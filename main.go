package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

// List of backend servers (Replace with your actual backend URLs)
var backends = []string{
	"http://localhost:8081",
	"http://localhost:8082",
	"http://localhost:8083",
}

var currentIndex uint64 = 0

// Function to get the next backend using Round-Robin
func getNextBackend() string {
	index := atomic.AddUint64(&currentIndex, 1) % uint64(len(backends))
	return backends[index]
}

// Reverse proxy function
func reverseProxy(c *gin.Context) {
	targetURL := getNextBackend()
	remote, err := url.Parse(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid backend URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	c.Request.URL.Host = remote.Host
	c.Request.URL.Scheme = remote.Scheme
	c.Request.Host = remote.Host

	proxy.ModifyResponse = func(resp *http.Response) error {
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
		defer resp.Body.Close()
	
		// Print the response body
		fmt.Println("Response from", targetURL, ":", string(body))
	
		// Reset the response body so it can be read again
		resp.Body = io.NopCloser(bytes.NewReader(body))
		return nil
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()
	r.Any("/*proxyPath", reverseProxy)

	fmt.Println("Load Balancer running on port 8080...")
	r.Run(":8080") 
}
