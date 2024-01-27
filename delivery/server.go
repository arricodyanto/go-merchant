package delivery

import (
	"fmt"
	"go-merchant/config"
	"go-merchant/delivery/controller"
	"go-merchant/repository"
	"go-merchant/shared/service"
	"go-merchant/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	customerUC usecase.CustomerUsecase
	paymentUC  usecase.PaymentUsecase
	authUC     usecase.AuthUsecase
	jwtService service.JwtService
	engine     *gin.Engine
	host       string
}

func (s *Server) initRoute() {
	// ambilapi group dari file app config
	rg := s.engine.Group(config.ApiGroup)

	// routes from controller
	controller.NewAuthController(s.authUC, rg).Route()
	controller.NewPaymentController(s.paymentUC, rg).Route()
}

func (s *Server) Run() {
	// jalankan routes yang ada dari initRoute()
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server could not run at port (%s) because of error: %v", s.host, err.Error()))
	}
}

// buat bridge
func NewServer() *Server {
	// panggil config yang telah dibuat
	config, _ := config.NewConfig()

	// define repository
	customerRepo := repository.NewCustomerRepository()
	paymentRepo := repository.NewPaymentRepository()

	// define usecase
	customerUC := usecase.NewCustomerUsecase(customerRepo)
	paymentUC := usecase.NewPaymentUsecase(paymentRepo)
	jwtService := service.NewJwtService(config.TokenConfig)
	authUC := usecase.NewAuthUsecase(customerUC, jwtService)

	// config gin to default config
	engine := gin.Default()
	host := fmt.Sprintf(":%s", config.ApiPort)

	return &Server{
		customerUC: customerUC,
		paymentUC:  paymentUC,
		jwtService: jwtService,
		authUC:     authUC,
		engine:     engine,
		host:       host,
	}
}
