package config

import (
    "context"
    "log"

    firebase "firebase.google.com/go/v4"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/option"
)


var (
	App      *firebase.App
	Firestore *firestore.Client
)

func InitFirebase() {
	opt := option.WithCredentialsFile("toko-indrajaya-firebase-adminsdk-fbsvc-0c37ccf24c.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}

	App = app
	Firestore = client
	log.Println("Firebase connected successfully.")
}
