package services

import (
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"strconv"
	"time"
)

type PostService struct {
	*Service
	NotificationService *NotificationService
	MessageService      *MessagesService
}

func NewPostService(s *Service, ns *NotificationService, ms *MessagesService) *PostService {
	return &PostService{
		Service:             s,
		NotificationService: ns,
		MessageService:      ms,
	}
}

func (s *PostService) LikeCountComment(obj *model.Comment) (int, error) {
	var likeCount int64

	err := s.RedisAdapter.GetOrSet([]string{"comment", obj.ID, "like"}, &likeCount, func() (interface{}, error) {
		likeCount = s.DB.Model(obj).Association("Likes").Count()

		return likeCount, nil

	}, time.Minute*5)

	if err != nil {
		return 0, err
	}

	return int(likeCount), nil
}

func (s *PostService) ReplyCount(obj *model.Comment) (int, error) {
	var replyCount int64

	err := s.RedisAdapter.GetOrSet([]string{"comment", obj.ID, "reply"}, &replyCount, func() (interface{}, error) {
		replyCount = s.DB.Model(obj).Association("Comments").Count()

		return replyCount, nil

	}, time.Minute*5)

	if err != nil {
		return 0, err
	}

	return int(replyCount), nil
}

func (s *PostService) LikedComment(userID string, obj *model.Comment) (*bool, error) {
	boolean := false

	err := s.RedisAdapter.GetOrSet([]string{"liked", obj.ID, userID}, &boolean, func() (interface{}, error) {
		var commentLike *model.CommentLike

		if err := s.DB.First(&commentLike, "comment_id = ? AND user_id = ?", obj.ID, userID).Error; err == nil && commentLike != nil {
			boolean = true
		}

		return &boolean, nil

	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return &boolean, nil
}

func (s *PostService) CreatePost(userID string, newPost model.NewPost) (*model.Post, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	boolVar := false
	post := &model.Post{
		ID:           uuid.NewString(),
		UserID:       userID,
		User:         user,
		Content:      newPost.Content,
		Privacy:      newPost.Privacy,
		LikeCount:    0,
		CommentCount: 0,
		ShareCount:   0,
		GroupID:      newPost.GroupID,
		Files:        newPost.Files,
		Liked:        &boolVar,
		CreatedAt:    time.Now(),
	}

	if err := s.DB.Save(&post).Error; err != nil {
		return nil, err
	}

	for _, vsb := range newPost.Visibility {
		visibility := &model.PostVisibility{
			PostID: post.ID,
			UserID: *vsb,
		}

		if err := s.DB.Save(visibility).Error; err != nil {
			return nil, err
		}
	}

	for i, _ := range newPost.Tags {
		tagModel := &model.PostTag{
			PostID: post.ID,
			UserID: *newPost.Tags[i],
		}

		if err := s.DB.Create(tagModel).Error; err != nil {
			return nil, err
		}
	}

	if err := s.DB.
		Preload("PostTags.User").
		Preload("Visibility.User").
		First(&post).Error; err != nil {
		return nil, err
	}

	go func() {
		s.NotificationService.CreatePostNotification(userID, *user, post.ID)
	}()

	return post, nil
}

func (s *PostService) CreateComment(userID string, newComment model.NewComment) (*model.Comment, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	boolVar := false
	comment := &model.Comment{
		ID:              uuid.NewString(),
		UserID:          userID,
		User:            user,
		Content:         newComment.Content,
		Liked:           &boolVar,
		LikeCount:       0,
		ReplyCount:      0,
		ParentPostID:    newComment.ParentPost,
		ParentCommentID: newComment.ParentComment,
		CreatedAt:       time.Now(),
	}

	if err := s.DB.Save(&comment).Error; err != nil {
		return nil, err
	}

	go func() {
		s.NotificationService.CreateCommentNotification(userID, *user, newComment)
	}()

	if newComment.ParentComment != nil {
		if err := s.RedisAdapter.Del([]string{"comment", *newComment.ParentComment, "reply"}); err != nil {
			return nil, err
		}
	}

	if newComment.ParentPost != nil {
		if err := s.RedisAdapter.Del([]string{"posts", *newComment.ParentPost, "comment"}); err != nil {
			return nil, err
		}
	}

	return comment, nil
}

func (s *PostService) SharePost(userID string, postID string) (*string, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	conv, err := s.MessageService.CreateConversation(userID, user.Username)

	if err != nil {
		return nil, err
	}

	_, err = s.MessageService.SendMessage(userID, conv.ID, nil, nil, &postID)

	if err != nil {
		return nil, err
	}

	var post *model.Post

	if err := s.DB.First(&post, "id = ?", postID).Error; err != nil {
		return nil, err
	}

	post.ShareCount = post.ShareCount + 1

	if err := s.DB.Save(&post).Error; err != nil {
		return nil, err
	}

	go func() {
		s.NotificationService.CreateShareNotification(userID, *user, postID)
	}()

	return &conv.ID, nil
}

func (s *PostService) LikePost(userID string, postID string) (*model.PostLike, error) {
	var postLike *model.PostLike

	if err := s.DB.First(&postLike, "post_id = ? AND user_id = ?", postID, userID).Error; err != nil || postLike == nil {
		postLike = &model.PostLike{
			PostID: postID,
			UserID: userID,
		}
		if err := s.DB.Save(&postLike).Error; err != nil {
			return nil, err
		}

		go func() {
			s.NotificationService.CreateLikeNotification(userID, postID)
		}()

	} else {
		if err := s.DB.Delete(&postLike).Error; err != nil {
			return nil, err
		}
	}

	if err := s.RedisAdapter.Del([]string{"liked", postID, userID}); err != nil {
		return nil, err
	}
	if err := s.RedisAdapter.Del([]string{"post", postID, "like"}); err != nil {
		return nil, err
	}

	return postLike, nil
}

func (s *PostService) Likecomment(userID string, commentID string) (*model.CommentLike, error) {
	var commentLike *model.CommentLike

	if err := s.DB.First(&commentLike, "comment_id = ? AND user_id = ?", commentID, userID).Error; err != nil {
		commentLike = &model.CommentLike{
			CommentID: commentID,
			UserID:    userID,
		}
		if err := s.DB.Save(&commentLike).Error; err != nil {
			return nil, err
		}
	} else {
		if err := s.DB.Delete(&commentLike).Error; err != nil {
			return nil, err
		}
	}

	s.RedisAdapter.Del([]string{"liked", commentID, userID})
	s.RedisAdapter.Del([]string{"comment", commentID, "like"})
	return commentLike, nil
}

func (s *PostService) DeletePost(postID string, userID string) (*string, error) {
	var post *model.Post

	if err := s.DB.First(&post, "id = ?", postID).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Delete(&model.Post{}, "id = ?", postID).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"posts", postID}); err != nil {
		return nil, err
	}

	if post.GroupID != nil {
		if err := s.RedisAdapter.Del([]string{"group", *post.GroupID, "user", userID}); err != nil {
			return nil, err
		}
	}

	return &postID, nil
}

func (s *PostService) LikeCountPost(obj *model.Post) (int, error) {
	var count int64

	err := s.RedisAdapter.GetOrSet([]string{"post", obj.ID, "like"}, &count, func() (interface{}, error) {
		count = s.DB.Model(obj).Association("Likes").Count()

		return count, nil

	}, time.Minute*5)

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (s *PostService) CommentCount(obj *model.Post) (int, error) {
	var commentCount int64

	err := s.RedisAdapter.GetOrSet([]string{"post", obj.ID, "comment"}, &commentCount, func() (interface{}, error) {
		commentCount = s.DB.Model(obj).Association("Comments").Count()

		return commentCount, nil

	}, time.Minute*5)

	if err != nil {
		return 0, err
	}

	return int(commentCount), nil
}

func (s *PostService) Group(obj *model.Post) (*model.Group, error) {
	if obj.GroupID == nil {
		return nil, nil
	}

	var group *model.Group

	err := s.RedisAdapter.GetOrSet([]string{"group", *obj.GroupID}, &group, func() (interface{}, error) {
		if err := s.DB.
			Find(&group, "id = ?", *obj.GroupID).Error; err != nil {
			return nil, err
		}

		return group, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return group, nil
}

func (s *PostService) LikedPost(userID string, obj *model.Post) (*bool, error) {
	boolean := false
	var postLike *model.PostLike

	cacheKey := []string{"post", obj.ID, userID}
	err := s.RedisAdapter.GetOrSet(cacheKey, &boolean, func() (interface{}, error) {
		if err := s.DB.First(&postLike, "post_id = ? AND user_id = ?", obj.ID, userID).Error; err == nil && postLike != nil {
			boolean = true
		}

		return &boolean, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return &boolean, nil
}

func (s *PostService) GetPosts(userID string, pagination model.Pagination) ([]*model.Post, error) {
	var posts []*model.Post

	cacheKeys := []string{"posts", userID, strconv.Itoa(pagination.Start), strconv.Itoa(pagination.Limit)}

	err := s.RedisAdapter.GetOrSet(cacheKeys, &posts, func() (interface{}, error) {
		subQueryFriend := s.DB.
			Select("*").
			Where("(sender_id = ? AND receiver_id = posts.user_id) or (sender_id = posts.user_id AND receiver_id = ?)", userID, userID).
			Table("friends")

		subQueryPrivate := s.DB.
			Select("user_id").
			Where("(post_id = posts.id)").
			Table("post_visibilities")

		subQueryGroup := s.DB.
			Select("group_id").
			Where("user_id = ? AND approved = ?", userID, true).
			Table("members")

		if err := s.DB.
			Order("created_at desc").
			Preload("User").
			Preload("User").
			Preload("Likes").
			Preload("Comments").
			Preload("Visibility.User").
			Preload("PostTags.User").
			Offset(pagination.Start).
			Limit(pagination.Limit).
			Find(&posts, "(privacy = ? OR (privacy = ? AND EXISTS(?)) OR (privacy = ? AND ? IN (?)) OR group_id IN (?))", "public", "friend", subQueryFriend, "specific", userID, subQueryPrivate, subQueryGroup).Error; err != nil {
			return nil, err
		}

		return posts, nil
	}, time.Minute*2)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetGroupPosts(userID string, groupID string, pagination model.Pagination) ([]*model.Post, error) {
	var posts []*model.Post

	cacheKey := []string{"group", "posts", groupID, strconv.Itoa(pagination.Start), strconv.Itoa(pagination.Limit)}

	err := s.RedisAdapter.GetOrSet(cacheKey, &posts, func() (interface{}, error) {
		if err := s.DB.
			Order("created_at desc").
			Preload("User").
			Preload("Likes").
			Preload("Comments").
			Offset(pagination.Start).
			Limit(pagination.Limit).
			Find(&posts, "group_id = ?", groupID).Error; err != nil {
			return nil, err
		}

		return posts, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetCommentPost(postID string) ([]*model.Comment, error) {
	var comments []*model.Comment

	cacheKey := []string{"post", postID, "comment"}

	err := s.RedisAdapter.GetOrSet(cacheKey, &comments, func() (interface{}, error) {

		if err := s.DB.
			Preload("User").
			Preload("Likes").
			Preload("Comments").
			Preload("Comments.User").
			Find(&comments, "parent_post_id = ?", postID).Error; err != nil {
			return nil, err
		}

		return comments, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *PostService) GetFilteredPosts(userID string, filter string, pagination model.Pagination) ([]*model.Post, error) {
	var posts []*model.Post

	cacheKey := []string{"posts", filter, strconv.Itoa(pagination.Start), strconv.Itoa(pagination.Limit)}

	err := s.RedisAdapter.GetOrSet(cacheKey, &posts, func() (interface{}, error) {
		subQueryFriend := s.DB.
			Select("*").
			Where("(sender_id = ? AND receiver_id = posts.user_id) or (sender_id = posts.user_id AND receiver_id = ?)", userID, userID).
			Table("friends")

		subQueryPrivate := s.DB.
			Select("user_id").
			Where("(post_id = posts.id)").
			Table("post_visibilities")

		subQueryGroup := s.DB.
			Select("group_id").
			Where("user_id = ? AND approved = ?", userID, true).
			Table("members")

		if err := s.DB.
			Order("created_at desc").
			Preload("User").
			Preload("User").
			Preload("Likes").
			Preload("Comments").
			Preload("Visibility.User").
			Preload("PostTags.User").
			Offset(pagination.Start).
			Limit(pagination.Limit).
			Find(&posts, "id = ? OR ((privacy = ? OR (privacy = ? AND EXISTS(?)) OR (privacy = ? AND ? IN (?)) OR group_id IN (?)) AND LOWER(content) LIKE LOWER(?))", filter, "public", "friend", subQueryFriend, "specific", userID, subQueryPrivate, subQueryGroup, "%"+filter+"%").Error; err != nil {
			return nil, err
		}

		return posts, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetGroupHomePosts(userID string, pagination model.Pagination) ([]*model.Post, error) {
	var posts []*model.Post

	cacheKey := []string{"post", "group", userID, strconv.Itoa(pagination.Start), strconv.Itoa(pagination.Limit)}

	err := s.RedisAdapter.GetOrSet(cacheKey, &posts, func() (interface{}, error) {
		subQueryGroupMembers := s.DB.
			Select("group_id").
			Where("user_id = ? AND approved = ?", userID, true).
			Table("members")

		subQueryGroup := s.DB.
			Select("group_id").
			Where("privacy = 'public'").
			Table("groups")

		if err := s.DB.
			Order("created_at desc").
			Preload("User").
			Preload("User").
			Preload("Likes").
			Preload("Comments").
			Preload("Visibility.User").
			Preload("PostTags.User").
			Offset(pagination.Start).
			Limit(pagination.Limit).
			Find(&posts, "group_id IN (?) OR group_id IN (?)", subQueryGroupMembers, subQueryGroup).Error; err != nil {
			return nil, err
		}

		return posts, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
