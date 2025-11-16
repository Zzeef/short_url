package link

import "short_url/internal/storage"

type LinkRepo struct {
	store *storage.Mongo
}

func NewRepository(store *storage.Mongo) *LinkRepo {
	return &LinkRepo{store: store}
}
