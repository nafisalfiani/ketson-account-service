package usecase

import (
	"github.com/nafisalfiani/ketson-account-service/domain"
	"github.com/nafisalfiani/ketson-account-service/usecase/role"
	"github.com/nafisalfiani/ketson-account-service/usecase/user"
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/log"
)

type Usecases struct {
	User user.Interface
	Role role.Interface
}

func Init(logger log.Interface, broker broker.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		User: user.Init(logger, dom.User, broker),
		Role: role.Init(logger, dom.Role),
	}
}
