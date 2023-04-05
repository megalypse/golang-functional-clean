package database

import (
	"context"

	"github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo"
)

func CloseDatabaseSession(ctx context.Context) {
	mongo.CloseSession(ctx)
}
