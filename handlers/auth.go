package handlers

import (
	authdtos "absence-click/dtos/auth"
	dtos "absence-click/dtos/result"
	"absence-click/models"
	"absence-click/packages/bcrypt"
	"absence-click/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerAuth struct {
	AuthRepositories repositories.AuthRepositories
}

func HandlerAuth(AuthRepositories repositories.AuthRepositories) *handlerAuth {
	return &handlerAuth{AuthRepositories}
}

func (h *handlerAuth) Register(c *gin.Context) {
	request := new(authdtos.AuthRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dtos.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dtos.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dtos.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: password,
	}

	data, err := h.AuthRepositories.Register(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	c.JSON(
		http.StatusOK,
		dtos.SuccessResult{
			Code: http.StatusOK,
			Data: data,
		},
	)
}
