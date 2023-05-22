package repository

import (
	"context"
	"log"

	. "alerts/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

type AlertsRepository struct {
	Server       string
	DatabaseName string
	Client       *mongo.Client
	Database     *mongo.Database
}

const (
	COLLECTION = "alerts"
)

func (m *AlertsRepository) Connect() {
	clientOptions := options.Client().ApplyURI(m.Server)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	m.Client = client
	m.Database = client.Database(m.DatabaseName)
}

func (m *AlertsRepository) FindAll() ([]Alert, error) {
	collection := m.Database.Collection(COLLECTION)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var alerts []Alert
	err = cursor.All(context.Background(), &alerts)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}

func (m *AlertsRepository) FindById(id string) (Alert, error) {
	collection := m.Database.Collection(COLLECTION)
	obId, _ := primitive.ObjectIDFromHex(id)

	var alert Alert
	err := collection.FindOne(context.Background(), bson.M{"_id": obId}).Decode(&alert)
	if err != nil {
		return Alert{}, err
	}

	return alert, nil
}

func (m *AlertsRepository) Insert(alert Alert) error {
	collection := m.Database.Collection(COLLECTION)

	_, err := collection.InsertOne(context.Background(), alert)
	if err != nil {
		return err
	}

	return nil
}

func (m *AlertsRepository) Delete(alert Alert) error {
	collection := m.Database.Collection(COLLECTION)

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": alert.ID})
	if err != nil {
		return err
	}

	return nil
}

func (m *AlertsRepository) Update(alert Alert) error {
	collection := m.Database.Collection(COLLECTION)

	update := bson.M{"$set": alert}

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": alert.ID}, update)
	if err != nil {
		return err
	}

	return nil
}
