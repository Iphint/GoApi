package controller

import (
	"goapi/app/auth"
	"goapi/app/config"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateTransaction(c echo.Context) error {
	db := config.GetDB()

	user := c.Get("user").(*auth.Claims)
	userID := user.ID
	productIDStr := c.FormValue("product_id")
	if productIDStr == "" {
		return c.JSON(http.StatusBadRequest, "product_id is required")
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid product_id")
	}
	createdAt := time.Now()

	query := "INSERT INTO transactions (user_id, product_id, created_at) VALUES (?, ?, ?)"
	result, err := db.Exec(query, userID, productID, createdAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":        "transactions successfully created",
		"transaction_id": id,
	})
}
func ShowTransactions(c echo.Context) error {
	db := config.GetDB()

	query := "SELECT * FROM transactions"
	rows, err := db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var transactions []map[string]interface{}
	for rows.Next() {
		var id int
		var userID int
		var productID int
		var createdAt []uint8

		err := rows.Scan(&id, &userID, &productID, &createdAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
        if err != nil {
            return c.JSON(http.StatusInternalServerError, err.Error())
        }
		transaction := map[string]interface{}{
			"id":         id,
			"user_id":    userID,
			"product_id": productID,
			"created_at": createdAtTime,
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transactions)
}
func DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	db := config.GetDB()

	query := "DELETE FROM transactions WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "transaction successfully deleted",
	})
}
