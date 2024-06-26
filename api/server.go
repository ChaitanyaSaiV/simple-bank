package api

import (
	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(), // Initialize the router here
	}
	server.router.GET("/ping", server.Ping)
	server.router.POST("/accounts", server.CreateAccount)
	server.router.GET("/account/:id", server.GetAccount)
	server.router.GET("/listaccounts", server.ListAccounts)
	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
