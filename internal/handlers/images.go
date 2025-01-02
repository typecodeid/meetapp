package handlers

import (
	"io"
	utils "meetapp/pkg/database"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// PostImage godoc
// @Summary Upload an image
// @Description Upload an image
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"
// @Success 200 {object} map[string]interface{}
// @Router /images [post]
func PostImage(c echo.Context) error {
	// Ambil file dari form data
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to get the image file"})
	}

	// Buka file untuk membaca
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to open the image file"})
	}
	defer src.Close()

	// Generate ID dan nama file unik
	imageID := uuid.New().String()
	uniqueFileName := imageID + filepath.Ext(file.Filename)

	// Tentukan lokasi penyimpanan file
	savePath := filepath.Join("uploads", uniqueFileName)
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	// Simpan file di server
	dst, err := os.Create(savePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save the image file"})
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to copy the image file"})
	}

	// Metadata gambar
	image := struct {
		ID       string
		Name     string
		MimeType string
		URL      string
		Size     int64
	}{
		ID:       imageID,
		Name:     file.Filename,
		MimeType: file.Header.Get("Content-Type"),
		URL:      "/uploads/" + uniqueFileName,
		Size:     file.Size,
	}

	// Simpan metadata ke database
	query := `
		INSERT INTO images (id, name, mime_type, url, width, height, size)
		VALUES ($1, $2, $3, $4, 0, 0, $5)
	`
	_, err = utils.DB.Exec(query, image.ID, image.Name, image.MimeType, image.URL, image.Size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert image metadata into the database"})
	}

	// Respon sukses
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Image uploaded and metadata saved successfully",
		"data":    image,
	})
}
