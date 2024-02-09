package domain

import (
	"github.com/nafisalfiani/ketson-account-service/domain/role"
	"github.com/nafisalfiani/ketson-account-service/domain/user"
	"github.com/nafisalfiani/ketson-account-service/lib/broker"
	"github.com/nafisalfiani/ketson-account-service/lib/cache"
	"github.com/nafisalfiani/ketson-account-service/lib/log"
	"github.com/nafisalfiani/ketson-account-service/lib/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

type Domains struct {
	User user.Interface
	Role role.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Client, cache cache.Interface, broker broker.Interface) *Domains {
	return &Domains{
		User: user.Init(logger, json, db.Database("ketson_account_db").Collection("user"), cache, broker),
		Role: role.Init(logger, json, db.Database("ketson_account_db").Collection("role"), cache),
	}
}
