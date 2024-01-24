package userDomain

import "errors"

var (
	// Mensaje de error para cuando el usuario ya existe.
	ErrUserAlreadyExists = errors.New("user_already_exists")
	// Mensaje de error para cuando no se ha podido crear el usuario.
	ErrCreateUser = errors.New("error_create_user")
	// Mensaje de error para cuando el nombre de usuario ya existe.
	ErrUsernameAlreadyExists = errors.New("username_already_exists")
)
