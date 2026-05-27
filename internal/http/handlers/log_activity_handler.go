package handlers

import (
	"net/http"

	"log-service/internal/logactivity"
	"log-service/internal/response"

	"github.com/gin-gonic/gin"
)

type LogActivityHandler struct {
	service *logactivity.Service
}

func NewLogActivityHandler(service *logactivity.Service) *LogActivityHandler {
	return &LogActivityHandler{
		service: service,
	}
}

func (h *LogActivityHandler) Create(c *gin.Context) {
	var req logactivity.CreateLogActivityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	log, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, log)
}

func (h *LogActivityHandler) FindAll(c *gin.Context) {
	logs, err := h.service.FindAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, logs)
}
