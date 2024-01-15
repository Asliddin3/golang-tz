package controller

import (
	"github.com/Asliddin3/tz/config"
	"github.com/Asliddin3/tz/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db  *gorm.DB
	log *logger.MyLogger
	cfg *config.Config
}

func NewHandler(db *gorm.DB, log *logger.MyLogger, cfg config.Config) *Handler {
	return &Handler{
		db:  db,
		log: log,
		cfg: &cfg,
	}
}

func (h *Handler) Init(server *gin.Engine) {
	h.log.Debug("Initializing handlers")
	api := server.Group("api")
	{
		h.NewPersonController(api)
	}
}

type Response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, Response{message})
}
