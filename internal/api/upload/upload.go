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
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	params := url.Values{}
	params.Set("timestamp", timestamp)

	signature, err := api.SignParameters(params, os.Getenv("CLOUDINARY_SECRET"))
	if err != nil {
		return fmt.Errorf("Failed to generate signature: %v", err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"timestamp": timestamp,
		"signature": signature,
	})
}
