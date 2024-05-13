package domain

import (
	"context"
	"net/http"
)

type (
	User struct {
		ID       string `json:"id"`
		UserName string `json:"username"`
	}

	UserEntity struct {
		ID       string
		UserName string
		Password string
	}

	UserRepository interface {
		FetchByUserName(ctx context.Context, username string) (*UserEntity, error)
	}

	UserService interface {
		FetchByUserName(ctx context.Context, username string) (*User, error)
	}

	UserHandler interface {
		FetchByUserName() http.HandlerFunc
	}
)
