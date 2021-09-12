package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config is data required to connect to the database.
type Config struct {
	Host         string
	Port         string
	DatabaseName string
	Username     string
	Password     string
}

var _client *mongo.Client
var _database *mongo.Database
var _ctx = context.TODO()

// InitialiseDatabase initializes the connection to the database.
func InitialiseDatabase(config Config) {
	fmt.Println("Connecting to database.")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(_ctx, clientOptions)
	if err != nil {
		fmt.Println("Println to connect to database.", err)
		panic(err)
	}

	fmt.Println("Connected to database.")

	_client = client
	_database = client.Database(config.DatabaseName)
}

// Ping is used to check if the database is still connected to the app.
func Ping() bool {
	err := _client.Ping(_ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Failed to ping database.", err)
		return false
	}

	return true
}

// Insert adds a new entry to the database.
func Insert(collectionName string, object interface{}) (bool, error) {
	collection := _database.Collection(collectionName)

	ok, err := collection.InsertOne(_ctx, object)
	if err != nil {
		return false, err
	}

	var inserted = true
	if ok.InsertedID == nil {
		fmt.Println("No elements inserted.")
		inserted = false
	}

	return inserted, nil
}

// InsertMultiple adds multiple entries to the database at once.
func InsertMultiple(collectionName string, object []interface{}) error {
	collection := _database.Collection(collectionName)

	_, err := collection.InsertMany(_ctx, object)
	if err != nil {
		return err
	}

	return nil
}

// Get retrieves an entry from the database.
func Get(collectionName string, filter interface{}, model interface{}) error {
	collection := _database.Collection(collectionName)
	fmt.Println(collection)
	encodedFilter, marshalErr := bson.Marshal(filter)
	if marshalErr != nil {
		return marshalErr
	}
	fmt.Println(encodedFilter)

	err := collection.FindOne(_ctx, encodedFilter).Decode(model)

	return err
}

// GetAll entries from the database.
func GetAll(collectionName string, model interface{}) error {
	collection := _database.Collection(collectionName)

	cursor, err := collection.Find(_ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get object", err)
	}

	if err = cursor.All(_ctx, model); err != nil {
		fmt.Println("Failed to transform object", err)
	}

	return err
}

// Count entries from the database.
func Count(collectionName string) (int64, error) {
	collection := _database.Collection(collectionName)

	count, err := collection.CountDocuments(_ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get object", err)
	}

	return count, err
}

// Delete removes an entry from the database.
func Delete(collectionName string, filter interface{}) (bool, error) {
	collection := _database.Collection(collectionName)

	ok, err := collection.DeleteOne(_ctx, filter)
	if err != nil {
		return false, err
	}

	var deleted = true
	if ok.DeletedCount == 0 {
		fmt.Println("No elements deleted.")
		deleted = false
	}

	return deleted, nil
}

// RemoveCollection removes a collection from the database.
func RemoveCollection(collectionName string) {
	err := _database.Collection(collectionName).Drop(_ctx)
	if err != nil {
		panic(err)
	}
}

// UpsertEntry updates an existing entry in the database.
func UpsertEntry(collectionName string, filter interface{}, update interface{}) (bool, error) {
	updated, err := modifyEntry(collectionName, filter, update, "$set")

	if err != nil {
		fmt.Println("Failed to upsert entry", err)
	}
	return updated, nil
}

// RemoveEntry updates an existing entry in the database.
func RemoveEntry(collectionName string, filter interface{}, update interface{}) (bool, error) {
	updated, err := modifyEntry(collectionName, filter, update, "$unset")

	if err != nil {
		fmt.Println("Failed to update entry", err)
		return false, err
	}
	return updated, nil
}

// AppendToEntry appends an new entry to an array in the database.
func AppendToEntry(collectionName string, filter interface{}, add interface{}) (bool, error) {
	updated, err := modifyEntry(collectionName, filter, add, "$push")

	if err != nil {
		fmt.Println("Failed to append to entry", err)
	}
	return updated, err
}

// RemoveFromEntry appends an new entry to an array in the database.
func RemoveFromEntry(collectionName string, filter interface{}, remove interface{}) (bool, error) {
	updated, err := modifyEntry(collectionName, filter, remove, "$pull")

	if err != nil {
		fmt.Println("Failed to remove from entry", err)
	}
	return updated, err
}

func modifyEntry(collectionName string, filter interface{}, modify interface{}, operation string) (bool, error) {
	collection := _database.Collection(collectionName)

	ok, err := collection.UpdateOne(_ctx, filter, bson.M{operation: modify})
	if err != nil {
		return false, err
	}

	var updated = false
	if ok.ModifiedCount > 0 {
		updated = true
	}
	return updated, nil
}
