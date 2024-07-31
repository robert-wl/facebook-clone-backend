package adapter

import (
	"context"
	"fmt"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

type NotificationAdapter struct {
	resolver *mutationResolver
}

func (r *mutationResolver) createPostNotification(ctx context.Context, user model.User, postID string) {
	var userIDs []string

	subQuery := r.DB.
		Model(&model.Friend{}).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

	subQueryBlocked := r.DB.
		Model(&model.BlockNotification{}).
		Where("(sender_id = ?)", user.ID).
		Select("DISTINCT receiver_id")

	if err := r.DB.
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

		if _, err := r.CreateNotification(ctx, *newNotification); err != nil {
			continue
		}
	}
}
func (r *mutationResolver) createLikeNotification(ctx context.Context, userID string, postID string) {
	var user *model.User

	if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
		return
	}

	var userIDs []string

	subQuery := r.DB.
		Model(&model.Friend{}).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

	subQueryBlocked := r.DB.
		Model(&model.BlockNotification{}).
		Where("(sender_id = ?)", user.ID).
		Select("DISTINCT receiver_id")

	if err := r.DB.
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

		if _, err := r.CreateNotification(ctx, *newNotification); err != nil {
			continue
		}
	}
}
func (r *mutationResolver) createCommentNotification(ctx context.Context, user model.User, newComment model.NewComment) {
	var users []*model.User

	subQuery := r.DB.
		Model(&model.Friend{}).
		Where("(sender_id = ? OR receiver_id = ? AND accepted = ?)", user.ID, user.ID, true).
		Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", user.ID)

	subQueryBlocked := r.DB.
		Model(&model.BlockNotification{}).
		Where("(sender_id = ?)", user.ID).
		Select("DISTINCT receiver_id")

	if err := r.DB.Find(&users, "id IN (?) AND id NOT IN (?) AND id != ?", subQuery, subQueryBlocked, user.ID).Error; err != nil {
		return
	}

	if newComment.ParentPost == nil {
		var comment *model.Comment

		if err := r.DB.First(&comment, "id = ?", newComment.ParentComment).Error; err != nil {
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

			if _, err := r.CreateNotification(ctx, *newNotification); err != nil {
				continue
			}
		}
		r.Redis.Del(ctx, fmt.Sprintf("comment:%s:reply", *newComment.ParentComment))
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
			if _, err := r.CreateNotification(ctx, *newNotification); err != nil {
				fmt.Println(err)
				continue
			}
		}
		r.Redis.Del(ctx, fmt.Sprintf("post:%s:comment", *newComment.ParentPost))
	}
}
