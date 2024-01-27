package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type TokenConfig struct {
	IssuerName       string `json:"IssuerName"`
	JwtSignatureKy   []byte `json:"JwtSignatureKy"`
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type Config struct {
	ApiConfig
	TokenConfig
}

func (c *Config) ConfigConfiguration() error {
	// Cek apakah file env ada
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}

	// simpan api port env ke variable
	c.ApiConfig = ApiConfig{ApiPort: os.Getenv("API_PORT")}

	// buat config untuk jwt token yang akan digenerate
	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	c.TokenConfig = TokenConfig{
		IssuerName:       os.Getenv("TOKEN_ISSUE"),
		JwtSignatureKy:   []byte(os.Getenv("TOKEN_SECRET")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiresTime:   time.Duration(tokenExpire) * time.Minute,
	}

	if c.ApiPort == "" || c.IssuerName == "" || c.JwtExpiresTime < 0 || len(c.JwtSignatureKy) == 0 {
		return fmt.Errorf("missing required environment")
	}

	return nil
}

// buat bridge
func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.ConfigConfiguration(); err != nil {
		return nil, err
	}
	return config, nil
}
