package controller

import (
	"database/sql"
	"goapi/app/auth"
	"goapi/app/config"
	"goapi/app/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateProductHandler(c echo.Context) error {
	db := config.GetDB()
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Mengambil nilai dari form
	name := c.FormValue("name")
	category := c.FormValue("category")
	price := c.FormValue("price")
	condition := c.FormValue("condition")

	// Mendapatkan user ID dari context (pastikan sudah ada middleware untuk menyetel user di context)
	user := c.Get("user").(*auth.Claims)
	product.UserID = user.ID

	// Menyimpan produk ke database
	query := "INSERT INTO products (user_id, name, category, price, `condition`) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, product.UserID, name, category, price, condition)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	product.ID = int(id)

	// Memastikan folder uploads ada
	uploadPath := "app/uploads"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err = os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	// Mengambil file gambar dari form (multiple)
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	files := form.File["images"]

	for _, file := range files {
		// Membuka file
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		// Menentukan lokasi penyimpanan file
		filename := strings.ReplaceAll(file.Filename, " ", "_")
		filepath := filepath.Join(uploadPath, filename)
		dst, err := os.Create(filepath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer dst.Close()

		// Menyalin konten file ke lokasi penyimpanan
		if _, err := io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Menyimpan informasi gambar ke tabel images
		imageQuery := "INSERT INTO images (product_id, path) VALUES (?, ?)"
		_, err = db.Exec(imageQuery, product.ID, filepath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "data successfully created",
		"id":      product.ID,
	})
}

func ShowProductHandler(c echo.Context) error {
	db := config.GetDB()
	id := c.Param("id")
	product := new(models.Product)
	query := "SELECT id, user_id, name, category, price, `condition` FROM products WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&product.ID, &product.UserID, &product.Name, &product.Category, &product.Price, &product.Condition)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "product not found")
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Mengambil gambar-gambar terkait dari tabel images
	imageQuery := "SELECT path FROM images WHERE product_id = ?"
	rows, err := db.Query(imageQuery, product.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	images := []string{}
	for rows.Next() {
		var imagePath string
		if err := rows.Scan(&imagePath); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		images = append(images, imagePath)
	}
	product.Images = images

	return c.JSON(http.StatusOK, product)
}

func ShowProductsHandler(c echo.Context) error {
	db := config.GetDB()
	products := []models.Product{}

	query := "SELECT id, user_id, name, category, price, `condition` FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.UserID, &product.Name, &product.Category, &product.Price, &product.Condition); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Mengambil gambar-gambar terkait dari tabel images
		imageQuery := "SELECT path FROM images WHERE product_id = ?"
		imageRows, err := db.Query(imageQuery, product.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer imageRows.Close()

		images := []string{}
		for imageRows.Next() {
			var imagePath string
			if err := imageRows.Scan(&imagePath); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			images = append(images, imagePath)
		}
		if err = imageRows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		product.Images = images
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func UpdateProducthandler(c echo.Context) error {
	db := config.GetDB()
	id := c.Param("id")
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := "UPDATE products SET user_id = ?, name = ?, category = ?, price = ?, condition = ?, image = ? WHERE id = ?"
	_, err := db.Exec(query, product.UserID, product.Name, product.Category, product.Price, product.Condition, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func DeleteProducthandler(c echo.Context) error {
	db := config.GetDB()
	id := c.Param("id")

	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "product deleted")
}
