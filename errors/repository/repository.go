package repository

import (
	"err/common"
	"fmt"
)

func GetUserRepository() (common.User, error) {
	//return common.User{ID: 42, Name: "Vasya"}, nil
	//return common.User{}, common.ErNotFound
	//return common.User{}, common.ErUnknown
	//return common.User{}, fmt.Errorf("database error=>: %w", common.ErUnknown)
	return common.User{}, fmt.Errorf("database error=>: %w", common.ErBrokenConnection)
}
