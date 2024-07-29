package database

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/helper"
	"math/rand"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

func generateUser() []model.User {
	var users []model.User

	db := GetDBInstance()
	for i := 0; i < 40; i++ {
		fmt.Println("Generating User")

		pw, _ := helper.EncryptPassword("password")

		fn := faker.FirstName()
		ln := faker.LastName()

		dob := time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000)))
		user := model.User{
			ID:                uuid.NewString(),
			FirstName:         fn,
			LastName:          ln,
			Username:          fmt.Sprintf("%s.%s", fn, ln),
			Email:             fmt.Sprintf("%s.%s@gmail.com", fn, ln),
			Password:          pw,
			Dob:               dob,
			Gender:            "",
			Active:            true,
			MiscId:            nil,
			Profile:           nil,
			Background:        nil,
			CreatedAt:         time.Now(),
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

func generatePost(user model.User) {
	db := GetDBInstance()

	num := rand.Intn(20) + 1
	for i := 0; i < num; i++ {
		fmt.Println("Generating Post")
		post := model.Post{
			ID:           uuid.NewString(),
			UserID:       user.ID,
			Content:      faker.Sentence(),
			Privacy:      "public",
			LikeCount:    0,
			CommentCount: 0,
			ShareCount:   0,
			CreatedAt:    time.Time{},
		}

		db.Create(&post)
	}
}

func generateFriend(sender model.User, receiver model.User) {
	db := GetDBInstance()

	fmt.Println("Generating Friend")
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
}

func FakeData() {
	users := generateUser()

	for _, user := range users {
		generatePost(user)
	}

	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			generateFriend(users[i], users[j])
		}
	}
}
