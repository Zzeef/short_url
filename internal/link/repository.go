package link

import (
	"context"
	"errors"
	"short_url/internal/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkRepo struct {
	store *storage.Mongo
}

func NewRepository(store *storage.Mongo) *LinkRepo {
	return &LinkRepo{store: store}
}

func (r *LinkRepo) DeleteRecord(ctx context.Context, code string) error {
	collection := r.store.DB.Collection("url")

	filter := bson.M{"shortCode": code}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *LinkRepo) UpdateUrlByCode(ctx context.Context, url string, code string) error {
	collection := r.store.DB.Collection("url")

	filter := bson.M{"shortCode": code}
	update := bson.M{"$set": bson.M{"url": url, "updatedAt": time.Now()}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *LinkRepo) GetRecordByColumn(ctx context.Context, column string, value string) (*Link, error) {
	collection := r.store.DB.Collection("url")

	var result Link
	err := collection.FindOne(ctx, bson.M{column: value}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func fillDefaultFields(link *Link) {
	if link.CreatedAt.IsZero() {
		link.CreatedAt = time.Now()
	}
	if link.UpdatedAt.IsZero() {
		link.UpdatedAt = time.Now()
	}
}

func (r *LinkRepo) Insert(ctx context.Context, link *Link) error {
	collection := r.store.DB.Collection("url")

	fillDefaultFields(link)
	_, err := collection.InsertOne(ctx, link)
	return err
}
