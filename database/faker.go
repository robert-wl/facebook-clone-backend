package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/helper"
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

		profile := fmt.Sprintf("https://i.pravatar.cc/300?img=%d", rand.Intn(70)+1)

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
			Profile:           &profile,
			Background:        nil,
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

func generatePost(user model.User) {
	images := []string{
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239889599-zeubp5w4zd-7afb0491c91b2f9e9aac56667c3be677.jpg?alt=media\u0026token=611497b0-4729-4a3c-a712-e5b030c76c98",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239889927-binnrjn0rls-Wonders-of-the-World-Pyramids-1030x538.png?alt=media\u0026token=60946674-0a1f-450f-88f6-3cdd452005e8",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239890364-4lgu1idb3f3-header.jpg?alt=media\u0026token=67cbc129-8050-4de2-b81d-632a00849b54",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239890642-bgv6gy7619g-wp4535284.webp?alt=media\u0026token=ae06e0fd-aa17-432c-badf-4cbe8022db30",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891088-2q3ng215db7-pixel-art-creature-sword-hyper-light-drifter-wallpaper-preview.jpg?alt=media\u0026token=9c3dfa20-6eec-4307-8a71-0ead6ff977cd",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891289-s5iziys6x8-aesthetic-super-mario-running-desktop-wallpaper-preview.jpg?alt=media\u0026token=d49fd349-994a-414e-8b46-d7de4a33e11d",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891629-1nprmy82w9j-elden-ring-landscape-game-art-video-game-art-video-games-hd-wallpaper-preview.jpg?alt=media\u0026token=4b4b01c3-f59e-4f9c-9914-43d320393a86",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239891872-f2bskviqt7p-HD-wallpaper-anime-landscape-ai-city.jpg?alt=media\u0026token=2cf51657-d5d0-40a5-b2fe-1b5fc1f41b30\",\"directory\": \"post/1722239891872-f2bskviqt7p-HD-wallpaper-anime-landscape-ai-city.jpg",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239892077-4ylhf2mrqdt-anime-landscape-anime-art-painting-sea-wallpaper-preview.jpg?alt=media\u0026token=b35c31fb-a4e8-4ca5-8345-a245c6f1c43d",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722239892440-bvg0j89wa9-88016860_p0_master1200.jpg?alt=media\u0026token=59dd9e3a-1d3f-43e3-b6e8-0eea996171b7",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722238113910-m8xehm01d1-image_5.png?alt=media\u0026token=c3a85375-58f6-4636-b435-a7fc04969b9b",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722238114728-depr6qvc1z7-image_6.png?alt=media\u0026token=7408b912-d806-4995-b05f-5006745f426e",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722238115256-9tfeismfm7w-88016860_p0_master1200.jpg?alt=media\u0026token=8256a127-59cc-4caf-8830-a42f94d66c41",
	}

	videos := []string{
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251656896-7ipol8g19se-Y2meta.app-CURIOSITY%20-%20Featuring%20Richard%20Feynman-(1080p).mp4?alt=media\u0026token=27ba7263-6665-4adc-807a-4c1838781faa",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251660098-2b4lh2hkqd1-Y2meta.app-ocean-(1080p).mp4?alt=media\u0026token=13589535-70bb-434e-ab90-aece2b8f6162",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251681222-59123q78wri-Y2meta.app-Indonesia%20_%20Cinematic%20Travel%20Video%20_%20Stock%20Footage-(1080p).mp4?alt=media\u0026token=49353b71-e427-41e5-859e-49e6293b0120",
		"https://firebasestorage.googleapis.com/v0/b/tpaweb-da706.appspot.com/o/post%2F1722251717038-b6zphnp94sf-Y2meta.app-JAPAN%20-%20Like%20You've%20Never%20Seen%20Before%20_%20Stock%20Footage-(1080p).mp4?alt=media\u0026token=f5332066-844c-49b3-8ed5-2a5eb37d54ac",
	}

	resolutions := [][]string{
		{"1920", "1080"},
		{"1280", "720"},
		{"1366", "768"},
		{"1600", "900"},
		{"800", "600"},
		{"1024", "768"},
		{"1280", "1024"},
		{"720", "480"},
	}

	db := GetDBInstance()

	num := rand.Intn(20) + 1
	for i := 0; i < num; i++ {
		fmt.Println("Generating Post")

		var files []*string
		if rand.Intn(10) > 4 {
			if rand.Intn(10) > 8 {
				video := videos[rand.Intn(len(videos))]
				data := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"video/mp4\"}", video, video)
				files = append(files, &data)

			} else {
				takeAmount := rand.Intn(10) + 1

				for i := 0; i < takeAmount; i++ {
					if rand.Intn(10) > 8 {
						image := images[rand.Intn(len(images))]
						data := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"image/jpeg\"}", image, image)

						files = append(files, &data)
					} else {
						resolution := resolutions[rand.Intn(len(resolutions))]
						id := rand.Intn(1000)

						image := fmt.Sprintf("https://picsum.photos/id/%d/%s/%s", id, resolution[0], resolution[1])
						data := fmt.Sprintf("{\"url\": \"%s\",\"directory\": \"%s\",\"type\": \"image/jpeg\"}", image, image)

						files = append(files, &data)
					}
				}
			}
		}

		post := model.Post{
			ID:           uuid.NewString(),
			UserID:       user.ID,
			Content:      faker.Sentence(),
			Privacy:      "public",
			LikeCount:    0,
			CommentCount: 0,
			ShareCount:   0,
			CreatedAt:    time.Now().Add(-time.Hour * time.Duration(90000+rand.Intn(1000000))),
			Files:        files,
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
