package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/auth_exam/mail"
	"github.com/web-gopro/auth_exam/models"
)

// import (
// 	"context"
// 	"encoding/json"

// 	"github.com/gin-gonic/gin"
// 	"github.com/web-gopro/auth_exam/token"

// 	"github.com/saidamir98/udevs_pkg/logger"
// )

func (h *Handler) UserCreate(ctx *gin.Context) {

	var req models.User

	ctx.BindJSON(&req)

	resp, err := h.storage.UserRepo().CreateUser(context.Background(), req)

	if err != nil {

		fmt.Println(err.Error())
		ctx.JSON(500, err.Error())
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetUserById(ctx *gin.Context) {

	var req models.GetById

	req.Id = ctx.Param("id")

	resp, err := h.storage.UserRepo().GetUserById(context.Background(), req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) CheckUser(ctx *gin.Context) {

	var reqBody models.Check_User

	err := ctx.BindJSON(&reqBody)

	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	isExists, err := h.storage.UserRepo().IsExists(context.Background(), models.Common{
		Table_name:  "users",
		Column_name: "email",
		Expvalue:    reqBody.Email,
	})

	if err != nil {

		ctx.JSON(500, err)
		return
	}

	if isExists.IsExists {
		ctx.JSON(201, models.CheckExists{
			Is_exists: isExists.IsExists,
			Status:    "sign-in",
		})
		return
	}
	otp := models.OtpData{
		Otp:   mail.GenerateOtp(6),
		Email: reqBody.Email,
	}

	otpdataB, err := json.Marshal(otp)

	if err != nil {

		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaa")
		ctx.JSON(500, err.Error())
		return
	}

	err = h.cache.Set(ctx, reqBody.Email, string(otpdataB), 120)

	err = mail.SendMail([]string{reqBody.Email}, otp.Otp)

	if err != nil {
		h.log.Error("errrr on Send mail", logger.Error(err))
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, models.CheckExists{
		Is_exists: isExists.IsExists,
		Status:    "registr",
	})

	ctx.JSON(201, "we sent otp")

}

// func (h *Handler) SignUp(ctx *gin.Context) {

// 	var otpData book_shop.OtpData

// 	var reqBody book_shop.UserCreateReq

// 	err := ctx.BindJSON(&reqBody)

// 	if err != nil {
// 		h.log.Error("errrr on ShouldBindJSON", logger.Error(err))
// 		ctx.JSON(500, err.Error())
// 		return
// 	}

// 	otpSData, err := h.cache.GetDell(ctx, reqBody.Email)

// 	if err != nil {
// 		ctx.JSON(500, err.Error())
// 		return
// 	}

// 	if otpSData == "" {
// 		ctx.JSON(201, "otp is expired")
// 		return
// 	}
// 	err = json.Unmarshal([]byte(otpSData), &otpData)

// 	if otpData.Otp != reqBody.Otp {

// 		ctx.JSON(405, "incorrect otp")
// 		return
// 	}

// 	reqBody.Password, err = helpers.HashPassword(reqBody.Password)

// 	claim, err := h.service.GetUserSevice().CreateUser(context.Background(), &reqBody)

// 	if err != nil {
// 		ctx.JSON(500, err.Error())
// 		return
// 	}

// 	accessToken, err := token.GenerateJWT(*&book_shop.Clamis{UserId: claim.UserId, UserRole: claim.UserRole})

// 	if err != nil {
// 		ctx.JSON(201, "registreted")
// 		return
// 	}

// 	ctx.JSON(201, accessToken)

// }

// func (h *Handler) SigIn(ctx *gin.Context) {

// 	var reqBody book_shop.UserLogIn

// 	err := ctx.BindJSON(&reqBody)

// 	if err != nil {
// 		ctx.JSON(500, err.Error())
// 		return
// 	}

// 	claim, err := h.service.GetUserSevice().UserLogin(ctx, &reqBody)

// 	if err != nil {
// 		if err.Error() == "password is incorrect" {
// 			ctx.JSON(405, err.Error())
// 			return
// 		}
// 		ctx.JSON(500, err.Error())
// 		return
// 	}

// 	accessToken, err := token.GenerateJWT(*claim)

// 	if err != nil {
// 		ctx.JSON(500, err.Error())
// 		return
// 	}

// 	ctx.JSON(201, accessToken)

// }
