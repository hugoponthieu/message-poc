package mongo

import (
	"context"
	"message/repository/types"
	"message/types/message"

	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepository struct {
	messageCollection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) *MongoRepository {
	messageCollection := database.Collection("messages")
	return &MongoRepository{
		messageCollection: messageCollection,
	}
}

func (r MongoRepository) Get(id string) (*message.Message, error) {
	var message message.Message
	err := r.messageCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r MongoRepository) MGet(ids []string) (*types.MGetResult, error) {
	var res types.MGetResult
	cursor, err := r.messageCollection.Find(context.Background(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var message message.Message
		err := cursor.Decode(&message)
		if err != nil {
			return nil, err
		}
		res.Messages = append(res.Messages, &message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r MongoRepository) Create(message *message.Message) (*message.Message, error) {
	result, err := r.messageCollection.InsertOne(context.Background(), message)
	if err != nil {
		return nil, err
	}
	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		hexID := oid.Hex()
		message.ID = &hexID
	}
	return message, nil
}

func (r MongoRepository) Update(id string, msg *message.UpdateMessage) (*message.Message, error) {
	var message message.Message
	err := r.messageCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": msg}).Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r MongoRepository) Delete(id string) error {
	_, err := r.messageCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r MongoRepository) GetByChannel(channelId string, page int, limit int) (*[]message.Message, error) {
	var messages []message.Message
	pagination_opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	cursor, err := r.messageCollection.Find(context.Background(), bson.M{"channel_id": channelId}, pagination_opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var message message.Message
		err := cursor.Decode(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &messages, nil
}

func (r MongoRepository) Search(query string, channelId *string, serverId *string, page int, limit int) (*[]message.Message, error) {
	var messages []message.Message
	pagination_opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	query_filter := bson.M{"$text": bson.M{"$search": query}}
	if channelId != nil {
		query_filter["channel_id"] = *channelId
	}
	if serverId != nil {
		query_filter["server_id"] = *serverId
	}
	cursor, err := r.messageCollection.Find(context.Background(), query_filter, pagination_opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var message message.Message
		err := cursor.Decode(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &messages, nil
}
