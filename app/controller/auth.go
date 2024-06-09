package controller

import (
	"goapi/app/auth"
	"goapi/app/config"
	"goapi/app/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Struct to parse the register request body
type RegisterRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

// Struct to parse the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	// Log values for debugging
	log.Printf("Register - Username: %s, Email: %s, Password: %s, RepeatPassword: %s", req.Username, req.Email, req.Password, req.RepeatPassword)

	if req.Password != req.RepeatPassword {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Passwords do not match"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error encrypting password"})
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	db := config.GetDB()
	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		// Log the error for debugging
		log.Printf("Error inserting user into database: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error inserting user into database"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User registered successfully"})
}

func Login(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	// Log values for debugging
	log.Printf("Login - Username: %s, Password: %s", req.Username, req.Password)

	db := config.GetDB()
	var user models.User
	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = ? OR email = ?", req.Username, req.Username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	token, err := auth.GenerateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error generating token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func GetUserById(c echo.Context) error {
	db := config.GetDB()
    id := c.Param("id")

    var user models.User
    err := db.QueryRow("SELECT id, username, email FROM users WHERE id =?", id).Scan(&user.ID, &user.Username, &user.Email)
    if err!= nil {
        return c.JSON(http.StatusNotFound, "user not found")
    }

    return c.JSON(http.StatusOK, user)
}
