package main

import (
	"context"
	"github.com/p1ass/openapi-playground/ogen/petstore"
)

type petService struct {
}

func (s *petService) AddPet(ctx context.Context, req *petstore.Pet, params petstore.AddPetParams) error {
	// TODO implement me
	panic("implement me")
}

func (s *petService) CreateUser(ctx context.Context, req *petstore.User) (*petstore.CreateUserDef, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) CreateUsersWithArrayInput(ctx context.Context, req []petstore.User) (*petstore.CreateUsersWithArrayInputDef, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) CreateUsersWithListInput(ctx context.Context, req []petstore.User) (*petstore.CreateUsersWithListInputDef, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) DeleteOrder(ctx context.Context, params petstore.DeleteOrderParams) (petstore.DeleteOrderRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) DeletePet(ctx context.Context, params petstore.DeletePetParams) error {
	// TODO implement me
	panic("implement me")
}

func (s *petService) DeleteUser(ctx context.Context, params petstore.DeleteUserParams) (petstore.DeleteUserRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) FindPetsByStatus(ctx context.Context, params petstore.FindPetsByStatusParams) (petstore.FindPetsByStatusRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) FindPetsByTags(ctx context.Context, params petstore.FindPetsByTagsParams) (petstore.FindPetsByTagsRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) GetInventory(ctx context.Context) (petstore.GetInventoryOK, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) GetOrderById(ctx context.Context, params petstore.GetOrderByIdParams) (petstore.GetOrderByIdRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) GetUserByName(ctx context.Context, params petstore.GetUserByNameParams) (petstore.GetUserByNameRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) LoginUser(ctx context.Context, params petstore.LoginUserParams) (petstore.LoginUserRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) LogoutUser(ctx context.Context) (*petstore.LogoutUserDef, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) StoreSubscribePost(ctx context.Context, req petstore.OptStoreSubscribePostReq) (*petstore.StoreSubscribePostCreated, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) UpdatePet(ctx context.Context, req *petstore.Pet, params petstore.UpdatePetParams) (petstore.UpdatePetRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) UpdatePetWithForm(ctx context.Context, req petstore.OptUpdatePetWithFormReq, params petstore.UpdatePetWithFormParams) error {
	// TODO implement me
	panic("implement me")
}

func (s *petService) UpdateUser(ctx context.Context, req *petstore.User, params petstore.UpdateUserParams) (petstore.UpdateUserRes, error) {
	// TODO implement me
	panic("implement me")
}

func (s *petService) UploadFile(ctx context.Context, req petstore.UploadFileReq, params petstore.UploadFileParams) (*petstore.ApiResponse, error) {
	// TODO implement me
	panic("implement me")
}

var _ petstore.Handler = &petService{}
