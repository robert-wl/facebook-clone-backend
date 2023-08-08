package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/TPAWebBack/graph"
	"github.com/yahkerobertkertasnya/TPAWebBack/graph/model"
	"github.com/yahkerobertkertasnya/TPAWebBack/helper"
	"github.com/yahkerobertkertasnya/TPAWebBack/helper/mail"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	activationId := uuid.NewString()

	user := &model.User{
		ID:         uuid.NewString(),
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Username:   input.Username,
		Email:      input.Email,
		Dob:        input.Dob,
		Gender:     input.Gender,
		Active:     false,
		MiscId:     &activationId,
		Profile:    nil,
		Background: nil,
	}

	if hashed, err := helper.EncryptPassword(input.Password); err != nil {
		return nil, err
	} else {
		user.Password = hashed
	}

	html := fmt.Sprintf(
		`
		<h1>Activate</h1>
		<a href="http://localhost:5173/activate/%s">Click here to activate your account</a>
		`, activationId)

	_, err := mail.SendVerification(user.Email, "Activate Account", html)
	if err != nil {
		return nil, err
	}

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ActivateUser is the resolver for the activateUser field.
func (r *mutationResolver) ActivateUser(ctx context.Context, id string) (*model.User, error) {
	var user *model.User

	if err := r.DB.First(&user, "not active and misc_id = ?", id).Update("active", true).Update("misc_id", nil).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticateUser is the resolver for the authenticateUser field.
func (r *mutationResolver) AuthenticateUser(ctx context.Context, email string, password string) (string, error) {
	var user *model.User

	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return "", fmt.Errorf("credentials not found")
	}

	if user.Active == false {
		return "", fmt.Errorf("user is not active")
	}

	if !helper.ComparePassword(user.Password, password) {
		return "", fmt.Errorf("incorrect password")
	}

	return helper.CreateJWT(user.ID)
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
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

	if err := r.DB.First(&model.User{}, "email = ?", email).Update("misc_id", forgotId).Error; err != nil {
		return false, err
	}

	return true, nil
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, id string, password string) (*model.User, error) {
	var user *model.User

	if err := r.DB.First(&user, "misc_id = ?", id).Error; err != nil {
		return nil, err
	}

	if helper.ComparePassword(user.Password, password) {
		return nil, fmt.Errorf("password cannot be the same")
	}

	if hashedP, err := helper.EncryptPassword(password); err != nil {
		return nil, err
	} else {
		user.Password = hashedP
	}

	user.MiscId = nil

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserProfile is the resolver for the updateUserProfile field.
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, profile string) (*model.User, error) {
	var user *model.User
	userID := ctx.Value("UserID")

	if err := r.DB.First(&user, "id = ?", userID).Update("profile", profile).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserBackground is the resolver for the updateUserBackground field.
func (r *mutationResolver) UpdateUserBackground(ctx context.Context, background string) (*model.User, error) {
	var user *model.User
	userID := ctx.Value("UserID")

	if err := r.DB.First(&user, "id = ?", userID).Update("background", background).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	var user *model.User
	userID := ctx.Value("UserID").(string)

	if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	if input.Password != "" {
		user.Password = input.Password
	}
	user.Gender = input.Gender

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	var friendCount int64
	var friend *model.Friend

	userID := ctx.Value("UserID").(string)

	if err := r.DB.Preload("Posts").Preload("Posts.User").First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Find(&model.Friend{}, "(sender_id = ? or receiver_id = ?) and accepted = true", userID, userID).Count(&friendCount).Error; err != nil {

	}

	user.FriendCount = int(friendCount)

	if err := r.DB.First(&friend, "(sender_id = ? and receiver_id = ?) or (sender_id = ? and receiver_id = ?)", userID, user.ID, user.ID, userID).Error; err != nil {
		user.Friended = "not friends"
	} else {
		if friend.Accepted {
			user.Friended = "friends"
		} else {
			user.Friended = "pending"
		}
	}

	return user, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// CheckActivateLink is the resolver for the checkActivateLink field.
func (r *queryResolver) CheckActivateLink(ctx context.Context, id string) (bool, error) {
	if err := r.DB.First(&model.User{}, "not active and misc_id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}

// CheckResetLink is the resolver for the checkResetLink field.
func (r *queryResolver) CheckResetLink(ctx context.Context, id string) (bool, error) {
	if err := r.DB.First(&model.User{}, "active and misc_id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}

// GetAuth is the resolver for the getAuth field.
func (r *queryResolver) GetAuth(ctx context.Context) (*model.User, error) {
	var user *model.User

	if err := r.DB.First(&user, "id = ?", ctx.Value("UserID")).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
