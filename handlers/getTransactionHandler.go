// handlers/user_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {
	signature := c.Param("signature")

	url := fmt.Sprintf("https://public-api.solanabeach.io/v1/transaction/%s", signature)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://solanabeach.io")
	req.Header.Set("Referer", "https://solanabeach.io/")

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

		bodyType := ""

		var bodySPL GetTransactionSPLResponse
		var bodySolana GetTransactionSolanaResponse

		// Attempt to map onto the response of a SPL transfer
		if err := json.Unmarshal(body, &bodySPL); err == nil {
			// Mapped to a regular Solana SPL response
			for _, account := range bodySPL.Accounts {

				if account.Account.Name == "SPL Token Program" {
					bodyType = "spl"
					break
				}
			}
		} else {
			fmt.Printf("SPL parsing error: %v\n", err)
		}

		// If the response is not an SPL transfer, attempt to map onto the response of a regular Solana transfer
		if bodyType == "" {
			if err := json.Unmarshal(body, &bodySolana); err == nil {
				// Mapped to a regular Solana transfer response
				bodyType = "solana"
			} else {
				fmt.Printf("Solana parsing error: %v\n", err)
			}
		}

		if bodyType == "solana" {

			response := gin.H{
				"type":              "solana",
				"tokenName":         "SOL",
				"from":              bodySolana.Accounts[0].Account.Address,
				"to":                bodySolana.Accounts[1].Account.Address,
				"amount":            float64(bodySolana.Instructions[2].Parsed.Transfer.Lamports) / 1000000000,
				"blockNumber":       bodySolana.BlockNumber,
				"transactionStatus": bodySolana.Meta.Status.Ok == nil,
				"blockTime":         bodySolana.Blocktime.Absolute,
			}

			c.JSON(http.StatusOK, response)

		} else if bodyType == "spl" {

			response := gin.H{
				"type":              "spl",
				"tokenName":         bodySPL.Accounts[4].Account.Name,
				"from":              bodySPL.Meta.PostTokenBalances[1].Owner.Address,
				"to":                bodySPL.Meta.PostTokenBalances[0].Owner.Address,
				"amount":            float64(bodySPL.Instructions[2].Parsed.TransferChecked.Amount) / math.Pow(10, float64(bodySPL.Instructions[2].Parsed.TransferChecked.Decimals)),
				"blockNumber":       bodySPL.BlockNumber,
				"transactionStatus": bodySPL.Meta.Status.Ok == nil,
				"blockTime":         bodySPL.Blocktime.Absolute,
			}

			c.JSON(http.StatusOK, response)
		} else if bodyType == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse transaction data"})
		}

	case 429:
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch transaction", "details": "Rate limit exceeded"})
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response: %v", err)
			return
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return
	default:
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch transaction", "details": "Unknown error", "statusCode": statusCode})
		return
	}
}
