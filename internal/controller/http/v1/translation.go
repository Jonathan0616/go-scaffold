package v1

import (
	"github.com/gin-gonic/gin"
	"go-scaffold/internal/entity"
	"net/http"

	"go-scaffold/internal/usecase"
	"go-scaffold/pkg/logger"
)

type translationRoutes struct {
	t usecase.Translation
	l logger.Interface
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *translationRoutes) history(c *gin.Context) {
	translations, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}
	c.JSON(http.StatusOK, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	Original    string `json:"original" binding:"required"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
func (r *translationRoutes) doTranslate(c *gin.Context) {
	var request doTranslateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	translation, err := r.t.Translate(
		c.Request.Context(),
		entity.Translation{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "translation service problems")
		return
	}

	c.JSON(http.StatusOK, translation)
}

func newTranslationRoutes(handler *gin.RouterGroup, t usecase.Translation, l logger.Interface) {
	r := &translationRoutes{t, l}

	h := handler.Group("translation")
	{
		h.GET("/history", r.history)
		h.POST("/do-translate", r.doTranslate)

	}
}
