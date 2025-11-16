package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	DB     *mongo.Database
	ctx    context.Context
	cancel context.CancelFunc
}

func NewMongo(uri, dbName string) *Mongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return &Mongo{
		DB:     client.Database(dbName),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (m *Mongo) Close() {
	defer m.cancel()
	if err := m.DB.Client().Disconnect(m.ctx); err != nil {
		log.Fatal(err)
	}
}
