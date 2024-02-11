package storage

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func MongoStorage() Storage {
	uri := os.Getenv("MONGO_CONNECTION_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	database := client.Database(os.Getenv("MONGO_DATABASE_NAME"))
	collection := database.Collection(os.Getenv("MONGO_COLLECTION_NAME"))

	return &mongoStorage{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (this *mongoStorage) StoreBlob(name string, content []byte) error {
	this.collection.InsertOne(context.Background(), content)

	return nil
}

func (this *mongoStorage) ListBlobs(pattern string) ([]string, error) {
	return nil, nil
}

func (this *mongoStorage) GetBlob(name string) ([]byte, error) {
	return nil, nil
}

func (this *mongoStorage) GetAllBlobs(pattern string) ([][]byte, error) {
	return nil, nil
}

func (this *mongoStorage) DeleteBlob(name string) error {
	return nil
}
