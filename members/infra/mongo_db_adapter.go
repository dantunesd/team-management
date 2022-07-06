package infra

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBAdapter struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewMongoDBAdapter(client *mongo.Client, dbName, collection string) *MongoDBAdapter {
	return &MongoDBAdapter{
		client:     client,
		dbName:     dbName,
		collection: collection,
	}
}

func (d *MongoDBAdapter) Get(fieldName, fieldValue string, output interface{}) error {
	return d.findOne(fieldName, fieldValue, output)
}

func (d *MongoDBAdapter) Create(output interface{}) (string, error) {
	return d.insertOne(output)
}

func (d *MongoDBAdapter) Update(fieldName, fieldValue string, output interface{}) (bool, error) {
	return d.updateOne(fieldName, fieldValue, output)
}

func (d *MongoDBAdapter) Delete(fieldName, fieldValue string) (bool, error) {
	return d.deleteOne(fieldName, fieldValue)
}

func (d *MongoDBAdapter) Filter(filters map[string]string, output interface{}) error {
	return d.find(filters, output)
}

func (d *MongoDBAdapter) findOne(fieldName, fieldValue string, output interface{}) error {
	collection := d.getCollection()
	filter := d.mountIdFilter(fieldName, fieldValue)
	err := collection.FindOne(context.TODO(), filter).Decode(output)

	if d.isNotFoundError(err) {
		return nil
	}
	return err
}

func (d *MongoDBAdapter) insertOne(output interface{}) (string, error) {
	collection := d.getCollection()
	result, err := collection.InsertOne(context.TODO(), output)
	return d.getIDFromResult(result, err)
}

func (d *MongoDBAdapter) updateOne(fieldName, fieldValue string, output interface{}) (bool, error) {
	collection := d.getCollection()
	filter := d.mountIdFilter(fieldName, fieldValue)
	update := bson.M{"$set": output}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	return result.ModifiedCount > 0, err
}

func (d *MongoDBAdapter) deleteOne(fieldName, fieldValue string) (bool, error) {
	collection := d.getCollection()
	filter := d.mountIdFilter(fieldName, fieldValue)
	result, err := collection.DeleteOne(context.TODO(), filter)
	return result.DeletedCount > 0, err
}

func (d *MongoDBAdapter) find(filters map[string]string, output interface{}) error {
	collection := d.getCollection()
	bsonFilters := d.mountBsonFilters(filters)
	cursor, err := collection.Find(context.TODO(), bsonFilters)
	if err != nil {
		return err
	}
	return cursor.All(context.TODO(), output)
}

func (d *MongoDBAdapter) getCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection(d.collection)
}

func (d *MongoDBAdapter) isNotFoundError(err error) bool {
	return err == nil || err == mongo.ErrNoDocuments
}

func (d *MongoDBAdapter) getIDFromResult(result *mongo.InsertOneResult, err error) (string, error) {
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), err
}

func (d *MongoDBAdapter) mountIdFilter(fieldName, fieldValue string) bson.M {
	objectId, _ := primitive.ObjectIDFromHex(fieldValue)
	return bson.M{fieldName: objectId}
}

func (d *MongoDBAdapter) mountBsonFilters(filters map[string]string) bson.D {
	bsonFilter := bson.D{}
	for field, value := range filters {
		bsonFilter = append(bsonFilter, bson.E{Key: field, Value: value})
	}
	return bsonFilter
}
