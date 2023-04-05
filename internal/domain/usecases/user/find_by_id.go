package user

import (
	"context"

	"github.com/megalypse/golang-functional-clean/internal/domain/models"
)

type FindUserByIdUsecase func(context.Context, string) (models.User, error)
