# Go Merchant API

Restful API to provide services between merchants and banks. Registered customers can transfer money
without a minimum limit to merchants via a payment gateway, which in this project uses Midtrans as
the payment gateway. This API does not use a database, all customers, payments & history data is
stored in a json file in the repository folder. As initial data, it is provided in the /assets
folder. The file in question is as follows.

```
/repository/json/customers.json
/repository/json/histories.json
```

## Postman Documentation

```
https://documenter.getpostman.com/view/17920856/2s9YyqiMvw
```

## Install & Dependence

- go

## Dataset Preparation

## Use

- clonning repo

```
git clone https://github.com/arricodyanto/go-merchant.git
```

- running project

```
go run .
```

## Pretrained model

## Available Endpoints

### Authorization Login

1. Login

- Method = POST
- Body

```
{
    "username": "yoelsefet12",
    "password": "password"
}
```

- Example Response

```
{
  "status": {
    "code": 201,
    "message": "Ok"
  },
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnb19tZXJjaGFudCIsImV4cCI6MTcwNjM2OTI0MywiaWF0IjoxNzA2MzY1NjQzLCJjdXN0b21lcl9pZCI6IjFjMWE2YmExLTEyMjUtNDM4NS1hMjdhLWE5YWI1NzdkNjFhMSJ9.SAkCGZ5xCCI53gtBV8bXP_3fOeJnZwph2p2Utx66yzY"
  }
}
```

2. Create Payment

- Header = AUTHORIZATION Bearer Token
- Method = POST
- Body

```
{
    "customer_id": "1c1a6ba1-1225-4385-a27a-a9ab577d61a1d",
    "amount": 1
}
```

- Example Response

```
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": "8d1f6fe7-39b5-4616-a464-da461081472d",
    "customer_id": "1c1a6ba1-1225-4385-a27a-a9ab577d61a1",
    "amount": 1,
    "token": "a674af11-27ed-4b24-8f84-dc286b08b8c2",
    "redirect_url": "https://app.midtrans.com/snap/v3/redirection/a674af11-27ed-4b24-8f84-dc286b08b8c2",
    "transaction_time": "2024-01-27T21:30:56.095469+07:00"
  }
}
```

3. Logout

- Header = AUTHORIZATION Bearer Token
- Method = POST
- Body

```
{
    "username": "yoelsefet12",
    "password": "password"
}
```

- Example Response

```
{
    "code": 200,
    "message": "Ok"
}
```

## Directory Hierarchy

```
|—— assets
|    |—— customers.json
|    |—— histories.json
|—— config
|    |—— app_config.go
|    |—— config.go
|—— delivery
|    |—— controller
|        |—— auth_controller.go
|        |—— payments_controller.go
|    |—— middleware
|        |—— auth_middleware.go
|    |—— server.go
|—— entity
|    |—— customers.go
|    |—— dto
|        |—— auth_dto.go
|        |—— payments_dto.go
|    |—— histories.go
|    |—— payments.go
|—— mock
|—— public
|—— repository
|    |—— customers_repository.go
|    |—— json
|        |—— customers.json
|        |—— histories.json
|    |—— payments_repository.go
|—— shared
|    |—— common
|        |—— json_data.go
|        |—— json_response.go
|    |—— model
|        |—— json_model.go
|        |—— mycustom_model.go
|    |—— service
|        |—— jwt_service.go
|—— usecase
|    |—— auth_usecase.go
|    |—— customers_usecase.go
|    |—— payments_usecase.go
|—— go.mod
|—— go.sum
|—— main.go
|—— .env
|—— .env.example
|—— .gitignore
```

## Credits

[Arrico Handyanto (Vercel)](https://arricohandyanto.vercel.app) -
[Arrico Handyanto (LinkedIn)](https://www.linkedin.com/in/arricohandyanto/)
