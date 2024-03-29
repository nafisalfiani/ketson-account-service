package user

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/ketson-account-service/entity"
	"github.com/nafisalfiani/ketson-account-service/lib/auth"
	"github.com/nafisalfiani/ketson-account-service/lib/log"
	"github.com/nafisalfiani/ketson-account-service/usecase/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcUser struct {
	log  log.Interface
	user user.Interface
	auth auth.Interface
}

func Init(log log.Interface, user user.Interface, auth auth.Interface, validator *validator.Validate) UserServiceServer {
	return &grpcUser{
		log:  log,
		user: user,
		auth: auth,
	}
}

func (u *grpcUser) mustEmbedUnimplementedUserServiceServer() {}

func (u *grpcUser) GetUser(ctx context.Context, req *User) (*User, error) {
	user, err := u.user.Get(ctx, fromProto(req))
	if err != nil {
		return nil, err
	}

	return toProto(user), nil
}

func (u *grpcUser) CreateUser(ctx context.Context, req *User) (*User, error) {
	newUser, err := u.user.Create(ctx, fromProto(req))
	if err != nil {
		return nil, err
	}

	return toProto(newUser), nil
}

func (u *grpcUser) UpdateUser(ctx context.Context, req *User) (*User, error) {
	user, err := u.user.Update(ctx, fromProto(req))
	if err != nil {
		return nil, err
	}

	return toProto(user), nil
}

func (u *grpcUser) DeleteUser(ctx context.Context, req *User) (*emptypb.Empty, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	if err := u.user.Delete(ctx, entity.User{
		Id: id,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *grpcUser) GetUsers(ctx context.Context, in *emptypb.Empty) (*UserList, error) {
	users, err := u.user.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &UserList{}
	for i := range users {
		res.Users = append(res.Users, toProto(users[i]))
	}

	return res, nil
}

func (u *grpcUser) VerifyUserEmail(ctx context.Context, in *User) (*User, error) {
	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.user.Update(ctx, entity.User{
		Id:              id,
		IsEmailVerified: true,
		UpdatedAt:       time.Now(),
		UpdatedBy:       id.Hex(),
	})
	if err != nil {
		return nil, err
	}

	return toProto(updatedUser), nil
}
