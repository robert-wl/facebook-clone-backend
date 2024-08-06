package database

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

var profile = readJSON("/profile-image.json")
var imageVeryLarge = readJSON("/image-very-large.json")
var reelVideo = readJSON("/reel-video.json")
var video = readJSON("/video.json")
var groupFile = readJSON("/group-file.json")
var combinedImages = readJSON("/image-large.json", "/image-medium.json", "/image-very-large.json", "/image-small.json", "/image-very-small.json")

type JSONFile struct {
	Url       string `json:"url"`
	Directory string `json:"directory"`
	Type      string `json:"type"`
}

func readJSON(name ...string) []JSONFile {
	var jsonFilesResult []JSONFile
	for _, n := range name {
		dir := fmt.Sprintf("./%s", n)
		file, err := os.Open(dir)

		if err != nil {
			panic(err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)

		data, err := io.ReadAll(file)

		if err != nil {
			panic(err)
		}

		var jsonFiles []JSONFile

		err = json.Unmarshal(data, &jsonFiles)

		if err != nil {
			fmt.Println("Error when unmarshalling", n)
			panic(err)
		}

		jsonFilesResult = append(jsonFilesResult, jsonFiles...)
	}

	return jsonFilesResult
}

func convertToStr(file JSONFile) string {
	jsonStr, err := json.Marshal(file)

	if err != nil {
		panic(err)
	}

	return string(jsonStr)
}

func generateImage(imgType string) JSONFile {
	if imgType == "very-large" {
		return imageVeryLarge[rand.Intn(len(imageVeryLarge))]
	}
	return combinedImages[rand.Intn(len(combinedImages))]
}

func generateProfile() JSONFile {
	return profile[rand.Intn(len(profile))]
}

func generateVideo() JSONFile {
	return video[rand.Intn(len(video))]
}

func generateReelVideo() JSONFile {
	return reelVideo[rand.Intn(len(reelVideo))]
}

func generateGroupFile() JSONFile {
	return groupFile[rand.Intn(len(groupFile))]
}

func generateUser() []model.User {
	var users []model.User

	db := GetDBInstance()

	for i := 0; i < 40; i++ {

		pw, _ := utils.EncryptPassword("password")

		fn := faker.FirstName()
		ln := faker.LastName()

		profile := generateProfile()
		background := generateImage("very-large")

		var gender string

		if rand.Intn(10) > 5 {
			gender = "Male"
		} else {
			gender = "Female"
		}

		dob := time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000)))
		user := model.User{
			ID:                uuid.NewString(),
			FirstName:         fn,
			LastName:          ln,
			Username:          fmt.Sprintf("%s%s", fn, ln),
			Email:             fmt.Sprintf("%s.%s@gmail.com", fn, ln),
			Password:          pw,
			Dob:               dob,
			Gender:            gender,
			Active:            true,
			MiscId:            nil,
			Profile:           &profile.Url,
			Background:        &background.Url,
			CreatedAt:         time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
			FriendCount:       0,
			MutualCount:       0,
			NotificationCount: 0,
			Friended:          "",
			Theme:             "light",
		}

		db.Create(&user)

		users = append(users, user)
	}

	return users
}

func generatePost(user model.User) []model.Post {
	var posts []model.Post
	db := GetDBInstance()

	num := rand.Intn(20) + 1
	for i := 0; i < num; i++ {
		fmt.Println("Generating Post")

		var files []*string
		if rand.Intn(10) > 4 {
			if rand.Intn(10) > 8 {
				data := generateVideo()

				jsonStr := convertToStr(data)
				files = append(files, &jsonStr)

			} else {
				takeAmount := rand.Intn(10) + 1

				for i := 0; i < takeAmount; i++ {
					image := generateImage("")

					jsonStr := convertToStr(image)
					files = append(files, &jsonStr)
				}
			}
		}

		post := model.Post{
			ID:         uuid.NewString(),
			UserID:     user.ID,
			Content:    faker.Sentence(),
			Privacy:    "public",
			ShareCount: rand.Intn(100),
			CreatedAt:  time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
			Files:      files,
		}

		db.Create(&post)

		posts = append(posts, post)
	}

	return posts
}

func generateFriend(users []model.User) {
	db := GetDBInstance()

	var wg sync.WaitGroup
	for i := 0; i < len(users); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := i + 1; j < len(users); j++ {
				if rand.Intn(10) > 7 {
					continue
				}

				fmt.Println("Generating Friend")
				sender := users[i]
				receiver := users[j]

				if rand.Intn(10) > 3 {
					friend := model.Friend{
						SenderID:   sender.ID,
						ReceiverID: receiver.ID,
						Accepted:   true,
					}
					friend2 := model.Friend{
						SenderID:   receiver.ID,
						ReceiverID: sender.ID,
						Accepted:   true,
					}

					db.Create(&friend)
					db.Create(&friend2)
				} else {
					friend := model.Friend{
						SenderID:   sender.ID,
						ReceiverID: receiver.ID,
						Accepted:   false,
					}
					db.Create(&friend)
				}

			}
		}(i)
	}
	wg.Wait()
}

func generatePostLike(users []model.User, posts []model.Post) {
	db := GetDBInstance()

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func(user model.User) {
			defer wg.Done()
			for _, post := range posts {
				if rand.Intn(10) > 4 {
					fmt.Println("Generating Post Like")
					postLike := model.PostLike{
						UserID: user.ID,
						PostID: post.ID,
					}

					db.Create(&postLike)
				}
			}
		}(user)
	}
	wg.Wait()
}

func generatePostComment(users []model.User, posts []model.Post) []model.Comment {
	db := GetDBInstance()

	var wg1 sync.WaitGroup
	var comments []model.Comment
	for _, user := range users {
		wg1.Add(1)
		go func(user model.User) {
			defer wg1.Done()
			for _, post := range posts {
				if rand.Intn(10) > 8 {
					fmt.Println("Generating Post Comment")
					postComment := model.Comment{
						ID:           uuid.NewString(),
						UserID:       user.ID,
						Content:      faker.Sentence(),
						ParentPostID: &post.ID,
						CreatedAt:    time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
					}

					db.Create(&postComment)
					comments = append(comments, postComment)
				}
			}
		}(user)
	}
	wg1.Wait()

	var wg2 sync.WaitGroup
	var commentReplies []model.Comment
	for _, comment := range comments {
		wg2.Add(1)
		go func(comment model.Comment) {
			defer wg2.Done()
			for _, user := range users {
				if rand.Intn(10) > 8 {
					fmt.Println("Generating Comment Reply")
					commentReply := model.Comment{
						ID:              uuid.NewString(),
						UserID:          user.ID,
						Content:         faker.Sentence(),
						ParentCommentID: &comment.ID,
						CreatedAt:       time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
					}

					db.Create(&commentReply)
					commentReplies = append(commentReplies, commentReply)
				}
			}
		}(comment)
	}

	commentReplies = append(commentReplies, comments...)
	wg2.Wait()

	return commentReplies
}

func generateCommentLike(users []model.User, comments []model.Comment) {
	db := GetDBInstance()

	var wg sync.WaitGroup
	for _, comment := range comments {
		wg.Add(1)
		go func(comment model.Comment) {
			defer wg.Done()
			for _, user := range users {
				if rand.Intn(10) > 7 {
					fmt.Println("Generating Comment Like")
					commentLike := model.CommentLike{
						UserID:    user.ID,
						CommentID: comment.ID,
					}

					if err := db.Create(&commentLike); err != nil {
						return
					}
				}
			}
		}(comment)
	}
	wg.Wait()

	fmt.Println("USER COUNT", len(users))
	fmt.Println("COMMENT COUNT", len(comments))

}

func generateConversation(users []model.User) {
	db := GetDBInstance()

	var wg sync.WaitGroup
	randUser := users
	for i := 0; i < len(users); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			rand.Shuffle(len(randUser), func(i, j int) {
				randUser[i], randUser[j] = randUser[j], randUser[i]
			})

			randLength := rand.Intn(len(users))

			for j := 0; j < randLength; j++ {
				if randUser[j].ID == randUser[i].ID {
					continue
				}

				if err := db.Find(&model.ConversationUsers{}, "user_id = ? AND user_id = ?", randUser[i].ID, randUser[j].ID).Error; err == nil {
					continue
				}

				fmt.Println("Generating Conversation")
				sender := randUser[i]
				receiver := randUser[j]

				conversation := model.Conversation{
					ID:                  uuid.NewString(),
					LastSentMessageTime: time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
				}

				db.Create(&conversation)

				cUser1 := model.ConversationUsers{
					ConversationID: conversation.ID,
					UserID:         sender.ID,
				}

				cUser2 := model.ConversationUsers{
					ConversationID: conversation.ID,
					UserID:         receiver.ID,
				}

				db.Create(&cUser1)
				db.Create(&cUser2)
			}
		}(i)
	}

	wg.Wait()
}

func generateStories(users []model.User) {
	db := GetDBInstance()

	colors := []string{
		"lightblue",
		"pink",
		"lightgray",
		"orange",
	}

	fonts := []string{
		"normal",
		"roman",
	}

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)

		go func(user model.User) {
			defer wg.Done()
			for i := 0; i < rand.Intn(10); i++ {
				fmt.Println("Generating Story")
				textBr := faker.Sentence()

				if rand.Intn(10) > 5 {
					story := model.Story{
						ID:        uuid.NewString(),
						UserID:    user.ID,
						Font:      &fonts[rand.Intn(len(fonts))],
						Color:     &colors[rand.Intn(len(colors))],
						Text:      &textBr,
						CreatedAt: time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(20))),
					}

					db.Create(&story)
				} else {
					image := generateImage("")
					story := model.Story{
						ID:        uuid.NewString(),
						UserID:    user.ID,
						Image:     &image.Url,
						CreatedAt: time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(20))),
					}

					db.Create(&story)
				}
			}
		}(user)
	}

	wg.Wait()
}

func generateReels(users []model.User) []model.Reel {
	db := GetDBInstance()

	var wg sync.WaitGroup
	var reels []model.Reel
	for _, user := range users {
		wg.Add(1)

		go func(user model.User) {
			defer wg.Done()
			for i := 0; i < rand.Intn(10); i++ {
				fmt.Println("Generating Reel")
				video := generateReelVideo()

				reel := model.Reel{
					ID:         uuid.NewString(),
					UserID:     user.ID,
					Content:    faker.Sentence(),
					Video:      video.Url,
					ShareCount: rand.Intn(100),
					CreatedAt:  time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
				}

				db.Create(&reel)

				reels = append(reels, reel)
			}
		}(user)
	}
	wg.Wait()

	return reels
}

func generateReelComment(users []model.User, reels []model.Reel) []model.ReelComment {
	db := GetDBInstance()

	var wg1 sync.WaitGroup
	comments := []model.ReelComment{}
	for _, user := range users {
		wg1.Add(1)

		go func(user model.User) {
			defer wg1.Done()

			for _, reel := range reels {
				if rand.Intn(10) > 8 {
					fmt.Println("Generating Reel Comment")
					comment := model.ReelComment{
						ID:           uuid.NewString(),
						UserID:       user.ID,
						Content:      faker.Sentence(),
						ParentReelID: &reel.ID,
						CreatedAt:    time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
					}

					db.Create(&comment)
					comments = append(comments, comment)
				}
			}
		}(user)
	}

	wg1.Wait()

	var wg2 sync.WaitGroup
	for _, comment := range comments {
		wg2.Add(1)

		go func(comment model.ReelComment) {
			defer wg2.Done()
			for _, user := range users {
				if rand.Intn(10) > 8 {
					fmt.Println("Generating Reel Comment Reply")
					commentReply := model.ReelComment{
						ID:              uuid.NewString(),
						UserID:          user.ID,
						Content:         faker.Sentence(),
						ParentCommentID: &comment.ID,
						CreatedAt:       time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
					}

					db.Create(&commentReply)
					comments = append(comments, commentReply)
				}
			}
		}(comment)
	}

	wg2.Wait()

	return comments

}

func generateReelLike(users []model.User, reels []model.Reel) {
	db := GetDBInstance()

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)

		go func(user model.User) {
			defer wg.Done()
			for _, reel := range reels {
				if rand.Intn(10) > 7 {
					fmt.Println("Generating Reel Like")
					reelLike := model.ReelLike{
						UserID: user.ID,
						ReelID: reel.ID,
					}

					db.Create(&reelLike)
				}
			}
		}(user)
	}
	wg.Wait()
}

func generateReelCommentLike(users []model.User, comments []model.ReelComment) {
	db := GetDBInstance()

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)

		go func(user model.User) {
			defer wg.Done()
			for _, comment := range comments {
				if rand.Intn(10) > 7 {
					fmt.Println("Generating Reel Comment Like")
					commentLike := model.ReelCommentLike{
						ReelCommentID: comment.ID,
						UserID:        user.ID,
					}

					db.Create(&commentLike)
				}
			}
		}(user)
	}
	wg.Wait()
}

func generateGroup(users []model.User) map[string][]model.Member {
	db := GetDBInstance()
	var groupMemberMap = make(map[string][]model.Member)

	groupNum := rand.Intn(100) + 50

	var wg sync.WaitGroup
	for i := 0; i < groupNum; i++ {
		fmt.Println("Generating Group")

		var privacy string

		if rand.Intn(10) > 6 {
			privacy = "public"
		} else {
			privacy = "private"
		}

		group := model.Group{
			ID:          uuid.NewString(),
			Name:        fmt.Sprintf("%s %s ", faker.Word(), faker.Word()),
			About:       faker.Paragraph(),
			Privacy:     privacy,
			Background:  generateImage("very-large").Url,
			MemberCount: 0,
			ChatID:      nil,
			CreatedAt:   time.Now().Add(-time.Hour * time.Duration(1+rand.Intn(2000))),
		}

		db.Create(&group)

		var groupUsers []model.Member
		wg.Add(1)

		go func(group model.Group) {
			defer wg.Done()
			for _, user := range users {
				if rand.Intn(10) > 5 {
					fmt.Println("Generating Group User")

					var role string

					if rand.Intn(10) > 5 {
						role = "admin"
					} else {
						role = "member"
					}

					approved := true
					if role == "member" && rand.Intn(10) > 8 {
						approved = false
					}

					groupUser := model.Member{
						GroupID:   group.ID,
						UserID:    user.ID,
						Requested: false,
						Approved:  approved,
						Role:      role,
					}

					groupUsers = append(groupUsers, groupUser)

					db.Create(&groupUser)
				}
			}
		}(group)

		wg.Wait()

		if len(groupUsers) == 0 {
			groupUser := model.Member{
				GroupID:   group.ID,
				UserID:    users[rand.Intn(len(users))].ID,
				Requested: false,
				Approved:  true,
				Role:      "admin",
			}

			groupUsers = append(groupUsers, groupUser)
			db.Create(&groupUser)
		}

		groupMemberMap[group.ID] = groupUsers

		group.MemberCount = len(groupUsers)

		db.Save(&group)
	}

	return groupMemberMap
}

func generateGroupConversation(groups map[string][]model.Member) {
	db := GetDBInstance()

	for groupID, members := range groups {

		conversation := model.Conversation{
			ID:                  uuid.NewString(),
			GroupID:             &groupID,
			LastSentMessageTime: time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
		}

		db.Create(&conversation)

		var group model.Group

		if err := db.Where("id = ?", members[0].GroupID).First(&group).Error; err != nil {
			continue
		}

		group.ChatID = &conversation.ID

		db.Save(&group)

		for i := 0; i < len(members); i++ {
			fmt.Println("Generating Group Conversation")

			conversation := model.ConversationUsers{
				UserID:         members[i].UserID,
				ConversationID: conversation.ID,
			}

			db.Create(&conversation)
		}
	}
}

func generateGroupPosts(groups map[string][]model.Member) map[string][]model.Post {
	db := GetDBInstance()

	var groupPosts = make(map[string][]model.Post)
	for _, members := range groups {
		generateAmount := rand.Intn(100) + 10

		var posts []model.Post

		for i := 0; i < generateAmount; i++ {
			fmt.Println("Generating Group Post")

			var files []*string
			if rand.Intn(10) > 4 {
				if rand.Intn(10) > 8 {
					data := generateVideo()
					jsonStr := convertToStr(data)
					files = append(files, &jsonStr)

				} else {
					takeAmount := rand.Intn(10) + 1

					for i := 0; i < takeAmount; i++ {
						image := generateImage("")
						jsonStr := convertToStr(image)
						files = append(files, &jsonStr)
					}
				}
			}

			post := model.Post{
				ID:         uuid.NewString(),
				UserID:     members[rand.Intn(len(members))].UserID,
				Content:    faker.Sentence(),
				Privacy:    "group",
				ShareCount: rand.Intn(100),
				CreatedAt:  time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
				Files:      files,
				GroupID:    &members[0].GroupID,
			}

			posts = append(posts, post)

			db.Create(&post)
		}

		groupPosts[members[0].GroupID] = posts
	}

	return groupPosts
}

func generateGroupFiles(groups map[string][]model.Member) {
	db := GetDBInstance()

	for _, members := range groups {
		for i := 0; i < rand.Intn(50); i++ {
			fmt.Println("Generating Group File")

			fileData := generateGroupFile()

			file := model.GroupFile{
				ID:         uuid.NewString(),
				UserID:     members[rand.Intn(len(members))].UserID,
				GroupID:    members[0].GroupID,
				Name:       fileData.Directory,
				Type:       fileData.Type,
				URL:        fileData.Url,
				UploadedAt: time.Now().Add(-time.Hour * time.Duration(900+rand.Intn(1000))),
			}

			db.Create(&file)
		}
	}
}

func generateGroupPostComment(members map[string][]model.Member, posts map[string][]model.Post) []model.Comment {
	db := GetDBInstance()

	var comments []model.Comment

	for groupID, posts := range posts {
		for _, user := range members[groupID] {
			for _, post := range posts {
				if rand.Intn(10) > 8 {
					fmt.Println("Generating Group Post Comment")
					comment := model.Comment{
						ID:           uuid.NewString(),
						UserID:       user.UserID,
						Content:      faker.Sentence(),
						ParentPostID: &post.ID,
						CreatedAt:    time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
					}

					db.Create(&comment)
					comments = append(comments, comment)
				}
			}
		}
	}

	return comments
}

func generateGroupPostLike(members map[string][]model.Member, posts map[string][]model.Post) {
	db := GetDBInstance()

	for groupID, posts := range posts {
		for _, user := range members[groupID] {
			for _, post := range posts {
				if rand.Intn(10) > 4 {
					fmt.Println("Generating Group Post Like")
					postLike := model.PostLike{
						UserID: user.UserID,
						PostID: post.ID,
					}

					db.Create(&postLike)
				}
			}
		}
	}
}

func FakeData() {

	users := generateUser()

	var posts []model.Post
	for _, user := range users {
		userPosts := generatePost(user)

		posts = append(posts, userPosts...)
	}

	generatePostLike(users, posts)

	comments := generatePostComment(users, posts)

	generateCommentLike(users, comments)

	generateFriend(users)

	generateConversation(users)

	generateStories(users)

	reels := generateReels(users)

	generateReelLike(users, reels)

	reelComments := generateReelComment(users, reels)

	generateReelCommentLike(users, reelComments)

	groupData := generateGroup(users)

	generateGroupConversation(groupData)

	groupPosts := generateGroupPosts(groupData)

	generateGroupFiles(groupData)

	generateGroupPostComment(groupData, groupPosts)

	generateGroupPostLike(groupData, groupPosts)

	fmt.Println("Data Generated")
}
