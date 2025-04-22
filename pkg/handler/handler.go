package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"smartbridge-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/vshep/v1", h.callVshep)
	return router
}

func (h *Handler) callVshep(c *gin.Context) {
	logrus.Info("Start CallVshep")

	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Request is not XML!"))
		return
	}

	statusCode, body, err := h.services.SendRequest(request)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Data(statusCode, "text/xml", body)
	logrus.Info("Stop CallVshep")
	return
}
