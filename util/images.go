package util

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func GetImage(path string) image.Image {
	var b []byte
	var photo image.Image

	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	_, err = base64.StdEncoding.Decode(b, content)
	if err != nil {
		b = content
	}

	ext := strings.Split(path, ".")[len(strings.Split(path, "."))-1]
	if ext == "png" {
		photo, err = png.Decode(bytes.NewReader(b))
	} else if ext == "jpeg" || ext == "jpg" {
		photo, err = jpeg.Decode(bytes.NewReader(b))
	} else if ext == "gif" {
		photo, err = gif.Decode(bytes.NewReader(b))
	} else {
		return nil
	}

	if err != nil {
		return nil
	}

	return photo
}
