package user

import (
	"context"

	"github.com/megalypse/golang-functional-clean/internal/domain/models"
)

type FindUserByIdRepository func(context.Context, string) (models.User, error)
