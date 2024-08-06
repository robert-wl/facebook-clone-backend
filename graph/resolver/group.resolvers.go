package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

// MemberCount is the resolver for the memberCount field.
func (r *groupResolver) MemberCount(ctx context.Context, obj *model.Group) (int, error) {
	return r.GroupService.MemberCount(obj)
}

// Joined is the resolver for the joined field.
func (r *groupResolver) Joined(ctx context.Context, obj *model.Group) (string, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.Joined(userID, obj)
}

// IsAdmin is the resolver for the isAdmin field.
func (r *groupResolver) IsAdmin(ctx context.Context, obj *model.Group) (bool, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.IsAdmin(userID, obj)
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, group model.NewGroup) (*model.Group, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.CreateGroup(userID, group)
}

// InviteToGroup is the resolver for the inviteToGroup field.
func (r *mutationResolver) InviteToGroup(ctx context.Context, groupID string, userID string) (*model.Member, error) {
	userIDSender := ctx.Value("UserID").(string)

	return r.GroupService.InviteToGroup(userIDSender, groupID, userID)
}

// HandleRequest is the resolver for the handleRequest field.
func (r *mutationResolver) HandleRequest(ctx context.Context, groupID string) (*model.Member, error) {
	userID := ctx.Value("UserID").(string)

	return r.GroupService.HandleRequest(userID, groupID)
}

// UpdateGroupBackground is the resolver for the updateGroupBackground field.
func (r *mutationResolver) UpdateGroupBackground(ctx context.Context, groupID string, background string) (*model.Group, error) {
	return r.GroupService.UpdateGroupBackground(groupID, background)
}

// UploadFile is the resolver for the uploadFile field.
func (r *mutationResolver) UploadFile(ctx context.Context, groupID string, file model.NewGroupFile) (*model.GroupFile, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.UploadFile(userID, groupID, file)
}

// DeleteFile is the resolver for the deleteFile field.
func (r *mutationResolver) DeleteFile(ctx context.Context, fileID string) (*bool, error) {
	return r.GroupService.DeleteFile(fileID)
}

// ApproveMember is the resolver for the approveMember field.
func (r *mutationResolver) ApproveMember(ctx context.Context, groupID string, userID string) (*model.Member, error) {
	return r.GroupService.ApproveMember(groupID, userID)
}

// DenyMember is the resolver for the denyMember field.
func (r *mutationResolver) DenyMember(ctx context.Context, groupID string, userID string) (*model.Member, error) {
	return r.GroupService.DenyMember(groupID, userID)
}

// KickMember is the resolver for the kickMember field.
func (r *mutationResolver) KickMember(ctx context.Context, groupID string, userID string) (*bool, error) {
	return r.GroupService.KickMember(groupID, userID)
}

// LeaveGroup is the resolver for the leaveGroup field.
func (r *mutationResolver) LeaveGroup(ctx context.Context, groupID string) (string, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.LeaveGroup(userID, groupID)
}

// PromoteMember is the resolver for the promoteMember field.
func (r *mutationResolver) PromoteMember(ctx context.Context, groupID string, userID string) (*model.Member, error) {
	return r.GroupService.PromoteMember(groupID, userID)
}

// GetGroup is the resolver for the getGroup field.
func (r *queryResolver) GetGroup(ctx context.Context, id string) (*model.Group, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.GetGroup(userID, id)
}

// GetGroupInvite is the resolver for the getGroupInvite field.
func (r *queryResolver) GetGroupInvite(ctx context.Context, id string) ([]*model.User, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.GetGroupInvite(userID, id)
}

// GetGroups is the resolver for the getGroups field.
func (r *queryResolver) GetGroups(ctx context.Context) ([]*model.Group, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.GetGroups(userID)
}

// GetJoinedGroups is the resolver for the getJoinedGroups field.
func (r *queryResolver) GetJoinedGroups(ctx context.Context) ([]*model.Group, error) {
	userID := ctx.Value("UserID").(string)
	return r.GroupService.GetJoinedGroups(userID)
}

// GetGroupFiles is the resolver for the getGroupFiles field.
func (r *queryResolver) GetGroupFiles(ctx context.Context, groupID string) ([]*model.GroupFile, error) {
	return r.GroupService.GetGroupFiles(groupID)
}

// GetJoinRequests is the resolver for the getJoinRequests field.
func (r *queryResolver) GetJoinRequests(ctx context.Context, groupID string) ([]*model.Member, error) {
	return r.GroupService.GetJoinRequests(groupID)
}

// GetFilteredGroups is the resolver for the getFilteredGroups field.
func (r *queryResolver) GetFilteredGroups(ctx context.Context, filter string, pagination model.Pagination) ([]*model.Group, error) {
	return r.GroupService.GetFilteredGroups(filter, pagination)
}

// Group returns graph.GroupResolver implementation.
func (r *Resolver) Group() graph.GroupResolver { return &groupResolver{r} }

type groupResolver struct{ *Resolver }
