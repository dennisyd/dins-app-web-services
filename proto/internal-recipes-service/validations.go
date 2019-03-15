package internal_recipes_service

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func (recipe Recipe) Validate(db *gorm.DB) {
	if len(recipe.Name) < 3 {
		db.AddError(errors.New("name must be specified and greater than 3 characters"))
	}
	if len(recipe.Description) < 10 {
		db.AddError(errors.New("description must be specified and greater than 10 characters"))
	}
}
