package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	client   *mongo.Client
	database *mongo.Database
}

func MongoStorage() Storage {
	// uri := os.Getenv("MONGO_CONNECTION_URI")
	uri := "mongodb://localhost:27017"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	database := client.Database("tempam")

	return &mongoStorage{
		client:   client,
		database: database,
	}
}

func (this *mongoStorage) StoreBlob(collectionName string, name string, content interface{}) error {
	collection := this.database.Collection(collectionName)

	// filter := bson.D{{"name", name}}
	result, err := collection.InsertOne(context.Background(), content)
	// result, err := collection.UpdateOne(context.Background(), filter, content, options.Update().SetUpsert(true))
	if err != nil {
		fmt.Println(result)
		return err
	}

	return nil
}

func (this *mongoStorage) ListBlobs(collectionName string, pattern string) ([]string, error) {
	return nil, nil
}

func (this *mongoStorage) GetBlob(collectionName string, name string) ([]byte, error) {
	return nil, nil
}

func (this *mongoStorage) GetAllBlobs(collectionName string, pattern string) ([][]byte, error) {
	return nil, nil
}

func (this *mongoStorage) DeleteBlob(collectionName string, name string) error {
	return nil
}
