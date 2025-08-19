package registrationmodel

import (
	"guestbook/models"
)

// RegisterModels mengembalikan semua model struct dalam bentuk slice of interface
func RegisterModels() []interface{} {
	return []interface{}{
		&models.Device{},

		// Tambahkan model lain di sini, cukup 1 baris
	}
}
