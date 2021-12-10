package http

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
)

// if in original source no file name, used this
const DEFAULT_FILE_NAME = "products.csv"

// TryDownloadFile
func TryDownloadFile(fileSaveDir string, downloadUrl string) (string, error) {
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return "", fmt.Errorf("download file by url: %w/n", err)
	}
	defer resp.Body.Close() // nolint

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	disposition, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return "", fmt.Errorf("parse media type fail: %w/n", err)
	}
	if disposition != "attachment" {
		return "", fmt.Errorf("url hasn't atteched file, header: %s", contentDisposition)
	}
	originalFileName := params["filename"]
	if originalFileName == "" {
		originalFileName = DEFAULT_FILE_NAME
	}

	fileSavePath := fmt.Sprintf("%s/%s", fileSaveDir, originalFileName)
	if _, err := os.Stat(fileSaveDir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(fileSaveDir, 0755)
		if mkdirErr != nil {
			return "", fmt.Errorf("create filder by path fail: %w/n", mkdirErr)
		}
	}

	out, err := os.Create(fileSavePath)
	if err != nil {
		return "", fmt.Errorf("create file fail: %w/n", err)
	}
	defer out.Close() // nolint

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("save file fail: %w/n", err)
	}
	return originalFileName, nil
}
