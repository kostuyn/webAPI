package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"webApi/internal/apperror"
	"webApi/internal/user"
	"webApi/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %v", err)
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectID.Hex(), nil
	}

	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectId to hex")
}

func (d *db) FindAll(ctx context.Context) (users []user.User, err error) {
	cursor, err := d.collection.Find(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %v", err)
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to read all documents from cursor: %v", err)
	}

	return users, nil
}

func (d *db) FindOne(ctx context.Context, id string) (user user.User, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, fmt.Errorf("failed to convert hex to objectId: %v", err)
	}

	filter := primitive.M{"_id": objectID}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user, apperror.ErrNotFound

		}
		return user, fmt.Errorf("failed to find user by id: %v", err)
	}

	err = result.Decode(&user)
	if err != nil {
		return user, fmt.Errorf("failed to decode user: %v", err)
	}

	return user, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": user}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return apperror.ErrNotFound
	}

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return apperror.ErrNotFound
	}

	return nil
}

func NewStorage(database *mongo.Database, collectionName string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collectionName),
		logger:     logger,
	}
}
