// Code generated by ogen, DO NOT EDIT.

package petstore

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddPet implements addPet operation.
	//
	// Add new pet to the store inventory.
	//
	// POST /pet
	AddPet(ctx context.Context, req *Pet, params AddPetParams) (*Error, error)
	// CreateUser implements createUser operation.
	//
	// This can only be done by the logged in user.
	//
	// POST /user
	CreateUser(ctx context.Context, req *User) (*CreateUserDef, error)
	// CreateUsersWithArrayInput implements createUsersWithArrayInput operation.
	//
	// Creates list of users with given input array.
	//
	// POST /user/createWithArray
	CreateUsersWithArrayInput(ctx context.Context, req []User) (*CreateUsersWithArrayInputDef, error)
	// CreateUsersWithListInput implements createUsersWithListInput operation.
	//
	// Creates list of users with given input array.
	//
	// POST /user/createWithList
	CreateUsersWithListInput(ctx context.Context, req []User) (*CreateUsersWithListInputDef, error)
	// DeleteOrder implements deleteOrder operation.
	//
	// For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will
	// generate API errors.
	//
	// DELETE /store/order/{orderId}
	DeleteOrder(ctx context.Context, params DeleteOrderParams) (DeleteOrderRes, error)
	// DeletePet implements deletePet operation.
	//
	// Deletes a pet.
	//
	// DELETE /pet/{petId}
	DeletePet(ctx context.Context, params DeletePetParams) (*Error, error)
	// DeleteUser implements deleteUser operation.
	//
	// This can only be done by the logged in user.
	//
	// DELETE /user/{username}
	DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error)
	// FindPetsByStatus implements findPetsByStatus operation.
	//
	// Multiple status values can be provided with comma separated strings.
	//
	// GET /pet/findByStatus
	FindPetsByStatus(ctx context.Context, params FindPetsByStatusParams) (FindPetsByStatusRes, error)
	// FindPetsByTags implements findPetsByTags operation.
	//
	// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
	//
	// Deprecated: schema marks this operation as deprecated.
	//
	// GET /pet/findByTags
	FindPetsByTags(ctx context.Context, params FindPetsByTagsParams) (FindPetsByTagsRes, error)
	// GetInventory implements getInventory operation.
	//
	// Returns a map of status codes to quantities.
	//
	// GET /store/inventory
	GetInventory(ctx context.Context) (GetInventoryOK, error)
	// GetOrderById implements getOrderById operation.
	//
	// For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions.
	//
	// GET /store/order/{orderId}
	GetOrderById(ctx context.Context, params GetOrderByIdParams) (GetOrderByIdRes, error)
	// GetPetById implements getPetById operation.
	//
	// Returns a single pet.
	//
	// GET /pet/{petId}
	GetPetById(ctx context.Context, params GetPetByIdParams) (GetPetByIdRes, error)
	// GetUserByName implements getUserByName operation.
	//
	// Get user by user name.
	//
	// GET /user/{username}
	GetUserByName(ctx context.Context, params GetUserByNameParams) (GetUserByNameRes, error)
	// LoginUser implements loginUser operation.
	//
	// Logs user into the system.
	//
	// GET /user/login
	LoginUser(ctx context.Context, params LoginUserParams) (LoginUserRes, error)
	// LogoutUser implements logoutUser operation.
	//
	// Logs out current logged in user session.
	//
	// GET /user/logout
	LogoutUser(ctx context.Context) (*LogoutUserDef, error)
	// StoreSubscribePost implements POST /store/subscribe operation.
	//
	// Add subscription for a store events.
	//
	// POST /store/subscribe
	StoreSubscribePost(ctx context.Context, req OptStoreSubscribePostReq) (*StoreSubscribePostCreated, error)
	// UpdatePet implements updatePet operation.
	//
	// Update an existing pet.
	//
	// PUT /pet
	UpdatePet(ctx context.Context, req *Pet, params UpdatePetParams) (UpdatePetRes, error)
	// UpdatePetWithForm implements updatePetWithForm operation.
	//
	// Updates a pet in the store with form data.
	//
	// POST /pet/{petId}
	UpdatePetWithForm(ctx context.Context, req OptUpdatePetWithFormReq, params UpdatePetWithFormParams) (*Error, error)
	// UpdateUser implements updateUser operation.
	//
	// This can only be done by the logged in user.
	//
	// PUT /user/{username}
	UpdateUser(ctx context.Context, req *User, params UpdateUserParams) (UpdateUserRes, error)
	// UploadFile implements uploadFile operation.
	//
	// Uploads an image.
	//
	// POST /pet/{petId}/uploadImage
	UploadFile(ctx context.Context, req UploadFileReq, params UploadFileParams) (*ApiResponse, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
