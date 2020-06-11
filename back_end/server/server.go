package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/vpoliakov01/CoverageMonitor/back_end/server/git"
)

// Config is the server config
type Config struct {
	Port int
	Mock bool
}

// Server is used to serve requests
type Server struct {
	Config
	Router *gin.Engine
}

// New creates a Server instance
func New(c *Config) *Server {
	return &Server{
		Config: *c,
		Router: gin.New(),
	}
}

// Serve serves requests
func (s *Server) Serve() error {
	publicAPI := s.Router.Group("/api")
	{
		if s.Mock {
			publicAPI.GET("/:github_org/:github_repo/info", git.GetRepoMockHandler)
		} else {
			publicAPI.GET("/:github_org/:github_repo/info", git.GetRepoHandler)
		}
		publicAPI.POST("/:github_org/:github_repo/test", git.TestRepoHandler)
	}
	return s.Router.Run(fmt.Sprintf(":%v", s.Port))
}
