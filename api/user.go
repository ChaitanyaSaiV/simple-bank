package api

import (
	"net/http"

	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/ChaitanyaSaiV/simple-bank/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type loginUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string  `json: "access_token"`
	User        userRes `json: "user"`
}

type userRes struct {
	UserName          string `json:"username"`
	FullName          string `json:"fullname"`
	Email             string `json:"email"`
	PasswordChangedAt string `json:"passwordchangedat"`
	CreatedAt         string `json:"createdat"`
}

func userResponse(User db.User) userRes {
	return userRes{
		UserName:          User.Username,
		FullName:          User.Username,
		Email:             User.Email,
		PasswordChangedAt: User.PasswordChangedAt.String(),
		CreatedAt:         User.CreatedAt.String(),
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		Username:       req.UserName,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	User, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userResponse(User))
}

type getUserReq struct {
	UserName string `form:"username" binding:"required,alphanum"`
}

func (server *Server) GetUser(ctx *gin.Context) {
	var req getUserReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	User, err := server.store.GetUser(ctx, req.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, userResponse(User))
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	access_token, err := server.tokenMaker.CreateToken(req.UserName, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := loginUserResponse{
		AccessToken: access_token,
		User:        userResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
