package validations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/juanfer2/api-rest-go/models"
)

func ValidateCreateTask(t models.Task) error {
	return validation.ValidateStruct(&t,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&t.Name, validation.Required, validation.Length(5, 50)),
	)
}
