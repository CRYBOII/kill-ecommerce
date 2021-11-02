package configs

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCN = ConnectDB()

var mongo_uri = os.Getenv("MONGO_URI")
var clientOptions = options.Client().ApplyURI(mongo_uri)

func ConnectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection Successfully!!")
	return client

}

func CheckConnection() int {
	err := mongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
