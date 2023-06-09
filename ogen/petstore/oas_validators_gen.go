// Code generated by ogen, DO NOT EDIT.

package petstore

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Address) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.City.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.City.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "city",
			Error: err,
		})
	}
	if err := func() error {
		if s.Country.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Country.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "country",
			Error: err,
		})
	}
	if err := func() error {
		if s.Street.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Street.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "street",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Category) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Name.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Name.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *DeleteOrderBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *DeleteOrderNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *DeleteUserBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *DeleteUserNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *Error) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           400,
			MaxSet:        true,
			Max:           599,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.Code)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s FindPetsByStatusOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s FindPetsByStatusStatusItem) Validate() error {
	switch s {
	case "available":
		return nil
	case "pending":
		return nil
	case "sold":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s FindPetsByTagsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetInventoryOK) Validate() error {
	var failures []validate.FieldError

	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetOrderByIdBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *GetOrderByIdNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *GetPetByIdBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *GetPetByIdNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *GetUserByNameBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *GetUserByNameNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *Order) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Quantity.Set {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(s.Quantity.Value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "quantity",
			Error: err,
		})
	}
	if err := func() error {
		if s.Status.Set {
			if err := func() error {
				if err := s.Status.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s OrderStatus) Validate() error {
	switch s {
	case "placed":
		return nil
	case "approved":
		return nil
	case "delivered":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *Pet) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Category.Set {
			if err := func() error {
				if err := s.Category.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "category",
			Error: err,
		})
	}
	if err := func() error {
		if s.PhotoUrls == nil {
			return errors.New("nil is invalid value")
		}
		if err := (validate.Array{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    20,
			MaxLengthSet: true,
		}).ValidateLength(len(s.PhotoUrls)); err != nil {
			return errors.Wrap(err, "array")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photoUrls",
			Error: err,
		})
	}
	if err := func() error {
		if s.Friend == nil {
			return nil // optional
		}
		if err := func() error {
			if err := s.Friend.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "pointer")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "friend",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Array{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    0,
			MaxLengthSet: false,
		}).ValidateLength(len(s.Tags)); err != nil {
			return errors.Wrap(err, "array")
		}
		var failures []validate.FieldError
		for i, elem := range s.Tags {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tags",
			Error: err,
		})
	}
	if err := func() error {
		if s.Status.Set {
			if err := func() error {
				if err := s.Status.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PetStatus) Validate() error {
	switch s {
	case "available":
		return nil
	case "pending":
		return nil
	case "sold":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *StoreSubscribePostReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.EventName.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "eventName",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s StoreSubscribePostReqEventName) Validate() error {
	switch s {
	case "orderInProgress":
		return nil
	case "orderShipped":
		return nil
	case "orderDelivered":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *Tag) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Name.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Name.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *UpdatePetBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *UpdatePetMethodNotAllowed) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *UpdatePetNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *UpdateUserBadRequest) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *UpdateUserNotFound) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s *User) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Pet.Set {
			if err := func() error {
				if err := s.Pet.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "pet",
			Error: err,
		})
	}
	if err := func() error {
		if s.Username.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    4,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Username.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "username",
			Error: err,
		})
	}
	if err := func() error {
		if s.FirstName.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.FirstName.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "firstName",
			Error: err,
		})
	}
	if err := func() error {
		if s.LastName.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.LastName.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "lastName",
			Error: err,
		})
	}
	if err := func() error {
		if s.Email.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        true,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Email.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if s.Password.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    8,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["/(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])/"],
				}).Validate(string(s.Password.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if err := func() error {
		if s.Phone.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["/^\\+(?:[0-9]-?){6,14}[0-9]$/"],
				}).Validate(string(s.Phone.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "phone",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Array{
			MinLength:    0,
			MinLengthSet: true,
			MaxLength:    0,
			MaxLengthSet: false,
		}).ValidateLength(len(s.Addresses)); err != nil {
			return errors.Wrap(err, "array")
		}
		var failures []validate.FieldError
		for i, elem := range s.Addresses {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "addresses",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserPet) Validate() error {
	switch s.Type {
	case PetUserPet:
		if err := s.Pet.Validate(); err != nil {
			return err
		}
		return nil
	case TagUserPet:
		if err := s.Tag.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}
