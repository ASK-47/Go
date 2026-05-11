package transport

import (
	"err/common"
	"err/service"
	"fmt"
)

func GetUserTransport() (common.User, error) {
	user, err := service.GetUserService()
	if err != nil {
		//return common.User{}, err
		return common.User{}, fmt.Errorf("service error=>: %w", err)

	}
	return user, nil
}
