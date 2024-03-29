package storage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	client   *mongo.Client
	database *mongo.Database
}

func MongoStorage() Storage {
	uri := os.Getenv("MONGO_CONNECTION_URI")
	fmt.Println(uri)
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

func (storage *mongoStorage) StoreBlob(collectionName string, name string, content interface{}) error {
	collection := storage.database.Collection(collectionName)

	// filter := bson.D{{"name", name}}
	result, err := collection.InsertOne(context.Background(), content)
	// result, err := collection.UpdateOne(context.Background(), filter, content, options.Update().SetUpsert(true))
	if err != nil {
		fmt.Println(result)
		return err
	}

	return nil
}

func (storage *mongoStorage) ListBlobs(collectionName string, pattern string) ([]string, error) {
	// TODO: finish this
	return nil, nil
}

func (storage *mongoStorage) GetBlob(collectionName string, name string) ([]byte, error) {
	// TODO: finish this
	return nil, nil
}

func (storage *mongoStorage) GetAllBlobs(collectionName string, filter interface{}, results interface{}) error {
	collection := storage.database.Collection(collectionName)
	cursor, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), results)
	if err != nil {
		return err
	}

	return nil
}

func (storage *mongoStorage) DeleteBlob(collectionName string, name string) error {
	// TODO: finish this
	return nil
}
