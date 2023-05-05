package main

import (
	"github.com/labstack/echo/v4"
	"github.com/p1ass/openapi-playground/oapi-codegen/petstore_echo"
)

type petService struct {
}

var _ petstore_echo.ServerInterface = &petService{}

func (p *petService) AddPet(ctx echo.Context, params petstore_echo.AddPetParams) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdatePet(ctx echo.Context, params petstore_echo.UpdatePetParams) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) FindPetsByStatus(ctx echo.Context, params petstore_echo.FindPetsByStatusParams) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) FindPetsByTags(ctx echo.Context, params petstore_echo.FindPetsByTagsParams) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeletePet(ctx echo.Context, petId int64, params petstore_echo.DeletePetParams) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetPetById(ctx echo.Context, petId int64) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdatePetWithForm(ctx echo.Context, petId int64) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UploadFile(ctx echo.Context, petId int64) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetInventory(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) PlaceOrder(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeleteOrder(ctx echo.Context, orderId string) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetOrderById(ctx echo.Context, orderId int64) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) PostStoreSubscribe(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUser(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUsersWithArrayInput(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUsersWithListInput(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) LoginUser(ctx echo.Context, params petstore_echo.LoginUserParams) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) LogoutUser(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeleteUser(ctx echo.Context, username string) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetUserByName(ctx echo.Context, username string) error {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdateUser(ctx echo.Context, username string) error {
	// TODO implement me
	panic("implement me")
}
