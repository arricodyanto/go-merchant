package repository

import (
	"encoding/json"
	"fmt"
	"go-merchant/entity"
	"go-merchant/shared/common"
	"os"
	"path/filepath"
)

type CustomerRepository interface {
	GetByUsernamePassword(username, password string, isLoggedIn bool) (entity.Customer, error)
}

type customerRepository struct {
}

// GetByUsernamePassword implements CustomerRepository.
func (*customerRepository) GetByUsernamePassword(username, password string, isLoggedIn bool) (entity.Customer, error) {
	// buat variable berdasarkan model yang didefined
	var customers []entity.Customer
	var customer entity.Customer

	// baca data customer dari file json
	path := "repository/json/"
	fileName := "customers.json"
	err := common.ReadJsonFile(fileName, &customers)
	if err != nil {
		fmt.Printf("failed to read file: %v", err.Error())
		// buat file jika belum ada
		err = common.CreateJsonFile(fileName)
		if err != nil {
			return entity.Customer{}, err
		}
	}

	// cek data yang memiliki credential username password yang cocok
	for i, v := range customers {
		comparePassword := common.ComparePassword(password, v.Password)
		if v.Username == username && comparePassword {
			// ubah field is_logged_in menjadi true
			customers[i].IsLoggedIn = isLoggedIn
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
