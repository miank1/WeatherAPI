package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Weather API")

	// Note: You'll need to replace YOUR_API_KEY with your actual Visual Crossing API key
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/delhi?unitGroup=us&key=X2AM8B666SQSMPZNKMD7QBEAE&contentType=json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	// Start HTTP server
	http.HandleFunc("/weather", weatherHandler)
	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Weather API Server is running!")
}
