package repository

import (
	"github.com/megalypse/golang-functional-clean/internal/data/repository/user"
	repoImpl "github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo/user"
)

var saveUserRepository user.SaveUserRepository

func init() {
	saveUserRepository = repoImpl.MongoSaveUserRepository()
}

func GetSaveUserRepository() user.SaveUserRepository {
	return saveUserRepository
}
