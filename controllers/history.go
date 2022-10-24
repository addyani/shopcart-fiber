package controllers

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/task1/database"
)

type HistoryController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitHistoryController(s *session.Store) *HistoryController {
	db := database.InitDb()
	return &HistoryController{Db: db, store: s}
}
