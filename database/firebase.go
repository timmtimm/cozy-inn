package database

import (
	"context"
	"cozy-inn/util"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func InitDB() *firestore.Client {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: util.GetFirebaseEnv("project_id")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	return client
}
