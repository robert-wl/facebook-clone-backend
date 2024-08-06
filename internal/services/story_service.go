package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

type StoryService struct {
	*Service
	NotificationService *NotificationService
}

func NewStoryService(s *Service, ns *NotificationService) *StoryService {
	return &StoryService{
		Service:             s,
		NotificationService: ns,
	}
}

func (s *StoryService) CreateTextStory(userID string, input model.NewTextStory) (*model.Story, error) {

	story := &model.Story{
		ID:        uuid.NewString(),
		UserID:    userID,
		Text:      &input.Text,
		Font:      &input.Font,
		Color:     &input.Color,
		CreatedAt: time.Now(),
	}

	if err := s.DB.Save(&story).Error; err != nil {
		return nil, err
	}

	go func() {
		s.NotificationService.CreateStoryNotification(userID, story.ID)
	}()

	if err := s.RedisAdapter.Del([]string{"story", userID}); err != nil {
		return nil, err
	}

	return story, nil
}

func (s *StoryService) CreateImageStory(userID string, input model.NewImageStory) (*model.Story, error) {
	story := &model.Story{
		ID:        uuid.NewString(),
		UserID:    userID,
		Image:     &input.Image,
		CreatedAt: time.Now(),
	}

	if err := s.DB.Save(&story).Error; err != nil {
		return nil, err
	}

	go func() {
		s.NotificationService.CreateStoryNotification(userID, story.ID)
	}()

	if err := s.RedisAdapter.Del([]string{"story", userID}); err != nil {
		return nil, err
	}

	return story, nil
}

func (s *StoryService) GetStories(username string) ([]*model.Story, error) {
	var stories []*model.Story
	var user *model.User

	err := s.RedisAdapter.GetOrSet([]string{username}, &user, func() (interface{}, error) {
		if err := s.DB.First(&user, "username = ?", username).Error; err != nil {
			return nil, err
		}

		return user, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	err = s.RedisAdapter.GetOrSet([]string{"story", user.ID}, &stories, func() (interface{}, error) {
		if err := s.DB.Find(&stories, "user_id = ? AND DATE_PART('day', ? - created_at) = 0", user.ID, time.Now()).Error; err != nil {
			return nil, err
		}

		return stories, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return stories, nil
}

func (s *StoryService) GetUserWithStories(userID string) ([]*model.User, error) {
	var users []*model.User
	var friendIDs []string

	err := s.RedisAdapter.GetOrSet([]string{"story", userID}, &users, func() (interface{}, error) {
		if err := s.DB.
			Model(&model.Friend{}).
			Where("sender_id = ? OR receiver_id = ?", userID, userID).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID).
			Find(&friendIDs).Error; err != nil {
			return nil, err
		}

		friendIDs = append(friendIDs, userID)

		var userIDs []string
		if err := s.DB.Model(&model.Story{}).Where("user_id IN (?) AND DATE_PART('day', ? - created_at) = 0", friendIDs, time.Now()).Select("user_id").Find(&userIDs).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Model(&model.User{}).Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
			return nil, err
		}

		return users, nil
	}, 10*time.Second)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *StoryService) User(obj *model.Story) (*model.User, error) {
	var user *model.User

	err := s.RedisAdapter.GetOrSet([]string{obj.UserID}, &user, func() (interface{}, error) {
		if err := s.DB.First(&user, "id = ?", obj.UserID).Error; err != nil {
			return nil, err
		}

		return user, nil
	}, 10*60*time.Minute)

	if err != nil {
		return nil, err
	}

	return user, nil
}
