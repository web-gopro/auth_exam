package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/models"
)

func(h *Handler)SysUserCreate(ctx *gin.Context){

	reqBody:=models.SysUserCretReq{}

	ctx.BindJSON(&reqBody)

	resp,err:=h.storage.SysUserRepo().CreateSysUser(context.Background(),reqBody)

		if err != nil {

		ctx.JSON(500, err.Error())
	}

	ctx.JSON(201,resp)
}