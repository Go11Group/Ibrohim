package handler

import (
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Client(c *gin.Context, port string) ([]byte, error) {
	method := c.Request.Method
	url := c.Request.URL.Path
	body := c.Request.Body

	client := http.Client{}
	req, err := http.NewRequest(method, "http://localhost:808"+port+url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return resBody, nil
}
