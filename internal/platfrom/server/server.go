package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
	"github.com/osmait/blog-go/internal/find"
	"github.com/osmait/blog-go/internal/platfrom/server/handler/auth"
	"github.com/osmait/blog-go/internal/platfrom/server/handler/health"
	"github.com/osmait/blog-go/internal/platfrom/server/handler/post"
	"github.com/osmait/blog-go/internal/platfrom/server/handler/user"
)

type Server struct {
	httpAddr          string
	engine            *gin.Engine
	CreateUseService  creating.CreateUseService
	CreatePostService creating.CreatePostService
	findUserService   find.FindUserService
}

func NewServer(host string, port uint, creatinUserService creating.CreateUseService, creatingPostService creating.CreatePostService, findUserService find.FindUserService) Server {
	srv := Server{
		engine:            gin.New(),
		httpAddr:          fmt.Sprintf("%s:%d", host, port),
		CreateUseService:  creatinUserService,
		CreatePostService: creatingPostService,
		findUserService:   findUserService,
	}
	srv.routes()
	return srv

}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) routes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/users", user.CreateUser(s.CreateUseService))
	s.engine.POST("/posts", post.Createpost(s.CreatePostService))
	s.engine.POST("/login", auth.Login(s.findUserService))

}
