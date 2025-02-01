package api

import (
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/errors"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/gin-gonic/gin"
)

type PermissionApi struct {
	ac_svc services.AccountService
	r_svc  services.RoleService
	p_svc  services.PermissionService
}

func NewPermissionApi(sf services.ServiceFactory) PermissionApi {
	return PermissionApi{
		ac_svc: sf.Accounts,
		r_svc:  sf.Roles,
		p_svc:  sf.Permissions,
	}
}

func (h PermissionApi) GetAll(c *gin.Context) {
	ctx, span := utils.TracerWithGinContext(c, "api.GetAll")
	accounts, account, err := h.ac_svc.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	roles, role, err := h.r_svc.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	permissions, permission, err := h.p_svc.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	span.End()
	c.JSON(http.StatusOK, gin.H{
		"accounts":    accounts,
		"account":     account,
		"roles":       roles,
		"role":        role,
		"permissions": permissions,
		"permission":  permission,
	})
}
