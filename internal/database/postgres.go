package database

import (
	"fmt"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func GetDBInstance() *gorm.DB {
	if database == nil {
		host := utils.GetDotENVVariable("DATABASE_HOST", "localhost")
		user := utils.GetDotENVVariable("DATABASE_USER", "postgres")
		password := utils.GetDotENVVariable("DATABASE_PASSWORD", "postgres")
		dbname := utils.GetDotENVVariable("DATABASE_NAME", "facebook")
		port := utils.GetDotENVVariable("DATABASE_PORT", "5432")
		sslmode := utils.GetDotENVVariable("DATABASE_SSLMODE", "disable")
		timezone := utils.GetDotENVVariable("DATABASE_TIMEZONE", "Asia/Jakarta")

		url := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s", host, user, password, port, sslmode, timezone)

		db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))

		url = fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		database = db
	}

	return database
}

func DropDatabase() {
	db := GetDBInstance()
	db.Migrator().DropTable(
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
}

func MigrateDatabase() {
	db := GetDBInstance()
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

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
