package helpers

import (
	"boilerplate/database"
)

// Shortcut to quickly delete a model.
// Takes the model type you want to delete & id of the record you want to delete.
func DeleteModel(model interface{}, id string) error {
	var err error
	err = database.Client.First(&model, id).Error
	if err != nil {
		return err
	}

	err = database.Client.Delete(&model).Error

	if err != nil {
		return err
	}

	return nil
}

// Shortcut to quickly create a model.
// Takes the model you want to create
// Provide a reference to your model, so you can access gorm model details after
func CreateModel(model interface{}) error {
	tx := database.Client.Begin()

	if err := tx.Create(model).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
