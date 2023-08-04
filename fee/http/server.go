package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/hosting/gin/middleware"
	log "github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/fee/controller"
	"go.opencensus.io/plugin/ochttp"
	"io"
	"net/http"
	"os"
)

type Server interface {
	Serve() error
}

type serverImpl struct {
	server   *http.Server
	endpoint controller.Controller
	config   config.AppConfig
	logger   log.Logger
}

func NewServer(endpoint controller.Controller, config config.AppConfig, logger log.Logger) Server {
	i := &serverImpl{
		endpoint: endpoint, config: config, logger: logger,
	}

	i.initServer()
	return i
}

func (i *serverImpl) Serve() error {
	i.logger.Infof(" 🚀 Listening HTTP on port %d", i.config.Http.Port)
	return i.server.ListenAndServe()
}

func (i *serverImpl) initServer() {
	if i.config.Http.EnableLogGin {
		// Send gin logs to writer
		logWriter := log.Writer()
		if logWriter != nil && *logWriter != nil {
			// write logs to both std log and custom log writer
			gin.DefaultWriter = io.MultiWriter(os.Stdout, *logWriter)
		}
	} else {
		gin.DefaultWriter = io.Discard
	}

	gin.SetMode(i.config.Http.GinMode)
	r := i.initRouter()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", i.config.Http.Port),
		Handler: &ochttp.Handler{
			Handler:          r,
			IsPublicEndpoint: true,
		},
	}
	i.server = srv
}

func (i *serverImpl) initRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.NewCORSMiddleware())

	i.registerRoutes(r)

	return r
}

func (i *serverImpl) registerRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"version": "1",
		})
	})

	r.GET("/calculation", i.endpoint.FeeCalculation)
}
