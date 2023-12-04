package api

import (
	db "github.com/RC-002/Banking_backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Serves HTTP requests for banking service

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new HTTP server and sets up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Starts server on the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
