package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// Product struct to hold product information
type Product struct {
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
	Brand       string   `json:"brand"`
	Price       string   `json:"price"`
	Currency    string   `json:"currency"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	var product Product

	// On HTML element with the script type application/ld+json
	c.OnHTML("script[type='application/ld+json']", func(e *colly.HTMLElement) {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(e.Text), &data)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			return
		}

		// Extract relevant data
		if data["@type"] == "Product" {
			product.Name = data["name"].(string)
			product.URL = data["url"].(string)
			for _, img := range data["image"].([]interface{}) {
				product.Image = append(product.Image, img.(string))
			}
			product.Description = data["description"].(string)
			product.Brand = data["brand"].(map[string]interface{})["name"].(string)
			product.Price = data["offers"].(map[string]interface{})["price"].(string)
			product.Currency = data["offers"].(map[string]interface{})["priceCurrency"].(string)
		}
	})

	// URL of the page to scrape
	url := "https://thegrommet.com/product/kitchen/dbchopper"
	c.Visit(url)

	// Convert product data to JSON
	productJSON, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Println("Error marshaling product to JSON:", err)
		return
	}

	// Save to file
	file, err := os.Create("product.json")
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(productJSON)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Product information saved to product.json")
}
