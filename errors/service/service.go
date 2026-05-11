package service

import (
	"err/common"
	"err/repository"
	"fmt"
)

func GetUserService() (common.User, error) {
	user, err := repository.GetUserRepository()
	if err != nil {
		//return common.User{}, err
		return common.User{}, fmt.Errorf("repository error=>: %w", err)
	}
	return user, nil
}
