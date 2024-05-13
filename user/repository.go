package user

import (
	"context"
	"database/sql"
	"wire/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUserName(ctx context.Context, username string) (*domain.UserEntity, error) {

	userEntity := domain.UserEntity{
		ID:       "1",
		UserName: "test",
	}

	return &userEntity, nil
}
