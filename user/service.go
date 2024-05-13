package user

import (
	"context"
	"wire/domain"
)

type service struct {
	repo domain.UserRepository
}

func (s *service) FetchByUserName(ctx context.Context, username string) (*domain.User, error) {
	user, err := s.repo.FetchByUserName(ctx, username)
	return &domain.User{
		ID:       user.ID,
		UserName: user.UserName,
	}, err
}
