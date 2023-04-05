package user

import (
	"context"

	repository "github.com/megalypse/golang-functional-clean/internal/data/repository/user"
	"github.com/megalypse/golang-functional-clean/internal/domain/models"
	"github.com/megalypse/golang-functional-clean/internal/domain/usecases/user"
	"github.com/megalypse/golang-functional-clean/internal/factory/database"
)

func FindUserById(
	findByIdRepository repository.FindUserByIdRepository,
) user.FindUserByIdUsecase {
	return func(ctx context.Context, s string) (models.User, error) {
		defer database.CloseDatabaseSession(ctx)

		user, err := findByIdRepository(ctx, s)
		return user, err
	}
}
