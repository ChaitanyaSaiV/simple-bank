package api

import (
	"fmt"

	token "github.com/ChaitanyaSaiV/simple-bank/token"
	"github.com/ChaitanyaSaiV/simple-bank/util"

	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/gin-gonic/gin"
)

type Server struct {
	tokenMaker token.Maker
	store      *db.Store
	router     *gin.Engine
	config     util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		router:     gin.Default(), // Initialize the router here
		tokenMaker: tokenMaker,
	}

	server.router.GET("/ping", server.Ping)
	server.router.POST("/users", server.CreateUser)
	server.router.POST("/user/login", server.loginUser)

	authRoutes := server.router.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.GET("/account/:id", server.GetAccount)
	authRoutes.GET("/listaccounts", server.ListAccounts)
	authRoutes.POST("/transfer", server.CreateTransfer)
	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/users", server.GetUser)
	return server, nil
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
