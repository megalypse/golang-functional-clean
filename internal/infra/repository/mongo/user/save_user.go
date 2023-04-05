package user

import (
	"context"

	"github.com/megalypse/golang-functional-clean/internal/data/repository/user"
	"github.com/megalypse/golang-functional-clean/internal/domain/models"
	"github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo"
)

func MongoSaveUserRepository() user.SaveUserRepository {
	return func(ctx context.Context, u models.User) error {
		usersCollection := mongo.GetCollection(mongo.GetSession(ctx), "users")

		_, err := usersCollection.InsertOne(ctx, u)
		return err
	}
}
