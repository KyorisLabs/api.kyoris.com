package sharedApplication

import (
	authApplication "api.kyoris.com/core/auth/application"
	trainerApplication "api.kyoris.com/core/trainer/application"
	userApplication "api.kyoris.com/core/user/application"
)

type Services struct {
	UserService    userApplication.UserService
	AuthService    authApplication.AuthService
	TrainerService trainerApplication.TrainerService
}
