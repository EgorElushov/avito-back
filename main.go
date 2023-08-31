package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {

	router := gin.Default()
	router.GET("/user/:id", getUserByID)
	router.POST("")

	router.Run("localhost:8080")
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	db, err := sql.Open(driver, dataSourceName)

	result, err := db.ExecContext(ctx,
		"SELECT segment FROM user_segments WHERE user_id=$1",
		id,
	)
	if err != nil {
		c.IndentedJSON(http.StatusOK, result)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
