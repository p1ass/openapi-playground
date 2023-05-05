package main

import (
	"context"
	"github.com/p1ass/openapi-playground/ogen/petstore"
)

type petService struct {
}

var _ petstore.Handler = &petService{}

func (p *petService) AddPet(ctx context.Context, req *petstore.Pet, params petstore.AddPetParams) (*petstore.Error, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUser(ctx context.Context, req *petstore.User) (*petstore.CreateUserDef, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUsersWithArrayInput(ctx context.Context, req []petstore.User) (*petstore.CreateUsersWithArrayInputDef, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUsersWithListInput(ctx context.Context, req []petstore.User) (*petstore.CreateUsersWithListInputDef, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeleteOrder(ctx context.Context, params petstore.DeleteOrderParams) (petstore.DeleteOrderRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeletePet(ctx context.Context, params petstore.DeletePetParams) (*petstore.Error, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeleteUser(ctx context.Context, params petstore.DeleteUserParams) (petstore.DeleteUserRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) FindPetsByStatus(ctx context.Context, params petstore.FindPetsByStatusParams) (petstore.FindPetsByStatusRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) FindPetsByTags(ctx context.Context, params petstore.FindPetsByTagsParams) (petstore.FindPetsByTagsRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetInventory(ctx context.Context) (petstore.GetInventoryOK, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetOrderById(ctx context.Context, params petstore.GetOrderByIdParams) (petstore.GetOrderByIdRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetUserByName(ctx context.Context, params petstore.GetUserByNameParams) (petstore.GetUserByNameRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) LoginUser(ctx context.Context, params petstore.LoginUserParams) (petstore.LoginUserRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) LogoutUser(ctx context.Context) (*petstore.LogoutUserDef, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) StoreSubscribePost(ctx context.Context, req petstore.OptStoreSubscribePostReq) (*petstore.StoreSubscribePostCreated, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdatePet(ctx context.Context, req *petstore.Pet, params petstore.UpdatePetParams) (petstore.UpdatePetRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdatePetWithForm(ctx context.Context, req petstore.OptUpdatePetWithFormReq, params petstore.UpdatePetWithFormParams) (*petstore.Error, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdateUser(ctx context.Context, req *petstore.User, params petstore.UpdateUserParams) (petstore.UpdateUserRes, error) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UploadFile(ctx context.Context, req petstore.UploadFileReq, params petstore.UploadFileParams) (*petstore.ApiResponse, error) {
	// TODO implement me
	panic("implement me")
}
