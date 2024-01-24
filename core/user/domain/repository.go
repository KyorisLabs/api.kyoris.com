package userDomain

type UserRepository interface {
	Create(user *UserDTO) error
	FindByEmail(email string) *UserDTO
	FindById(id uint) *UserDTO
	FindByUsername(username string) *UserDTO
}
