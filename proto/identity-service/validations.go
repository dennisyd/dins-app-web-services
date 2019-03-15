package identity_service

import (
	fmt "fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

// Validate validates the user before saving to the db
func (u User) Validate(db *gorm.DB) {
	err := validation.Errors{
		"email":      validation.Validate(u.Email, validation.Required, is.Email),
		"first_name": validation.Validate(u.FirstName, validation.Required, validation.Length(3, 50)),
		"last_name":  validation.Validate(u.LastName, validation.Required, validation.Length(3, 50)),
		"password":   validation.Validate(u.Password, validation.Required, validation.Length(6, 50)),
	}
	for k, v := range err {
		if v != nil {
			db.AddError(fmt.Errorf("%s: %s", k, v.Error()))
		}
	}
}
