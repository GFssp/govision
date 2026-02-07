package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func SendRequest(url string, imageBytes *bytes.Buffer, apiKey string) (error, []byte) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	encodedImage := base64.StdEncoding.EncodeToString(imageBytes.Bytes())
	if err := writer.WriteField("image", encodedImage); err != nil {
		return fmt.Errorf("Error writing image field: %v", err), nil
	}

	writer.Close()
	fullURL := fmt.Sprintf("%s?key=%s", url, apiKey)
	req, err := http.NewRequest("POST", fullURL, &requestBody)
	if err != nil {
		return fmt.Errorf("Error setting request to %v, %v", url, err), nil
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request to %v: %v", url, err), nil
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %v", err), nil
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error calling URL (status %d): %v", resp.StatusCode, string(respBytes)), nil
	}

	return nil, respBytes
}
