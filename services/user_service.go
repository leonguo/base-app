package services

import (
	//"errors"
	"echo-app/models"
)

// UserService handles CRUID operations of a user datamodel,
// it depends on a user repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type UserService interface {
	GetAll() []models.User
	GetByID(id int64) (models.User, bool)
	//GetByUsernameAndPassword(username, userPassword string) (models.User, bool)
	//DeleteByID(id int64) bool
	//
	//Update(id int64, user models.User) (models.User, error)
	//UpdatePassword(id int64, newPassword string) (models.User, error)
	//UpdateUsername(id int64, newUsername string) (models.User, error)
	//
	//Create(userPassword string, user models.User) (models.User, error)
}

// NewUserService returns the default user service.
//func NewUserService(repo repositories.UserRepository) UserService {
//	return &userService{
//		repo: repo,
//	}
//}
//
//type userService struct {
//	repo repositories.UserRepository
//}
//
//// GetAll returns all users.
//func (s *userService) GetAll() []models.User {
//	return s.repo.SelectMany(func(_ models.User) bool {
//		return true
//	}, -1)
//}
//
//// GetByID returns a user based on its id.
//func (s *userService) GetByID(id int64) (models.User, bool) {
//	return s.repo.Select(func(m models.User) bool {
//		return m.ID == id
//	})
//}