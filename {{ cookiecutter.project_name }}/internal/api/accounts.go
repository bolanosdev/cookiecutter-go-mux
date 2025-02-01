package api

import (
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/errors"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"

	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	store db.Store
}

func NewAccountApi(store db.Store) AccountApi {
	return AccountApi{
		store: store,
	}
}

func (h AccountApi) GetAll(c *gin.Context) {
	accounts, err := h.store.GetAccount(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}
	c.JSON(http.StatusOK, accounts)
}
