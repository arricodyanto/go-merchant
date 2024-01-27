package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-merchant/config"
	"go-merchant/entity"
	"go-merchant/entity/dto"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payload entity.Payment) (entity.Payment, error)
	UpdateStatus(id string) (entity.Payment, error)
}

type paymentRepository struct {
}

// Create implements PaymentRepository.
func (*paymentRepository) Create(payload entity.Payment) (entity.Payment, error) {
	// baca file dari file customers.json untuk dicocokan dengan payload.customer_id
	var customers []entity.Customer

	path := "repository/json/"
	customerFileName := "customers.json"
	file, err := os.ReadFile(filepath.Join(path, customerFileName))
	if err != nil {
		return entity.Payment{}, fmt.Errorf("failed to read file: %v", err.Error())
	}
	json.Unmarshal(file, &customers)

	// cari customer dengan payload.customer_id yang cocok
	var matchedCustomer entity.Customer
	for _, v := range customers {
		// ubah valuenya ke null agar tidak masuk ke histories.json
		v.IsLoggedIn = false
		v.Password = ""

		// cek id
		if v.ID == payload.CustomerID {
			matchedCustomer = v
			break
		}
	}

	if matchedCustomer.ID == "" {
		return entity.Payment{}, fmt.Errorf("customer not found with ID: %s", payload.CustomerID)
	}

	var payment entity.Payment

	// generate uuid for id payment
	paymentID := uuid.New().String()

	// make request body
	requestObj := dto.SnapRequestDto{
		TransactionDetails: dto.TransactionDetailDto{
			OrderID:     paymentID,
			GrossAmount: payload.Amount,
		},
		CreditCard: dto.CreditCardDto{
			Secure: true,
		},
	}
	requestBody, _ := json.Marshal(requestObj)

	// make request to midtrans
	req, err := http.NewRequest(http.MethodPost, config.PaymentSnap, bytes.NewBuffer(requestBody))
	if err != nil {
		return entity.Payment{}, fmt.Errorf("failed to create request: %v", err.Error())
	}

	// Menambahkan header Content-Type
	req.Header.Set("Content-Type", "application/json")

	// pasang authorization heeader ke request
	username := os.Getenv("MIDTRANS_SERVER_KEY")
	password := ""
	req.SetBasicAuth(username, password)

	// bertindak sebagai client
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return entity.Payment{}, fmt.Errorf("failed do request: %v", err.Error())

	}

	if response.StatusCode != http.StatusCreated {
		return entity.Payment{}, fmt.Errorf("failed to retrieve response, status code: %d", response.StatusCode)
	}

	// make response body
	responseBody, _ := io.ReadAll(response.Body)

	// convert response body ke struct
	var responseData dto.SnapResponseDto
	_ = json.Unmarshal(responseBody, &responseData)

	// masukkan value payload ke payment
	payment.ID = paymentID
	payment.CustomerID = payload.CustomerID
	payment.Amount = payload.Amount
	payment.Token = responseData.Token
	payment.RedirectURL = responseData.RedirectURL
	payment.TransactionTime = time.Now()

	// read isi data dari file histories.json
	var histories []entity.History

	fileName := "histories.json"
	file, err = os.ReadFile(filepath.Join(path, fileName))
	if err != nil {
		fmt.Printf("failed to read file: %v", err.Error())
		// buat file jika belum ada
		file, err := os.Create(filepath.Join(path, fileName))
		if err != nil {
			return entity.Payment{}, fmt.Errorf("failed to create customers file: %v", err.Error())
		}
		defer file.Close()
	}
	json.Unmarshal(file, &histories)

	// buat object baru untuk disimpan ke dalam file histories.json
	historiesJson := entity.History{
		ID:                paymentID,
		Customer:          matchedCustomer,
		Amount:            payment.Amount,
		RedirectURL:       payment.RedirectURL,
		Currency:          "",
		PaymentType:       "",
		TransactionStatus: "Waiting Payment",
		TransactionTime:   payment.TransactionTime,
		SettlementTime:    time.Time{},
		ExpiryTime:        time.Time{},
	}

	// append createdData ke histories varible
	histories = append(histories, historiesJson)

	// ubah ke format json dengan marshal untuk disimpan ke history.json
	createdData, err := json.MarshalIndent(histories, "", "  ")
	if err != nil {
		return entity.Payment{}, fmt.Errorf("failed to marshal createdData: %v", err.Error())
	}

	// save data ke file history.json
	err = os.WriteFile(filepath.Join(path, fileName), createdData, 0644)
	if err != nil {
		return entity.Payment{}, fmt.Errorf("failed to save createdData to json: %v", err.Error())
	}

	return payment, nil
}

// UpdateStatus implements PaymentRepository.
func (*paymentRepository) UpdateStatus(id string) (entity.Payment, error) {
	panic("unimplemented")
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{}
}
