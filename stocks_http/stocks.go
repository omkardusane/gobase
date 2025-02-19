package stocks

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Main() {
	connStr := "user=username dbname=stocks sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/stocks", func(c *gin.Context) {
		rows, err := db.Query("SELECT symbol, price, timestamp FROM stocks ORDER BY timestamp DESC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var stocks []gin.H
		for rows.Next() {
			var symbol string
			var price float64
			var timestamp string
			if err := rows.Scan(&symbol, &price, &timestamp); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			stocks = append(stocks, gin.H{"symbol": symbol, "price": price, "timestamp": timestamp})
		}

		c.JSON(http.StatusOK, stocks)
	})

	r.Run(":8080")
}
