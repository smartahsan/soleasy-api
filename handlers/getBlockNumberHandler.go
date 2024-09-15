// handlers/user_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlockNumber(c *gin.Context) {

	url := "https://public-api.solanabeach.io/v1/latest-blocks?limit=1"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("If-None-Match", `W/"aa8-FWgvG9s4kXEyrgfi70cigsvXXM8"`)
	req.Header.Set("Origin", "https://solanabeach.io")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://solanabeach.io/")
	req.Header.Set("Sec-CH-UA", `"Chromium";v="128", "Not;A=Brand";v="24", "Google Chrome";v="128"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error making request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transaction", "err": err})
		return
	}

	defer resp.Body.Close()

	statusCode := resp.StatusCode

	switch statusCode {
	case 200:
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Printf("Error reading response: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read transaction data"})
			return
		}

		var bodyMapped GetBlockNumberResponse

		json.Unmarshal(body, &bodyMapped)

		blockNumber := bodyMapped[0].Blocknumber

		c.JSON(http.StatusOK, gin.H{"blockNumber": blockNumber})

	case 429:
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch transaction", "details": "Rate limit exceeded"})
		return
	default:
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch transaction", "details": "Unknown error", "statusCode": statusCode})
		return
	}
}
