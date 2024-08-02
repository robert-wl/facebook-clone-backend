package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

// LikeCount is the resolver for the likeCount field.
func (r *commentResolver) LikeCount(ctx context.Context, obj *model.Comment) (int, error) {
	return r.PostService.LikeCountComment(obj)
}

// ReplyCount is the resolver for the replyCount field.
func (r *commentResolver) ReplyCount(ctx context.Context, obj *model.Comment) (int, error) {
	return r.PostService.ReplyCount(obj)
}

// Liked is the resolver for the liked field.
func (r *commentResolver) Liked(ctx context.Context, obj *model.Comment) (*bool, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.LikedComment(userID, obj)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, newPost model.NewPost) (*model.Post, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.CreatePost(userID, newPost)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, newComment model.NewComment) (*model.Comment, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.CreateComment(userID, newComment)
}

// SharePost is the resolver for the sharePost field.
func (r *mutationResolver) SharePost(ctx context.Context, userID string, postID string) (*string, error) {
	return r.PostService.SharePost(userID, postID)
}

// LikePost is the resolver for the likePost field.
func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*model.PostLike, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.LikePost(userID, postID)
}

// Likecomment is the resolver for the likecomment field.
func (r *mutationResolver) Likecomment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.Likecomment(userID, commentID)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postID string) (*string, error) {
	return r.PostService.DeletePost(postID)
}

// LikeCount is the resolver for the likeCount field.
func (r *postResolver) LikeCount(ctx context.Context, obj *model.Post) (int, error) {
	return r.PostService.LikeCountPost(obj)
}

// CommentCount is the resolver for the commentCount field.
func (r *postResolver) CommentCount(ctx context.Context, obj *model.Post) (int, error) {
	return r.PostService.CommentCount(obj)
}

// Group is the resolver for the group field.
func (r *postResolver) Group(ctx context.Context, obj *model.Post) (*model.Group, error) {
	return r.PostService.Group(obj)
}

// Liked is the resolver for the liked field.
func (r *postResolver) Liked(ctx context.Context, obj *model.Post) (*bool, error) {
	userID := ctx.Value("UserID").(string)
	return r.PostService.LikedPost(userID, obj)
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: GetPost - getPost"))
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, pagination model.Pagination) ([]*model.Post, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.GetPosts(userID, pagination)
}

// GetGroupPosts is the resolver for the getGroupPosts field.
func (r *queryResolver) GetGroupPosts(ctx context.Context, groupID string, pagination model.Pagination) ([]*model.Post, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.GetGroupPosts(userID, groupID, pagination)
}

// GetCommentPost is the resolver for the getCommentPost field.
func (r *queryResolver) GetCommentPost(ctx context.Context, postID string) ([]*model.Comment, error) {
	return r.PostService.GetCommentPost(postID)
}

// GetFilteredPosts is the resolver for the getFilteredPosts field.
func (r *queryResolver) GetFilteredPosts(ctx context.Context, filter string, pagination model.Pagination) ([]*model.Post, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.GetFilteredPosts(userID, filter, pagination)
}

// GetGroupHomePosts is the resolver for the getGroupHomePosts field.
func (r *queryResolver) GetGroupHomePosts(ctx context.Context, pagination model.Pagination) ([]*model.Post, error) {
	userID := ctx.Value("UserID").(string)

	return r.PostService.GetGroupHomePosts(userID, pagination)
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

type commentResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
