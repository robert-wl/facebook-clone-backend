package services

import (
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"time"
)

type FriendsService struct {
	*Service
	NotificationService *NotificationService
}

func NewFriendsService(s *Service, ns *NotificationService) *FriendsService {
	return &FriendsService{
		Service:             s,
		NotificationService: ns,
	}
}

func (s *FriendsService) AddFriend(userID string, friendInput model.FriendInput) (*model.Friend, error) {
	var friendModel *model.Friend

	if err := s.DB.First(&friendModel, "sender_id = ? and receiver_id = ?", friendInput.Sender, friendInput.Receiver).Error; err != nil {
		friend := &model.Friend{
			SenderID:   friendInput.Sender,
			ReceiverID: friendInput.Receiver,
			Accepted:   false,
		}

		if err := s.DB.Save(&friend).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Preload("Sender").Preload("Receiver").First(&friend, "sender_id = ? and receiver_id = ?", friend.SenderID, friend.ReceiverID).Error; err != nil {
			return nil, err
		}

		if err := s.RedisAdapter.Del([]string{"friends", friendInput.Sender}); err != nil {
			return nil, err
		}
		if err := s.RedisAdapter.Del([]string{"friends", friendInput.Receiver}); err != nil {
			return nil, err
		}

		go func() {
			s.NotificationService.CreateFriendRequestNotification(userID, *friend)
		}()

		return friend, nil
	} else {
		if err := s.RedisAdapter.Del([]string{"friends", friendInput.Sender}); err != nil {
			return nil, err
		}
		if err := s.RedisAdapter.Del([]string{"friends", friendInput.Receiver}); err != nil {
			return nil, err
		}

		return nil, s.DB.Delete(&friendModel).Error
	}
}

func (s *FriendsService) AcceptFriend(userID string, friend string) (*model.Friend, error) {
	var friendModel *model.Friend

	if err := s.DB.First(&friendModel, "sender_id = ? and receiver_id = ?", friend, userID).Update("Accepted", true).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"friends", userID}); err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"friends", friend}); err != nil {
		return nil, err
	}

	return friendModel, nil
}

func (s *FriendsService) RejectFriend(userID string, friend string) (*model.Friend, error) {

	friendModel := &model.Friend{
		SenderID:   friend,
		ReceiverID: userID,
		Accepted:   false,
	}

	if err := s.DB.Delete(&model.Friend{}, "(sender_id = ? AND receiver_id = ?)", friend, userID).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"friends", userID}); err != nil {
		return nil, err
	}
	if err := s.RedisAdapter.Del([]string{"friends", friend}); err != nil {
		return nil, err
	}

	return friendModel, nil
}

func (s *FriendsService) GetFriends(userID string) ([]*model.User, error) {
	var users []*model.User

	err := s.RedisAdapter.GetOrSet([]string{"friends", userID}, &users, func() (interface{}, error) {
		subQuery := s.DB.
			Model(&model.Friend{}).
			Where("((sender_id = ? OR receiver_id = ?) AND accepted = ?)", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID)

		if err := s.DB.Find(&users, "id IN (?)", subQuery).Error; err != nil {
			return nil, err
		}

		return users, nil
	}, 60*time.Minute)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *FriendsService) GetFriendRequests(userID string) ([]*model.User, error) {
	var users []*model.User

	err := s.RedisAdapter.GetOrSet([]string{"friends", userID, "request"}, &users, func() (interface{}, error) {
		subQuery := s.DB.
			Model(&model.Friend{}).
			Where("receiver_id = ? AND accepted = ?", userID, false).
			Select("DISTINCT sender_id")

		if err := s.DB.Find(&users, "id IN (?)", subQuery).Error; err != nil {
			return nil, err
		}

		return users, nil
	}, 60*time.Minute)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *FriendsService) GetUserFriends(username string) ([]*model.User, error) {
	var user *model.User
	var users []*model.User

	if err := s.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	err := s.RedisAdapter.GetOrSet([]string{"friends", user.ID}, &users, func() (interface{}, error) {
		subQuery := s.DB.
			Model(&model.Friend{}).
			Where("((sender_id = ? OR receiver_id = ?) AND accepted = true)", user.ID, user.ID).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

		if err := s.DB.Find(&users, "id IN (?)", subQuery).Error; err != nil {
			return nil, err
		}

		return users, nil
	}, 60*time.Minute)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *FriendsService) GetUserMutuals(userID string, username string) ([]*model.User, error) {
	var users []*model.User
	var user *model.User
	var friendIDs []string
	var myFriendIDs []string

	if err := s.DB.Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	err := s.RedisAdapter.GetOrSet([]string{"friends", userID, user.ID, "mutuals"}, &users, func() (interface{}, error) {
		if err := s.DB.
			Model(&model.Friend{}).
			Where("sender_id = ? OR receiver_id = ?", user.ID, user.ID).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID).
			Find(&friendIDs).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Model(&model.Friend{}).
			Where("(sender_id = ? OR receiver_id = ?) AND accepted = ?", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID).Find(&myFriendIDs).Error; err != nil {
			return nil, err
		}
		if err := s.DB.Find(&users, "id IN (?) AND id IN (?)", friendIDs, myFriendIDs).Error; err != nil {
			return nil, err
		}

		return users, nil
	}, 60*time.Minute)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *FriendsService) GetPeopleMightKnow(userID string) ([]*model.User, error) {
	var userIds []string
	var userFriendIds []string
	var users []*model.User

	err := s.RedisAdapter.GetOrSet([]string{"user", userID, "people_might_know"}, &users, func() (interface{}, error) {
		if err := s.DB.Model(&model.Friend{}).
			Where("(sender_id = ? OR receiver_id = ?) AND accepted = ?", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID).Find(&userIds).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Model(&model.Friend{}).
			Where("(sender_id IN (?) OR receiver_id IN (?)) AND accepted = ?", userIds, userIds, true).
			Select("DISTINCT CASE WHEN sender_id IN (?) THEN receiver_id ELSE sender_id END", userIds).Find(&userFriendIds).Error; err != nil {
			return nil, err
		}

		if err := s.DB.
			Limit(5).
			Find(&users, "id IN (?) AND id NOT IN (?) AND id != ?", userFriendIds, userIds, userID).Error; err != nil {
			return nil, err
		}

		return users, nil
	}, 60*time.Minute)

	if err != nil {
		return nil, err
	}

	return users, nil
}
