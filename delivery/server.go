package delivery

import (
	"fmt"
	"go-merchant/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
	// ambilapi group dari file app config
	rg := s.engine.Group(config.ApiGroup)
	fmt.Print(rg)

	// routes from controller

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

	// config gin to default config
	engine := gin.Default()
	host := fmt.Sprintf(":%s", config.ApiPort)

	return &Server{
		engine: engine,
		host:   host,
	}
}
