package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/models"
	"github.com/web-gopro/auth_exam/token"
)

// @Summary Create a new Role
// @Router /api/super/role [post]
// @Description Create a new role user (admin, buxgalter, etc.) with role assignment
// @Tags role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysuser body models.CreateRoleRequest true "Role object to create"
// @Success 201 {object} models.Role "Successfully created"
// @Failure 400 {string} string "Invalid input"
// @Failure 409 {string} string "Role already exists"
// @Failure 500 {string} string "Internal server error"
func (h *Handler) RoleCreate(ctx *gin.Context) {

	reqbody := models.CreateRoleRequest{}
	err := ctx.BindJSON(&reqbody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	tokenString := ctx.GetHeader("Authorization")

	tokenClaim, err := token.ParseJWT(tokenString)

	createdBy := &tokenClaim.UserId

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	resp, err := h.storage.RoleRepo().Create(context.Background(), &reqbody, *createdBy)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(201, resp)

}


// @Summary Update Role
// @Router /api/super/role [put]
// @Description Update system user role (admin, buxgalter, etc.) with role assignment
// @Tags role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysuser body models.UpdateRoleRequest true "SysUser object to create"
// @Success 201 {object} models.Role "Successfully Updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 409 {string} string "SysUser already exists"
// @Failure 500 {string} string "Internal server error"
func (h *Handler) RoleUpdate(ctx *gin.Context) {

	reqbody := models.UpdateRoleRequest{}
	ctx.BindJSON(&reqbody)

	resp,err:=h.storage.RoleRepo().Update(context.Background(),&reqbody)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200,resp)

}


// @Summary Get Role by ID
// @Description Get system role details by ID
// @Tags role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} models.Role "Successfully Retrieved"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Role not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/super/role/{id} [get]
func (h *Handler) RoleGetById(ctx *gin.Context) {

	var req models.GetById

	req.Id = ctx.Param("id")

	resp, err := h.storage.RoleRepo().GetByID(context.Background(), req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

