package postgresql

import (
	"fmt"
	"github.com/yahkerobertkertasnya/TPAWebBack/graph/model"
	"github.com/yahkerobertkertasnya/TPAWebBack/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

const defaultDatabase = "host=localhost user=postgres password=postgres dbname=TPAWeb port=5432 sslmode=disable TimeZone=Asia/Jakarta"

func GetInstance() *gorm.DB {
	if database == nil {
		dsn := helper.GetDotENVVariable("DATABASE_URL", defaultDatabase)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		database = db
	}

	return database
}

func MigrateDatabase() {
	db := GetInstance()

	err := db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.PostVisibility{},
		&model.PostTag{},
		&model.PostLike{},
		&model.Comment{},
		&model.CommentLike{},
		&model.Friend{},
		&model.Story{},
		&model.Conversation{},
		&model.Message{},
		&model.ConversationUsers{},
		&model.Reel{},
		&model.ReelLike{},
		&model.ReelComment{},
		&model.ReelCommentLike{},
		&model.Group{},
		&model.Member{},
		&model.GroupFile{},
		&model.Notification{},
		&model.BlockNotification{},
	)

	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}
