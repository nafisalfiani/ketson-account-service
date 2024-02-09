package config

import (
	"github.com/nafisalfiani/ketson-account-service/handler/grpc"
	"github.com/nafisalfiani/ketson-account-service/lib/auth"
	"github.com/nafisalfiani/ketson-account-service/lib/broker"
	"github.com/nafisalfiani/ketson-account-service/lib/cache"
	"github.com/nafisalfiani/ketson-account-service/lib/log"
	"github.com/nafisalfiani/ketson-account-service/lib/nosql"
	"github.com/nafisalfiani/ketson-account-service/lib/security"
)

type Application struct {
	Auth     auth.Config     `env:"AUTH"`
	Log      log.Config      `env:"LOG"`
	Security security.Config `env:"SECURITY"`
	NoSql    nosql.Config    `env:"NO_SQL"`
	Cache    cache.Config    `env:"CACHE"`
	Broker   broker.Config   `env:"BROKER"`
	Grpc     grpc.Config     `env:"GRPC"`
}
