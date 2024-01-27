package repository

import (
	"encoding/json"
	"fmt"
	"go-merchant/entity"
	"os"
	"path/filepath"
)

type CustomerRepository interface {
	GetByUsernamePassword(username, password string) (entity.Customer, error)
}

type customerRepository struct {
}

// GetByUsernamePassword implements CustomerRepository.
func (*customerRepository) GetByUsernamePassword(username, password string) (entity.Customer, error) {
	// buat variable berdasarkan model yang didefined
	var customers []entity.Customer
	var customer entity.Customer

	// baca data customer dari file json
	path := "repository/json/"
	fileName := "customers.json"
	file, err := os.ReadFile(filepath.Join(path, fileName))
	if err != nil {
		fmt.Printf("failed to read file: %v", err.Error())
		// buat file jika belum ada
		file, err := os.Create(filepath.Join(path, fileName))
		if err != nil {
			return entity.Customer{}, fmt.Errorf("failed to create customers file: %v", err.Error())
		}
		defer file.Close()
	}
	json.Unmarshal(file, &customers)

	// cek data yang memiliki credential username password yang cocok
	for i, v := range customers {
		if v.Username == username && v.Password == password {
			// ubah field is_logged_in menjadi true
			customers[i].IsLoggedIn = true
			customer = customers[i]

			// ubah ke format json dengan marshal
			updatedData, err := json.MarshalIndent(customers, "", "  ")
			if err != nil {
				return entity.Customer{}, fmt.Errorf("failed to marshal updatedData: %v", err.Error())
			}

			// sumpan perubahan ke file json
			err = os.WriteFile(filepath.Join(path, fileName), updatedData, 0644)
			if err != nil {
				return entity.Customer{}, fmt.Errorf("failed to save updatedData to json: %v", err.Error())
			}

			return customer, nil
		}
	}
	// apabila tidak ada yang cocok return error
	return entity.Customer{}, fmt.Errorf("credentials did not match")
}

// buat bridge
func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}
