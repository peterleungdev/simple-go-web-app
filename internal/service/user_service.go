package service

import "webapp/internal/model"

type UserService struct {
	*Service
}

func (s *UserService) GetUserList() (list []*model.User, err error) {
	return
}

func (s *UserService) DeleteUser(id int64) (err error) {

	return
}
