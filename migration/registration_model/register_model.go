package registrationmodel

import (
	"guestbook_backend/models"
)

// RegisterModels mengembalikan semua model struct dalam bentuk slice of interface
func RegisterModels() []interface{} {
	return []interface{}{
		&models.Device{},
		&models.AccessCard{},
		&models.Company{},
		&models.Department{},
		&models.Division{},
		&models.Employee{},
		&models.FrontDeskStaff{},
		&models.Section{},
		&models.Visitor{},
		&models.Visit{},
		// Tambahkan model lain di sini, cukup 1 baris
	}
}
