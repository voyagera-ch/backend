package imageHandler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/voyagera/backend/auth"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username, err := auth.ValidateRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	username = username

	// Get the image path
	imagePath := "./image-storage/gopher.png"

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Image not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file extension to set the correct content type
	ext := filepath.Ext(imagePath)
	var contentType string
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	default:
		http.Error(w, "Unsupported image type.", http.StatusUnsupportedMediaType)
		return
	}

	// Set the content type
	w.Header().Set("Content-Type", contentType)

	// Serve the file
	http.ServeFile(w, r, imagePath)
}
