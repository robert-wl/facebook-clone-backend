package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"time"
)

type NotificationService struct {
	*Service
}

func NewNotificationService(s *Service) *NotificationService {
	return &NotificationService{
		Service: s,
	}
}

func (s *NotificationService) CreateNotification(userID string, notification model.NewNotification) (*model.Notification, error) {
	var blocked *model.BlockNotification

	if err := s.DB.First(&blocked, "sender_id = ? AND receiver_id = ?", userID, notification.UserID).Error; err == nil && blocked != nil {
		return nil, nil
	}

	newNotification := &model.Notification{
		ID:        uuid.NewString(),
		Message:   notification.Message,
		UserID:    notification.UserID,
		SenderID:  userID,
		Seen:      false,
		PostID:    notification.PostID,
		ReelID:    notification.ReelID,
		StoryID:   notification.StoryID,
		GroupID:   notification.GroupID,
		CreatedAt: time.Now(),
	}

	if err := s.DB.Save(&newNotification).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"notification", notification.UserID}); err != nil {
		return nil, err
	}
	return newNotification, nil
}

func (s *NotificationService) GetUnreadNotifications(userID string) ([]*model.Notification, error) {
	var notifications []*model.Notification

	err := s.RedisAdapter.GetOrSet([]string{"notification", userID, "unread"}, &notifications, func() (interface{}, error) {
		if err := s.DB.
			Order("created_at DESC").
			Preload("Sender").
			Find(&notifications, "user_id = ? AND seen = false", userID).Error; err != nil {
			return nil, err
		}

		for _, notification := range notifications {
			notification.Seen = true
		}

		go func() {
			for _, notification := range notifications {
				if err := s.DB.Save(&notification).Error; err != nil {
					continue
				}
			}
		}()

		return notifications, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	for _, notification := range notifications {
		notification.Seen = false
	}

	return notifications, nil
}

func (s *NotificationService) BlockUser(userID string, username string) (*model.BlockNotification, error) {
	var user *model.User

	if err := s.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	var blockNotif *model.BlockNotification

	if err := s.DB.First(&blockNotif, "sender_id = ? AND receiver_id = ?", user.ID, userID).Error; err == nil && blockNotif != nil {

		if err := s.DB.Delete(&blockNotif).Error; err != nil {
			return nil, err
		}

		if err := s.RedisAdapter.Del([]string{"blocked", userID, user.ID}); err != nil {
			return nil, err
		}

		return blockNotif, nil
	}

	blockNotif = &model.BlockNotification{
		SenderID:   user.ID,
		ReceiverID: userID,
	}

	if err := s.DB.Save(&blockNotif).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"blocked", userID, user.ID}); err != nil {
		return nil, err
	}

	return blockNotif, nil
}

func (s *NotificationService) GetNotifications(userID string) ([]*model.Notification, error) {
	var notifications []*model.Notification

	err := s.RedisAdapter.GetOrSet([]string{"notification", userID}, &notifications, func() (interface{}, error) {
		if err := s.DB.
			Order("created_at DESC").
			Preload("Sender").
			Find(&notifications, "user_id = ?", userID).Error; err != nil {
			return nil, err
		}

		return notifications, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (s *NotificationService) CreatePostNotification(userID string, user model.User, postID string) {
	var userIDs []string

	subQuery := s.DB.
		Model(&model.Friend{}).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

	subQueryBlocked := s.DB.
		Model(&model.BlockNotification{}).
		Where("(sender_id = ?)", user.ID).
		Select("DISTINCT receiver_id")

	if err := s.DB.
		Model(&model.User{}).
		Where("id IN (?) AND id NOT IN (?) AND id != ?", subQuery, subQueryBlocked, user.ID).
		Select("id").
		Find(&userIDs).Error; err != nil {
		return
	}

	for _, userId := range userIDs {

		newNotification := &model.NewNotification{
			Message: fmt.Sprintf("%s %s posted a new post", user.FirstName, user.LastName),
			UserID:  userId,
			PostID:  &postID,
			ReelID:  nil,
			StoryID: nil,
			GroupID: nil,
		}

		if _, err := s.CreateNotification(userID, *newNotification); err != nil {
			continue
		}
	}
}
func (s *NotificationService) CreateLikeNotification(userID string, postID string) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return
	}

	var userIDs []string

	subQuery := s.DB.
		Model(&model.Friend{}).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

	subQueryBlocked := s.DB.
		Model(&model.BlockNotification{}).
		Where("(sender_id = ?)", user.ID).
		Select("DISTINCT receiver_id")

	if err := s.DB.
		Model(&model.User{}).
		Where("id IN (?) AND id NOT IN (?) AND id != ?", subQuery, subQueryBlocked, user.ID).
		Select("id").
		Find(&userIDs).Error; err != nil {
		return
	}

	for _, userId := range userIDs {

		newNotification := &model.NewNotification{
			Message: fmt.Sprintf("%s %s liked a post", user.FirstName, user.LastName),
			UserID:  userId,
			PostID:  &postID,
			ReelID:  nil,
			StoryID: nil,
			GroupID: nil,
		}

		if _, err := s.CreateNotification(userID, *newNotification); err != nil {
			continue
		}
	}
}
func (s *NotificationService) CreateCommentNotification(userID string, user model.User, newComment model.NewComment) {
	var users []*model.User

	subQuery := s.DB.
		Model(&model.Friend{}).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

	subQueryBlocked := s.DB.
		Model(&model.BlockNotification{}).
		Where("(sender_id = ?)", user.ID).
		Select("DISTINCT receiver_id")

	if err := s.DB.Find(&users, "id IN (?) AND id NOT IN (?) AND id != ?", subQuery, subQueryBlocked, user.ID).Error; err != nil {
		return
	}

	if newComment.ParentPost == nil {
		var comment *model.Comment

		if err := s.DB.First(&comment, "id = ?", newComment.ParentComment).Error; err != nil {
			return
		}

		for _, userDat := range users {

			newNotification := &model.NewNotification{
				Message: fmt.Sprintf("%s %s replied a comment", user.FirstName, user.LastName),
				UserID:  userDat.ID,
				PostID:  comment.ParentPostID,
				ReelID:  nil,
				StoryID: nil,
				GroupID: nil,
			}

			if _, err := s.CreateNotification(userID, *newNotification); err != nil {
				continue
			}
		}
	} else {
		for _, userDat := range users {

			newNotification := &model.NewNotification{
				Message: fmt.Sprintf("%s %s commented on a post", user.FirstName, user.LastName),
				UserID:  userDat.ID,
				PostID:  newComment.ParentPost,
				ReelID:  nil,
				StoryID: nil,
				GroupID: nil,
			}

			fmt.Println(newNotification)
			if _, err := s.CreateNotification(userID, *newNotification); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func (s *NotificationService) CreateShareNotification(userID string, user model.User, postID string) {
	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return
	}

	newNotification := &model.NewNotification{
		Message: fmt.Sprintf("%s %s released a new story", user.FirstName, user.LastName),
		UserID:  userID,
		PostID:  &postID,
		ReelID:  nil,
		StoryID: nil,
		GroupID: nil,
	}

	if _, err := s.CreateNotification(userID, *newNotification); err != nil {
		return
	}
}

func (s *NotificationService) CreateFriendRequestNotification(userID string, friend model.Friend) {
	newNotification := &model.NewNotification{
		Message: fmt.Sprintf("%s %s sent you a friend request", friend.Sender.FirstName, friend.Sender.LastName),
		UserID:  friend.ReceiverID,
		PostID:  nil,
		ReelID:  nil,
		StoryID: nil,
		GroupID: nil,
	}

	if _, err := s.CreateNotification(userID, *newNotification); err != nil {
		return
	}
}

func (s *NotificationService) CreateNewGroupNotification(userID string, groupID string) {
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

	for _, userId := range userIDs {

		newNotification := &model.NewNotification{
			Message: fmt.Sprintf("%s %s created a new group", user.FirstName, user.LastName),
			UserID:  userId,
			PostID:  nil,
			ReelID:  nil,
			StoryID: nil,
			GroupID: &groupID,
		}

		if _, err := s.CreateNotification(userID, *newNotification); err != nil {
			continue
		}
	}
}

func (s *NotificationService) CreateGroupInvitationNotification(userID string, user model.User, groupID string, inviteID string) {
	newNotification := &model.NewNotification{
		Message: fmt.Sprintf("%s %s invited you to join a group", user.FirstName, user.LastName),
		UserID:  inviteID,
		PostID:  nil,
		ReelID:  nil,
		StoryID: nil,
		GroupID: &groupID,
	}

	if _, err := s.CreateNotification(userID, *newNotification); err != nil {
		return
	}
}

func (s *NotificationService) CreateStoryNotification(userID string, storyID string) {
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

	for _, userId := range userIDs {

		newNotification := &model.NewNotification{
			Message: fmt.Sprintf("%s %s released a new story", user.FirstName, user.LastName),
			UserID:  userId,
			PostID:  nil,
			ReelID:  nil,
			StoryID: &storyID,
			GroupID: nil,
		}

		if _, err := s.CreateNotification(userID, *newNotification); err != nil {
			continue
		}
	}
}
