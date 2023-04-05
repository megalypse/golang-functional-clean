package repository

import (
	"github.com/megalypse/golang-functional-clean/internal/data/repository/user"
	repoImpl "github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo/user"
)

var findUserByIdRepository user.FindUserByIdRepository

func init() {
	findUserByIdRepository = repoImpl.MongoFindUserByIdRepository()
}

func GetFindUserByIdRepository() user.FindUserByIdRepository {
	return findUserByIdRepository
}
