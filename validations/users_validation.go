package validations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/juanfer2/api-rest-go/models"
)

func ValidateCreateUser(t models.User) error {
	return validation.ValidateStruct(&t,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&t.Username, validation.Required),
		validation.Field(&t.Password, validation.Required),
	)
}
