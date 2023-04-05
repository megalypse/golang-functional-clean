package user

import (
	"context"

	repository "github.com/megalypse/golang-functional-clean/internal/data/repository/user"
	"github.com/megalypse/golang-functional-clean/internal/domain/models"
	"github.com/megalypse/golang-functional-clean/internal/domain/usecases/user"
	"github.com/megalypse/golang-functional-clean/internal/factory/database"
)

func SaveUser(
	saveUserRepository repository.SaveUserRepository,
) user.SaveUserUsecase {
	return func(ctx context.Context, u models.User) error {
		defer database.CloseDatabaseSession(ctx)

		err := saveUserRepository(ctx, u)
		return err
	}
}
