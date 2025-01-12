// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"database/sql"
)

// Injectors from wire.go:

func Wire(db *sql.DB) *hanlder {
	userRepository := ProvideRepository()
	userService := ProvideService(userRepository)
	userHanlder := ProvideHandler(userService)
	return userHanlder
}
