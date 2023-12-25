package main

import (
	database "alert-service/database/sqlc"
	"context"
)

type Alerter interface {
	// creates an alert and pushes to postgres and redis for indexing it in a sorted set
	Create(context.Context) error

	// Get all your alerts from postgres. Can filter them too
	Read(context.Context) error

	// Update an alert in postgres and redis
	Update(context.Context) error

	// Delete an alert from postgres and redis
	Delete(context.Context) error
}

type alert struct{
	cache Cacher
	db database.Querier
}

func NewAlertService(cache Cacher, db database.Querier) Alerter {
	return &alert{
		cache: cache,
		db: db,
	}
}

func (a *alert) Create(ctx context.Context) error {
	return nil
}

func (a *alert) Read(ctx context.Context) error {
	return nil
}

func (a *alert) Update(ctx context.Context) error {
	return nil
}

func (a *alert) Delete(ctx context.Context) error {
	return nil
}
