package firebase

import (
	"context"
	"firebase.google.com/go/storage"
	"fmt"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/helper"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const defaultCredentialApp = "serviceAccountKey.json"

var firebaseApp *firebase.App
var client *storage.Client

func GetAppInstance() (*firebase.App, error) {

	if firebaseApp == nil {
		credential := helper.GetDotENVVariable("FIREBASE_URL")

		fmt.Println(credential)

		if credential == "" {
			credential = defaultCredentialApp
		}

		opt := option.WithCredentialsFile(credential)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing app: %v", err)
		}

		firebaseApp = app
	}
	return firebaseApp, nil
}

func GetStorageClientInstance() (*storage.Client, error) {

	if client == nil {
		app, err := GetAppInstance()

		if err != nil {
			return nil, err
		}

		cli, err := app.Storage(context.Background())

		if err != nil {
			return nil, err
		}

		client = cli
	}

	return client, nil
}
