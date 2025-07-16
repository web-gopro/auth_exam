package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/models"
	"github.com/web-gopro/auth_exam/pkg/helpers"
	"github.com/web-gopro/auth_exam/token"
)

func (h *Handler) SysUserSinUp(ctx *gin.Context) {

	var reqBody models.SysUserCretReq

	err := ctx.BindJSON(&reqBody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	tokenString := ctx.GetHeader("Authorization")

	tokenClaim, err := token.ParseJWT(tokenString)

	reqBody.CreatedBy = &tokenClaim.UserId

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	reqBody.Password, err = helpers.HashPassword(reqBody.Password)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	resp, err := h.storage.SysUserRepo().CreateSysUser(context.Background(), reqBody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)

}

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
