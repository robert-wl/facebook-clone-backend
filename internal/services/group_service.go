package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

type GroupService struct {
	*Service
	NotificationService *NotificationService
}

func NewGroupService(s *Service, ns *NotificationService) *GroupService {
	return &GroupService{
		Service:             s,
		NotificationService: ns,
	}
}

func (s *GroupService) ClearGroupCache(groupID string, userID string) error {
	if err := s.RedisAdapter.Del([]string{"group", groupID}); err != nil {
		return err
	}

	if err := s.RedisAdapter.Del([]string{"group", "joined", userID}); err != nil {
		return err
	}

	if err := s.RedisAdapter.Del([]string{"group", "all", userID}); err != nil {
		return err
	}

	return nil
}

func (s *GroupService) MemberCount(obj *model.Group) (int, error) {
	var count int

	fmt.Println("CALLED")

	err := s.RedisAdapter.GetOrSet([]string{"group", obj.ID, "memberCount"}, &count, func() (interface{}, error) {
		count = int(s.DB.Model(&obj).Association("Members").Count())

		fmt.Println("COUNTER", s.DB.Model(&obj).Association("Members").Count())
		return count, nil
	}, 10*time.Minute)

	fmt.Println("COUNT", count)

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *GroupService) Joined(userID string, obj *model.Group) (string, error) {
	status := "not joined"
	var member *model.Member

	err := s.RedisAdapter.GetOrSet([]string{"group", obj.ID, "joined", userID}, &status, func() (interface{}, error) {
		if err := s.DB.First(&member, "group_id = ? AND user_id = ?", obj.ID, userID).Error; err == nil && member != nil {
			if member.Requested && !member.Approved {
				status = "pending"
			} else if member.Approved && !member.Requested {
				status = "joined"
			} else {
				status = "not accepted"
			}
		}

		return status, nil
	}, 10*time.Minute)

	if err != nil {
		return status, err
	}

	return status, nil
}

func (s *GroupService) IsAdmin(userID string, obj *model.Group) (bool, error) {
	var isAdmin = false
	err := s.RedisAdapter.GetOrSet([]string{"group", obj.ID, "isAdmin", userID}, &isAdmin, func() (interface{}, error) {
		if err := s.DB.First(&model.Member{}, "group_id = ? AND user_id = ? and role = ?", obj.ID, userID, "Admin").Error; err != nil {
			return false, nil
		}

		return true, nil
	}, 10*time.Minute)

	if err != nil {
		return false, err
	}

	return isAdmin, nil
}

func (s *GroupService) CreateGroup(userID string, group model.NewGroup) (*model.Group, error) {
	newGroup := &model.Group{
		ID:         uuid.NewString(),
		Name:       group.Name,
		About:      group.About,
		Privacy:    group.Privacy,
		Background: "",
		CreatedAt:  time.Now(),
	}

	if err := s.DB.Save(&newGroup).Error; err != nil {
		return nil, err
	}

	conversation := &model.Conversation{
		ID: uuid.NewString(),
	}

	if err := s.DB.Save(&conversation).Error; err != nil {
		return nil, err
	}

	conversationUser := &model.ConversationUsers{
		ConversationID: conversation.ID,
		UserID:         userID,
	}

	if err := s.DB.Save(&conversationUser).Error; err != nil {
		return nil, err
	}

	member := &model.Member{
		GroupID:  newGroup.ID,
		UserID:   userID,
		Approved: true,
		Role:     "Admin",
	}

	if err := s.DB.Save(&member).Error; err != nil {
		return nil, err
	}

	newGroup.ChatID = &conversation.ID

	if err := s.DB.Save(&newGroup).Error; err != nil {
		return nil, err
	}

	if group.Privacy != "public" {
		return newGroup, nil
	}

	go func() {
		s.NotificationService.CreateNewGroupNotification(userID, newGroup.ID)
	}()

	if err := s.ClearGroupCache(newGroup.ID, userID); err != nil {
		return nil, err
	}

	return newGroup, nil
}

func (s *GroupService) InviteToGroup(userID string, groupID string, inviteID string) (*model.Member, error) {
	var user *model.User

	member := &model.Member{
		GroupID:   groupID,
		UserID:    inviteID,
		Approved:  false,
		Role:      "member",
		Requested: false,
	}

	if err := s.DB.Save(&member).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Find(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	go func() {
		s.NotificationService.CreateGroupInvitationNotification(userID, *user, groupID, inviteID)
	}()

	if err := s.RedisAdapter.Del([]string{"group", groupID, "invite", inviteID}); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *GroupService) HandleRequest(userID string, groupID string) (*model.Member, error) {
	var member *model.Member

	if err := s.DB.First(&member, "group_id = ? AND user_id = ?", groupID, userID).Error; err != nil {
		member = &model.Member{
			GroupID:   groupID,
			UserID:    userID,
			Requested: true,
			Approved:  false,
			Role:      "member",
		}

		if err := s.DB.Save(&member).Error; err != nil {
			return nil, err
		}

		if err := s.ClearGroupCache(groupID, userID); err != nil {
			return nil, err
		}

		return member, nil
	}

	if member.Approved == false {
		member.Requested = false
		member.Approved = true

		if err := s.DB.Save(&member).Error; err != nil {
			return nil, err
		}

		var group *model.Group

		if err := s.DB.First(&group, "id = ?", groupID).Error; err != nil {
			return nil, err
		}

		conversationUser := &model.ConversationUsers{
			ConversationID: *group.ChatID,
			UserID:         member.UserID,
		}

		if err := s.DB.Save(&conversationUser).Error; err != nil {
			return nil, err
		}

		if err := s.ClearGroupCache(groupID, userID); err != nil {
			return nil, err
		}

		return member, nil
	}

	if err := s.DB.Delete(&member).Error; err != nil {
		return nil, err
	}

	if err := s.ClearGroupCache(groupID, userID); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *GroupService) UpdateGroupBackground(groupID string, background string) (*model.Group, error) {
	var group *model.Group

	if err := s.DB.First(&group, "id = ?", groupID).Update("background", background).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"group"}); err != nil {
		return nil, err
	}

	return group, nil
}

func (s *GroupService) UploadFile(userID string, groupID string, file model.NewGroupFile) (*model.GroupFile, error) {
	newFile := &model.GroupFile{
		ID:         uuid.NewString(),
		GroupID:    groupID,
		Name:       file.Name,
		Type:       file.Type,
		URL:        file.URL,
		UserID:     userID,
		UploadedAt: time.Now(),
	}

	var fileCount int64
	count := 0

	for {
		if err := s.DB.Find(&model.GroupFile{}, "group_id = ? AND name = ? AND type = ?", newFile.GroupID, newFile.Name, newFile.Type).Count(&fileCount).Error; err != nil {
			break
		}

		if fileCount == 0 {
			break
		}

		count++
		split := strings.Split(file.Name, ".")

		if len(split) >= 2 {
			newFile.Name = strings.Join(split[:len(split)-1], "") + fmt.Sprintf(" (%d)", count) + "." + split[len(split)-1]
		} else {
			newFile.Name = fmt.Sprintf("%s (%d)", file.Name, count)
		}
	}

	if err := s.DB.Save(&newFile).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		Preload("UploadedBy").
		First(&newFile, "id = ?", newFile.ID).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"group", groupID, "files"}); err != nil {
		return nil, err
	}

	return newFile, nil
}

func (s *GroupService) DeleteFile(fileID string) (*bool, error) {
	boolean := true
	var groupFile *model.GroupFile
	if err := s.DB.First(&groupFile, "id = ?", fileID).Delete(&model.GroupFile{}).Error; err != nil {
		boolean = false
		return &boolean, err
	}

	if err := s.RedisAdapter.Del([]string{"group", groupFile.GroupID, "files"}); err != nil {
		boolean = false
		return &boolean, err
	}

	return &boolean, nil
}

func (s *GroupService) ApproveMember(groupID string, userID string) (*model.Member, error) {
	var member *model.Member

	if err := s.DB.First(&member, "group_id = ? AND user_id = ?", groupID, userID).Update("approved", true).Update("requested", false).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"group", "joined", userID}); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *GroupService) DenyMember(groupID string, userID string) (*model.Member, error) {
	var member *model.Member

	if err := s.DB.Delete(&member, "group_id = ? AND user_id = ?", groupID, userID).Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"group", "joined", userID}); err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"group", groupID}); err != nil {
		return nil, err
	}
	return member, nil
}

func (s *GroupService) KickMember(groupID string, userID string) (*bool, error) {
	boolean := true
	if err := s.DB.Delete(&model.Member{}, "group_id = ? AND user_id = ?", groupID, userID).Error; err != nil {
		boolean = false
		return &boolean, err
	}

	if err := s.RedisAdapter.Del([]string{"group", "joined", userID}); err != nil {
		boolean = false
		return &boolean, err
	}

	if err := s.RedisAdapter.Del([]string{"group", groupID}); err != nil {
		boolean = false
		return &boolean, nil
	}

	return &boolean, nil
}

func (s *GroupService) LeaveGroup(userID string, groupID string) (string, error) {
	var member *model.Member

	if err := s.DB.First(&member, "group_id = ? AND user_id = ?", groupID, userID).Error; err != nil {
		return "not found", err
	}

	if member.Role == "Admin" {
		var adminCount int64
		var memberCount int64

		if err := s.DB.Find(&model.Member{}, "group_id = ? AND role = ?", groupID, "Admin").Count(&adminCount).Error; err != nil {
			return "not found", err
		}

		if err := s.DB.Find(&model.Member{}, "group_id = ? AND role = ?", groupID, "member").Count(&memberCount).Error; err != nil {
			return "not found", err
		}

		if adminCount == 1 && memberCount != 0 {
			return "not allowed", nil
		}
	}

	if err := s.DB.Delete(&model.Member{}, "group_id = ? AND user_id = ?", groupID, userID).Error; err != nil {
		return "not found", err
	}

	var memberCount int64

	if err := s.DB.Find(&model.Member{}, "group_id = ?", groupID).Count(&memberCount).Error; err != nil {
		return "not found", err
	}

	if memberCount == 0 {
		if err := s.DB.Delete(&model.Group{}, "id = ?", groupID).Error; err != nil {
			return "not found", err
		}
	}

	if err := s.RedisAdapter.Del([]string{"group", groupID}); err != nil {
		return "unknown error", err
	}
	if err := s.RedisAdapter.Del([]string{"group", "joined", userID}); err != nil {
		return "unknown error", err
	}

	return "success", nil
}

func (s *GroupService) PromoteMember(groupID string, userID string) (*model.Member, error) {
	var member *model.Member

	if err := s.DB.First(&member, "group_id = ? AND user_id = ?", groupID, userID).Update("role", "Admin").Error; err != nil {
		return nil, err
	}

	if err := s.RedisAdapter.Del([]string{"group", groupID, "isAdmin", userID}); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *GroupService) GetGroup(userID string, id string) (*model.Group, error) {
	var group *model.Group

	err := s.RedisAdapter.GetOrSet([]string{"group", "user", id}, &group, func() (interface{}, error) {
		subQuery := s.DB.
			Select("user_id").
			Where("group_id = ? and approved = true and requested = false", id).
			Table("members")

		if err := s.DB.
			Preload("Members").
			Preload("Members.User").
			Preload("Chat").
			Preload("Posts").
			Preload("Posts.User").
			Find(&group, "id = ? AND (privacy = ? or (privacy = ? AND ? IN (?)))", id, "public", "private", userID, subQuery).Error; err != nil {
			return nil, err
		}

		var members []*model.Member
		for _, member := range group.Members {
			if member.Approved && !member.Requested {
				members = append(members, member)
			}
		}

		group.Members = members
		group.MemberCount = len(members)

		return group, nil

	}, 10*time.Minute)

	fmt.Println("GROUP D", group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (s *GroupService) GetGroupInvite(userID string, id string) ([]*model.User, error) {
	var friendIDs []string
	var friendMemberIDs []string
	var users []*model.User

	err := s.RedisAdapter.GetOrSet([]string{"group", id, "invite", userID}, &users, func() (interface{}, error) {

		if err := s.DB.
			Model(&model.Friend{}).
			Where("sender_id = ? OR receiver_id = ?", userID, userID).
			Select("DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END", userID).
			Find(&friendIDs).Error; err != nil {
			return nil, err
		}

		if err := s.DB.
			Model(&model.Member{}).
			Where("group_id = ? AND user_id IN (?)", id, friendIDs).
			Select("user_id").
			Find(&friendMemberIDs).Error; err != nil {
			return nil, err
		}

		if len(friendMemberIDs) != 0 {
			if err := s.DB.Find(&users, "id IN (?) AND id NOT IN (?)", friendIDs, friendMemberIDs).Error; err != nil {
				return nil, err
			}
		} else {
			if err := s.DB.Find(&users, "id IN (?)", friendIDs).Error; err != nil {
				return nil, err
			}
		}

		return users, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *GroupService) GetGroups(userID string) ([]*model.Group, error) {
	var groups []*model.Group

	err := s.RedisAdapter.GetOrSet([]string{"group", "all", userID}, &groups, func() (interface{}, error) {
		if err := s.DB.Find(&groups, "privacy = ?", "public").Error; err != nil {
			return nil, err
		}

		return groups, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *GroupService) GetJoinedGroups(userID string) ([]*model.Group, error) {
	var groupIDs []*string
	var groups []*model.Group

	err := s.RedisAdapter.GetOrSet([]string{"group", "joined", userID}, &groups, func() (interface{}, error) {
		if err := s.DB.Model(&model.Member{}).Where("user_id = ? AND approved = ? AND requested = ?", userID, true, false).Select("group_id").Find(&groupIDs).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Preload("Chat").Find(&groups, "id in (?)", groupIDs).Error; err != nil {
			return nil, err
		}

		return groups, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *GroupService) GetGroupFiles(groupID string) ([]*model.GroupFile, error) {
	var files []*model.GroupFile

	err := s.RedisAdapter.GetOrSet([]string{"group", groupID, "files"}, &files, func() (interface{}, error) {

		if err := s.DB.
			Order("uploaded_at ASC").
			Preload("UploadedBy").
			Find(&files, "group_id = ?", groupID).Error; err != nil {
			return nil, err
		}
		return files, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *GroupService) GetJoinRequests(groupID string) ([]*model.Member, error) {
	var members []*model.Member

	err := s.RedisAdapter.GetOrSet([]string{"group", groupID, "join_requests"}, &members, func() (interface{}, error) {
		if err := s.DB.Preload("User").
			Find(&members, "group_id = ? AND requested = ? AND approved = ?", groupID, true, false).Error; err != nil {
			return nil, err
		}

		return members, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return members, nil
}

func (s *GroupService) GetFilteredGroups(filter string, pagination model.Pagination) ([]*model.Group, error) {
	var groups []*model.Group

	cacheKey := []string{"group", filter, strconv.Itoa(pagination.Start), strconv.Itoa(pagination.Limit)}

	err := s.RedisAdapter.GetOrSet(cacheKey, &groups, func() (interface{}, error) {
		if err := s.DB.
			Offset(pagination.Start).
			Limit(pagination.Limit).
			Find(&groups, "LOWER(name) LIKE LOWER(?) OR LOWER(about) LIKE LOWER(?)", "%"+filter+"%", "%"+filter+"%").Error; err != nil {
			return nil, err
		}

		return groups, nil
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return groups, nil
}
