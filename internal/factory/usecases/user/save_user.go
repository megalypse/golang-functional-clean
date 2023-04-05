package user

import (
	caseImpl "github.com/megalypse/golang-functional-clean/internal/data/usecases/user"
	"github.com/megalypse/golang-functional-clean/internal/domain/usecases/user"
	repoFactory "github.com/megalypse/golang-functional-clean/internal/infra/repository/mongo/user"
)

var saveUserUsecase user.SaveUserUsecase

func init() {
	saveUserUsecase = caseImpl.SaveUser(
		repoFactory.MongoSaveUserRepository(),
	)
}

func GetSaveUserUsecase() user.SaveUserUsecase {
	return saveUserUsecase
}
