package user

import (
	"context"

	"github.com/megalypse/golang-functional-clean/internal/data/repository/user"
	"github.com/megalypse/golang-functional-clean/internal/domain/models"
	"github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo"
	"github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo/collections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MongoFindUserByIdRepository() user.FindUserByIdRepository {
	return func(ctx context.Context, id string) (models.User, error) {
		collection := mongo.GetCollection(mongo.GetSession(ctx), collections.CollectionUsers)

		result := collection.FindOne(ctx, bson.D{primitive.E{Key: "id", Value: id}})

		userDbe := UserDbe{}
		err := result.Decode(&userDbe)

		return userDbe.ToEntity(), err
	}
}
