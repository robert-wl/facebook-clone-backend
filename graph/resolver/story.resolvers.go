package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

// CreateTextStory is the resolver for the createTextStory field.
func (r *mutationResolver) CreateTextStory(ctx context.Context, input model.NewTextStory) (*model.Story, error) {
	userID := ctx.Value("UserID").(string)

	story := &model.Story{
		ID:        uuid.NewString(),
		UserID:    userID,
		Text:      &input.Text,
		Font:      &input.Font,
		Color:     &input.Color,
		CreatedAt: time.Now(),
	}

	if err := r.DB.Save(&story).Error; err != nil {
		return nil, err
	}

	go func() {
		var userIDs []string
		var user *model.User

		if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
			return
		}

		subQuery := r.DB.
			Model(&model.Friend{}).
			Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID)

		subQueryBlocked := r.DB.
			Model(&model.BlockNotification{}).
			Where("(sender_id = ?)", userID).
			Select("DISTINCT receiver_id")

		if err := r.DB.
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
				StoryID: &story.ID,
				GroupID: nil,
			}

			if _, err := r.CreateNotification(ctx, *newNotification); err != nil {
				continue
			}
		}
	}()

	r.Redis.Del(ctx, fmt.Sprintf(`user:%s:stories`, userID))

	return story, nil
}

// CreateImageStory is the resolver for the createImageStory field.
func (r *mutationResolver) CreateImageStory(ctx context.Context, input model.NewImageStory) (*model.Story, error) {
	userID := ctx.Value("UserID").(string)

	story := &model.Story{
		ID:        uuid.NewString(),
		UserID:    userID,
		Image:     &input.Image,
		CreatedAt: time.Now(),
	}

	if err := r.DB.Save(&story).Error; err != nil {
		return nil, err
	}

	go func() {
		var userIDs []string
		var user *model.User

		if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
			return
		}

		subQuery := r.DB.
			Model(&model.Friend{}).
			Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", userID, userID, true).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID)

		subQueryBlocked := r.DB.
			Model(&model.BlockNotification{}).
			Where("(sender_id = ?)", userID).
			Select("DISTINCT receiver_id")

		if err := r.DB.
			Model(&model.User{}).
			Where("id IN (?) AND id NOT IN (?) AND id != ?", subQuery, subQueryBlocked, userID).
			Select("id").
			Find(&userIDs).Error; err != nil {
			return
		}

		for _, userIdz := range userIDs {

			newNotification := &model.NewNotification{
				Message: fmt.Sprintf("%s %s released a new story", user.FirstName, user.LastName),
				UserID:  userIdz,
				PostID:  nil,
				ReelID:  nil,
				StoryID: &story.ID,
				GroupID: nil,
			}

			if _, err := r.CreateNotification(ctx, *newNotification); err != nil {
				continue
			}
		}
	}()

	r.Redis.Del(ctx, fmt.Sprintf(`user:%s:stories`, userID))

	return story, nil
}

// GetStories is the resolver for the getStories field.
func (r *queryResolver) GetStories(ctx context.Context, username string) ([]*model.Story, error) {
	var stories []*model.Story
	var user *model.User

	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	if storiesSerialized, err := r.Redis.Get(ctx, fmt.Sprintf(`user:%s:stories`, user.ID)).Result(); err != nil {
		if err := r.DB.Find(&stories, "user_id = ? AND DATE_PART('day', ? - created_at) = 0", user.ID, time.Now()).Error; err != nil {
			return nil, err
		}

		if serializedStories, err := json.Marshal(stories); err != nil {
			return nil, err
		} else {
			r.Redis.Set(ctx, fmt.Sprintf(`user:%s:stories`, user.ID), serializedStories, 10*60*time.Minute)
		}
	} else {
		if err := json.Unmarshal([]byte(storiesSerialized), &stories); err != nil {
			return nil, err
		}
	}

	return stories, nil
}

// GetUserWithStories is the resolver for the getUserWithStories field.
func (r *queryResolver) GetUserWithStories(ctx context.Context) ([]*model.User, error) {
	userID := ctx.Value("UserID").(string)

	var friendIDs []string
	if err := r.DB.
		Model(&model.Friend{}).
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID).
		Find(&friendIDs).Error; err != nil {
		return nil, err
	}

	friendIDs = append(friendIDs, userID)

	var userIDs []string
	if err := r.DB.Model(&model.Story{}).Where("user_id IN (?) AND DATE_PART('day', ? - created_at) = 0", friendIDs, time.Now()).Select("user_id").Find(&userIDs).Error; err != nil {
		return nil, err
	}

	var users []*model.User
	if err := r.DB.Model(&model.User{}).Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *storyResolver) User(ctx context.Context, obj *model.Story) (*model.User, error) {
	var user *model.User

	if err := r.DB.First(&user, "id = ?", obj.UserID).Error; err != nil {
		return nil, err
	}

	if serializedUser, err := r.Redis.Get(ctx, fmt.Sprintf(`user:%s`, obj.UserID)).Result(); err != nil {

		if err := r.DB.First(&user, "id = ?", obj.UserID).Error; err != nil {
			return nil, err
		}

		if serializedUser, err := json.Marshal(user); err != nil {
			return nil, err
		} else {
			r.Redis.Set(ctx, fmt.Sprintf(`user:%s`, obj.UserID), serializedUser, 10*60*time.Minute)
		}

	} else {
		if err := json.Unmarshal([]byte(serializedUser), &user); err != nil {
			return nil, err
		}
	}

	return user, nil
}

// Story returns graph.StoryResolver implementation.
func (r *Resolver) Story() graph.StoryResolver { return &storyResolver{r} }

type storyResolver struct{ *Resolver }
