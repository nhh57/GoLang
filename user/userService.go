package user

import "log"

type UserService interface {
	FindAll() ([]User, error)
	FindById(Id int) (User, error)
	Create(user UserRequest) (User, error)
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) FindAll() ([]User, error) {
	users, err := s.userRepository.FindAll()
	return users, err
}

func (s *userService) FindById(Id int) (User, error) {
	user, err := s.userRepository.FindById(Id)
	return user, err
}

func (s *userService) Create(userRequest UserRequest) (User, error) {
	log.Default().Printf("==========Start Create User Serv==========")
	users := User{
		Name:       userRequest.Name,
		Email:      userRequest.Email,
		UserName:   userRequest.UserName,
		Phone:      userRequest.Phone,
		DayOfBirth: userRequest.DayOfBirth,
		PassWord:   userRequest.PassWord,
	}
	log.Default().Printf("data:", users)
	users, err := s.userRepository.Create(users)
	log.Default().Printf("==========End Create User Serv==========")
	return users, err
}
