package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	helper2 "github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils/mail"
	"time"
)

type UserService struct {
	*Service
}

func NewUserService(s *Service) *UserService {
	return &UserService{
		Service: s,
	}
}

func (s *UserService) CreateUser(input model.NewUser) (*model.User, error) {
	activationId := uuid.NewString()

	user := &model.User{
		ID:         uuid.NewString(),
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Username:   input.Username,
		Email:      input.Email,
		Dob:        input.Dob,
		Gender:     input.Gender,
		Active:     true,
		MiscId:     &activationId,
		Profile:    nil,
		Background: nil,
		Theme:      "light",
	}

	if hashed, err := helper2.EncryptPassword(input.Password); err != nil {
		return nil, err
	} else {
		user.Password = hashed
	}

	//html := fmt.Sprintf(
	//	`
	//	<h1>Activate</h1>
	//	<a href="http://localhost:5173/activate/%s">Click here to activate your account</a>
	//	`, activationId)

	//_, err := mail.SendVerification(user.Email, "Activate Account", html)
	//if err != nil {
	//	return nil, err
	//}

	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	//s.Redis.Del(ctx, fmt.Sprintf("users"))

	return user, nil
}

func (s *UserService) ActivateUser(id string) (*model.User, error) {
	var user *model.User

	if err := s.DB.First(&user, "not active and misc_id = ?", id).Update("active", true).Update("misc_id", nil).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) AuthenticateUser(email string, password string) (string, error) {
	var user *model.User

	err := s.RedisAdapter.GetOrSet([]string{email}, &user, func() (interface{}, error) {
		if err := s.DB.First(&user, "email = ?", email).Error; err != nil {
			return nil, fmt.Errorf("credentials not found")
		}

		return user, nil
	}, time.Minute*60)

	fmt.Println("USER", user, err)
	if err != nil {
		return "", err
	}

	if user.Active == false {
		return "", fmt.Errorf("user is not active")
	}

	if !helper2.ComparePassword(user.Password, password) {
		return "", fmt.Errorf("incorrect password")
	}

	return helper2.CreateJWT(user.ID)
}

func (s *UserService) ForgotPassword(email string) (bool, error) {
	forgotId := uuid.NewString()

	html := fmt.Sprintf(
		`
		<h1>Reset Password</h1>
		<a href="http://localhost:5173/forgot/%s">Click here to reset your password</a>
		`, forgotId)

	_, err := mail.SendVerification(email, "Reset Password", html)
	if err != nil {
		return false, err
	}

	if err := s.DB.First(&model.User{}, "email = ?", email).Update("misc_id", forgotId).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (s *UserService) ResetPassword(id string, password string) (*model.User, error) {
	var user *model.User

	if err := s.DB.First(&user, "misc_id = ?", id).Error; err != nil {
		return nil, err
	}

	if helper2.ComparePassword(user.Password, password) {
		return nil, fmt.Errorf("password cannot be the same")
	}

	if hashedP, err := helper2.EncryptPassword(password); err != nil {
		return nil, err
	} else {
		user.Password = hashedP
	}

	user.MiscId = nil

	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUserProfile(userID string, profile string) (*model.User, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Update("profile", profile).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.DelType(user, []string{user.ID, user.Username, user.Email}); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUserBackground(userID string, background string) (*model.User, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Update("background", background).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.DelType(user, []string{user.ID, user.Username, user.Email}); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(userID string, input model.UpdateUser) (*model.User, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	if input.Password != "" {
		user.Password = input.Password
	}
	user.Gender = input.Gender

	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.DelType(user, []string{user.ID, user.Username, user.Email}); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateTheme(userID string, theme string) (*model.User, error) {
	var user *model.User

	if err := s.DB.First(&user, "id = ?", userID).Update("theme", theme).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.DelType(user, []string{user.ID, user.Username, user.Email}); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(userID string, username string) (*model.User, error) {
	var user *model.User

	var posts []*model.Post

	err := s.RedisAdapter.GetOrSet([]string{username}, &user, func() (interface{}, error) {
		if err := s.DB.First(&user, "username = ?", username).Error; err != nil {
			return nil, err
		}
		return user, nil
	}, time.Minute*60)

	if err != nil {
		return nil, err
	}

	subQueryFriend := s.DB.
		Select("*").
		Where("(sender_id = ? AND receiver_id = posts.user_id) or (sender_id = posts.user_id AND receiver_id = ?)", userID, userID).
		Table("friends")

	subQueryPrivate := s.DB.
		Select("user_id").
		Where("(post_id = posts.id)").
		Table("post_visibilities")

	if err := s.DB.
		Order("created_at desc").
		Preload("User").
		Preload("User").
		Preload("Likes").
		Preload("Comments").
		Preload("Visibility.User").
		Preload("PostTags.User").
		Find(&posts, "user_id = ? AND (privacy = ? OR (privacy = ? AND EXISTS(?)) OR (privacy = ? AND ? IN (?)) OR ?) AND group_id IS NULL", user.ID, "public", "friend", subQueryFriend, "specific", userID, subQueryPrivate, user.ID == userID).Error; err != nil {
		return nil, err
	}

	user.Posts = posts

	return user, nil
}

func (s *UserService) GetUsers() ([]*model.User, error) {
	var users []*model.User

	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CheckActivateLink(id string) (bool, error) {
	if err := s.DB.First(&model.User{}, "not active and misc_id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (s *UserService) CheckResetLink(id string) (bool, error) {
	if err := s.DB.First(&model.User{}, "active and misc_id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (s *UserService) GetAuth(userID string) (*model.User, error) {
	var user *model.User

	err := s.RedisAdapter.GetOrSet([]string{userID}, &user, func() (interface{}, error) {
		if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
			return nil, err
		}
		return user, nil
	}, time.Minute*60)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetFilteredUsers(userID string, filter string, pagination model.Pagination) ([]*model.User, error) {
	var users []*model.User

	if err := s.DB.
		Offset(pagination.Start).
		Limit(pagination.Limit).
		Where("id != ? AND (LOWER(first_name) LIKE LOWER(?) OR LOWER(last_name) LIKE LOWER(?) OR LOWER(username) LIKE LOWER(?))", userID, "%"+filter+"%", "%"+filter+"%", "%"+filter+"%").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) FriendCount(userID string) (int, error) {
	var friendCount int64

	err := s.RedisAdapter.GetOrSet([]string{"friend-count", userID}, &friendCount, func() (interface{}, error) {
		if err := s.DB.Find(&model.Friend{}, "(sender_id = ? or receiver_id = ?) and accepted = true", userID, userID).Count(&friendCount).Error; err != nil {
			return nil, err
		}

		return int(friendCount), nil
	}, time.Minute*60)

	if err != nil {
		return 0, err
	}

	return int(friendCount), nil
}

func (s *UserService) MutualCount(userID string, obj *model.User) (int, error) {
	var friendIDs []string
	var myFriendIDs []string
	var mutualCount int64

	err := s.RedisAdapter.GetOrSet([]string{"mutual-count", userID}, &mutualCount, func() (interface{}, error) {
		if err := s.DB.
			Model(&model.Friend{}).
			Where("sender_id = ? OR receiver_id = ?", obj.ID, obj.ID).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", obj.ID).
			Find(&friendIDs).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Model(&model.Friend{}).
			Where("sender_id = ? OR receiver_id = ? AND accepted = ?", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID).Find(&myFriendIDs).Error; err != nil {
			return nil, err
		}
		if err := s.DB.Find(&model.User{}, "id IN (?) AND id IN (?)", friendIDs, myFriendIDs).Count(&mutualCount).Error; err != nil {
			return nil, err
		}

		return int(mutualCount), nil
	}, time.Minute*60)

	if err != nil {
		return 0, err
	}

	return int(mutualCount), nil
}

func (s *UserService) NotificationCount(userID string) (int, error) {
	var notificationCount int64

	err := s.RedisAdapter.GetOrSet([]string{"notification-count", userID}, &notificationCount, func() (interface{}, error) {
		if err := s.DB.Find(&model.Notification{}, "user_id = ? AND seen = false", userID).Count(&notificationCount).Error; err != nil {
			return nil, err
		}

		return int(notificationCount), nil
	}, time.Minute*60)

	if err != nil {
		return 0, err
	}

	return int(notificationCount), nil
}

func (s *UserService) Friended(userID string, obj *model.User) (string, error) {
	var friend *model.Friend
	var status string

	if err := s.DB.First(&friend, "(sender_id = ? and receiver_id = ?) or (sender_id = ? and receiver_id = ?)", userID, obj.ID, obj.ID, userID).Error; err != nil {
		status = "not friends"
	} else {
		if friend.Accepted {
			status = "friends"
		} else {
			status = "pending"
		}
	}

	return status, nil
}

func (s *UserService) Blocked(userID string, obj *model.User) (bool, error) {
	var blocked bool
	err := s.RedisAdapter.GetOrSet([]string{"blocked", userID, obj.ID}, &blocked, func() (interface{}, error) {
		var blocked *model.BlockNotification

		if err := s.DB.First(&blocked, "sender_id = ? AND receiver_id = ?", userID, obj.ID).Error; err == nil && blocked != nil {
			return true, nil
		}

		return false, nil

	}, time.Minute*60)

	if err != nil {
		return false, err
	}

	return blocked, nil
}

func (s *UserService) GetRandomUsers(amount int) ([]*model.User, error) {
	var users []*model.User

	if err := s.DB.Order("RANDOM()").Limit(amount).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
