package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/models"
	"github.com/web-gopro/auth_exam/pkg/helpers"
	"github.com/web-gopro/auth_exam/token"
)

// @Summary Create a new sysuser
// @Router /api/super/sysuser_create [post]
// @Description Create a new system user (admin, buxgalter, etc.) with role assignment
// @Tags sysusers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysuser body models.SysUserCretReq true "SysUser object to create"
// @Success 201 {object} models.SysUserCreateResp "Successfully created"
// @Failure 400 {string} string "Invalid input"
// @Failure 409 {string} string "SysUser already exists"
// @Failure 500 {string} string "Internal server error"
func (h *Handler) SysUserCreate(ctx *gin.Context) {

	var reqBody models.SysUserCretReq

	err := ctx.BindJSON(&reqBody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	tokenString := ctx.GetHeader("Authorization")

	tokenClaim, err := token.ParseJWT(tokenString)

	createdBy:= &tokenClaim.UserId

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	reqBody.Password, err = helpers.HashPassword(reqBody.Password)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	resp, err := h.storage.SysUserRepo().CreateSysUser(context.Background(), reqBody,*createdBy)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)

}



// @Summary Login a sysuser
// @Router /api/admp/login [post]
// @Description login a system user
// @Tags admp
// @Accept json
// @Produce json
// @Param sysuser body models.LoginReq true "SysUser object to Login"
// @Success 201 {string} string "Successfully LogedIn"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
func (h *Handler) SysUserLogin(ctx *gin.Context) {

	var reqBody models.LoginReq

	ctx.BindJSON(&reqBody)

	claim, err := h.storage.SysUserRepo().SysUserLogin(ctx, reqBody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	accessToken, err := token.GenerateJWT(*claim)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(201, accessToken)

}


// GetSysUser godoc
// @Summary Get all system users
// @Description Only accessible by superadmin
// @Tags sysusers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} models.SysUser
// @Router /api/super/sysuser [get]
func (h *Handler) GetSysUser(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")

	claim, err := token.ParseJWT(tokenString)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	resp, err := h.storage.SysUserRepo().GetSysUser(context.Background(), models.GetById{Id: claim.UserId})
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(201, resp)

}
