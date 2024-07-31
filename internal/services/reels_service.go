package services

import (
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"time"
)

type ReelsService struct {
	*Service
}

func NewReelsService(s *Service) *ReelsService {
	return &ReelsService{
		Service: s,
	}
}

func (s *ReelsService) CreateReel(userID string, reel model.NewReel) (*model.Reel, error) {
	newReel := &model.Reel{
		ID:           uuid.NewString(),
		UserID:       userID,
		Content:      reel.Content,
		Video:        reel.Video,
		LikeCount:    0,
		CommentCount: 0,
		ShareCount:   0,
		CreatedAt:    time.Now(),
	}

	if err := s.DB.Save(&newReel).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		Preload("User").
		First(&newReel).Error; err != nil {
		return nil, err
	}

	go func() {
		var userIDs []string
		var user *model.User

		if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
			return
		}

		subQuery := s.DB.
			Model(&model.Friend{}).
			Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID)

		subQueryBlocked := s.DB.
			Model(&model.BlockNotification{}).
			Where("(sender_id = ?)", userID).
			Select("DISTINCT receiver_id")

		if err := s.DB.
			Model(&model.User{}).
			Where("id IN (?) AND id NOT IN (?) AND id != ?", subQuery, subQueryBlocked, userID).
			Select("id").
			Find(&userIDs).Error; err != nil {
			return
		}

	}()

	return newReel, nil
}

func (s *ReelsService) CreateReelComment(userID string, comment model.NewReelComment) (*model.ReelComment, error) {
	newComment := &model.ReelComment{
		ID:              uuid.NewString(),
		UserID:          userID,
		Content:         comment.Content,
		LikeCount:       0,
		ReplyCount:      0,
		ParentReelID:    comment.ParentReel,
		ParentCommentID: comment.ParentComment,
		CreatedAt:       time.Time{},
	}

	if err := s.DB.Save(&newComment).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		Preload("User").
		Preload("ParentReel").
		Preload("ParentReel.User").
		Preload("ParentComment").
		Preload("ParentComment.User").
		First(&newComment).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"reel", *comment.ParentReel}); err != nil {
		return nil, err
	}

	return newComment, nil
}

func (s *ReelsService) LikeReel(userID string, reelID string) (*model.ReelLike, error) {
	reelLike := &model.ReelLike{
		ReelID: reelID,
		UserID: userID,
	}

	if err := s.DB.First(&model.ReelLike{}, "reel_id = ? and user_id = ?", reelID, userID).Error; err == nil {

		if err := s.DB.Delete(&reelLike).Error; err != nil {
			return nil, err
		}

		return reelLike, nil
	}

	if err := s.DB.Save(&reelLike).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		Preload("User").
		First(&reelLike).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"reel", reelID, userID}); err != nil {
		return nil, err
	}

	return reelLike, nil
}

func (s *ReelsService) LikeReelComment(userID string, reelCommentID string) (*model.ReelCommentLike, error) {
	reelCommentLike := &model.ReelCommentLike{
		ReelCommentID: reelCommentID,
		UserID:        userID,
	}

	if err := s.DB.First(&model.ReelCommentLike{}, "reel_comment_id = ? and user_id = ?", reelCommentID, userID).Error; err == nil {

		if err := s.DB.Delete(&reelCommentLike).Error; err != nil {
			return nil, err
		}

		return reelCommentLike, nil
	}

	if err := s.DB.Save(&reelCommentLike).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		Preload("User").
		First(&reelCommentLike).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"reel", reelCommentID, "likeCount"}); err != nil {
		return nil, err
	}
	if err := s.RedisAdapter.Del([]string{"reel", reelCommentID, "comment", "like", userID}); err != nil {
		return nil, err
	}

	return reelCommentLike, nil
}

func (s *ReelsService) GetReels() ([]*string, error) {
	var reelsID []*string

	if err := s.DB.
		Model(&model.Reel{}).
		Order("RANDOM()").
		Select("id").
		Find(&reelsID).Error; err != nil {
		return nil, err
	}

	return reelsID, nil
}

func (s *ReelsService) GetReel(userID string, id string) (*model.Reel, error) {
	var reel *model.Reel

	err := s.RedisAdapter.GetOrSet([]string{"reel", id, userID}, &reel, func() (interface{}, error) {
		if err := s.DB.
			Preload("User").
			Preload("Likes").
			Preload("Comments").
			First(&reel, "id = ?", id).Error; err != nil {
			return nil, err
		}

		reel.LikeCount = int(s.DB.Model(reel).Association("Likes").Count())
		reel.CommentCount = int(s.DB.Model(reel).Association("Comments").Count())

		liked := false

		if err := s.DB.First(&model.ReelLike{}, "reel_id = ? AND user_id = ?", id, userID).Error; err == nil {
			liked = true
		}

		reel.Liked = &liked

		return reel, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return reel, nil
}

func (s *ReelsService) GetReelComments(reelID string) ([]*model.ReelComment, error) {
	var comments []*model.ReelComment

	err := s.RedisAdapter.GetOrSet([]string{"reel", reelID, "comments"}, &comments, func() (interface{}, error) {
		if err := s.DB.
			Preload("User").
			Preload("ParentReel").
			Preload("ParentComment").
			Preload("Likes").
			Preload("Comments").
			Preload("Comments.User").
			Find(&comments, "parent_reel_id = ?", reelID).Error; err != nil {
			return nil, err
		}

		return comments, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *ReelsService) LikeCount(obj *model.ReelComment) (int, error) {
	var count int

	err := s.RedisAdapter.GetOrSet([]string{"reel", obj.ID, "comment", "like_count"}, &count, func() (interface{}, error) {
		count = int(s.DB.Model(obj).Association("Likes").Count())

		return count, nil
	}, 10*time.Minute)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *ReelsService) ReplyCount(obj *model.ReelComment) (int, error) {
	var count int

	err := s.RedisAdapter.GetOrSet([]string{"reel", obj.ID, "comment", "reply_count"}, &count, func() (interface{}, error) {
		count = int(s.DB.Model(obj).Association("Comments").Count())

		return count, nil
	}, 10*time.Minute)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *ReelsService) Comments(obj *model.ReelComment) ([]*model.ReelComment, error) {
	var comments []*model.ReelComment

	err := s.RedisAdapter.GetOrSet([]string{"reel", obj.ID, "comment", "reply"}, &comments, func() (interface{}, error) {
		if err := s.DB.Preload("User").Find(&comments, "parent_comment_id = ?", obj.ID).Error; err != nil {
			return nil, err
		}

		return comments, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *ReelsService) Liked(userID string, obj *model.ReelComment) (*bool, error) {
	var count int64
	boolean := false

	err := s.RedisAdapter.GetOrSet([]string{"reel", obj.ID, "comment", "like", userID}, &boolean, func() (interface{}, error) {
		if err := s.DB.Find(&model.ReelCommentLike{}, "reel_comment_id = ? and user_id = ?", obj.ID, userID).Count(&count).Error; err == nil && count != 0 {
			boolean = true
		}

		return &boolean, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return &boolean, nil
}
