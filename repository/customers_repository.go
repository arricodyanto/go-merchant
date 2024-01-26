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
	for _, v := range customers {
		if v.Username == username && v.Password == password {
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
