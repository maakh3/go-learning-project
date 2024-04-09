package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type healthResponse struct {
	Status        string `json:"status"`
	Version       string `json:"version"`
	WorkerVersion string `json:"workerversion"`
	CommitHash    string `json:"commithash"`
}

func main() {
	resp, err := http.Get("https://r10clientv2uat.mobilefarm-nonprod.co.uk/health")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var health healthResponse
	json.Unmarshal(body, &health)

	fmt.Println("Raw Response\n", body)
	fmt.Println("\nMarshaled Response\n", health)
	fmt.Println("\nHealth.Status\n", health.Status)
	fmt.Println("\nHealth.Version\n", health.Version)
	fmt.Println("\nHealth.WorkerVersion\n", health.WorkerVersion)
	fmt.Println("\nHealth.CommitHash\n", health.CommitHash)
}
