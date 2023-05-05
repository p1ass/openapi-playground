package main

import (
	"github.com/p1ass/openapi-playground/oapi-codegen/petstore_chi"
	"net/http"
)

type petService struct {
}

var _ petstore_chi.ServerInterface = &petService{}

func (p *petService) AddPet(w http.ResponseWriter, r *http.Request, params petstore_chi.AddPetParams) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdatePet(w http.ResponseWriter, r *http.Request, params petstore_chi.UpdatePetParams) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) FindPetsByStatus(w http.ResponseWriter, r *http.Request, params petstore_chi.FindPetsByStatusParams) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) FindPetsByTags(w http.ResponseWriter, r *http.Request, params petstore_chi.FindPetsByTagsParams) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeletePet(w http.ResponseWriter, r *http.Request, petId int64, params petstore_chi.DeletePetParams) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetPetById(w http.ResponseWriter, r *http.Request, petId int64) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdatePetWithForm(w http.ResponseWriter, r *http.Request, petId int64) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UploadFile(w http.ResponseWriter, r *http.Request, petId int64) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetInventory(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeleteOrder(w http.ResponseWriter, r *http.Request, orderId string) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetOrderById(w http.ResponseWriter, r *http.Request, orderId int64) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) PostStoreSubscribe(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUsersWithArrayInput(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) LoginUser(w http.ResponseWriter, r *http.Request, params petstore_chi.LoginUserParams) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) LogoutUser(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) DeleteUser(w http.ResponseWriter, r *http.Request, username string) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) GetUserByName(w http.ResponseWriter, r *http.Request, username string) {
	// TODO implement me
	panic("implement me")
}

func (p *petService) UpdateUser(w http.ResponseWriter, r *http.Request, username string) {
	// TODO implement me
	panic("implement me")
}
