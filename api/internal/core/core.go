package core

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Core contains the main api logic
type Core struct {
	Router   *mux.Router
	Database *sql.DB
}

// NewCore is our mcore factory
func NewCore(router *mux.Router, database *sql.DB) *Core {
	return &Core{
		Router:   router,
		Database: database,
	}
}
