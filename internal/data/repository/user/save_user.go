package user

import (
	"context"

	"github.com/megalypse/golang-functional-clean/internal/domain/models"
)

type SaveUserRepository func(context.Context, models.User) error
