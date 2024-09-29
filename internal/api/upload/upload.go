package upload

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct{}

func (u *UploadHandler) GenerateSignature(c echo.Context) error {
	folder := c.Param("folder")

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	values := url.Values{}
	values.Set("folder", folder)
	values.Set("timestamp", timestamp)

	sig, err := api.SignParameters(values, os.Getenv("CLOUDINARY_SECRET"))
	if err != nil {
		return fmt.Errorf("Failed to generate signature: %v", err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"timestamp": timestamp,
		"signature": sig,
	})
}
