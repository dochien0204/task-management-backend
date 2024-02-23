package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TxMiddleware interface {
	DBTransactionMiddleware() gin.HandlerFunc
}

type MiddlewareRepository struct {
	db *gorm.DB
}

func NewMiddlewareRepository(db *gorm.DB) *MiddlewareRepository {
	return &MiddlewareRepository{
		db: db,
	}
}

func StatusInList(status int, statusList []int) bool {
	for _, value := range statusList {
		if value == status {
			return true
		}
	}

	return false
}

func (r *MiddlewareRepository) DBTransactionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := r.db.Begin()
		fmt.Println()
		log.Print("\033[32m", "[transaction] Beginning transactions", "\033[0m")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()
		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			fmt.Println()
			log.Print("\033[32m", "[transaction] committing transactions", "\033[0m")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error", err)
			}
		} else {
			fmt.Println()
			log.Println("\033[31m", "[transaction] rolling back transaction due to status code:", c.Writer.Status(), "\033[0m")
			txHandle.Rollback()
		}
	}
}
