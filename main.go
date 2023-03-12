// package main is the main package
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// import "github.com/timtoronto634/resas-cli/cmd"

func main() {
	// cmd.Execute()
	apiEndpoint := os.Getenv("RESAS_API_ENDPOINT")
	if apiEndpoint == "" {
		log.Println("failed in retrieving RESAS_API_ENDPOINT. please set RESAS_API_ENDPOINT in environment variable")
	}
	apiKey := os.Getenv("RESAS_API_KEY")
	if apiKey == "" {
		log.Println("failed in retrieving API_KEY. please set RESAS_API_KEY in environment variable")
	}
	client := &http.Client{}
	apiPath := "/api/v1/population/composition/perYear"
	req, err := http.NewRequest("GET", apiEndpoint+apiPath, nil)
	if err != nil {
		log.Printf("failed in creating request: %v", err)
	}
	values := req.URL.Query()
	values.Add("cityCode", "-")
	values.Add("prefCode", "13")
	req.URL.RawQuery = values.Encode()
	log.Printf("requesting %v\n", req.URL)

	req.Header.Add("X-API-KEY", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed in requesting: %v", err)
	}
	fmt.Println(resp.Status)
	io.Copy(os.Stdout, resp.Body)
	log.Println()

}
