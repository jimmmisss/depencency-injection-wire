package user

import (
	"github.com/google/wire"
	"sync"
	"wire/domain"
)

var (
	hdl     *hanlder
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,
		// bind to interface

		wire.Bind(new(domain.UserHandler), new(*hanlder)),
		wire.Bind(new(domain.UserService), new(*service)),
		wire.Bind(new(domain.UserRepository), new(*repository)),
	)
)

func ProvideHandler(svc domain.UserService) *hanlder {
	hdlOnce.Do(func() {
		hdl = &hanlder{
			svc: svc,
		}
	})
	return hdl
}

func ProvideService(repo domain.UserRepository) *service {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})
	return svc
}

func ProvideRepository() *repository {
	repoOnce.Do(func() {
		repo = &repository{}
	})
	return repo
}
