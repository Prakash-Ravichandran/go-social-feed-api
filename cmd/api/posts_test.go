package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
)

type UpdatePostPayloadTest struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func updatePost(user string, postID int, p UpdatePostPayloadTest, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()

	// Construct the URL for the update endpoint
	url := fmt.Sprintf("http://localhost:3000/posts/%d", postID)

	// Create the JSON payload
	b, _ := json.Marshal(p)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers as needed
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	t.Logf("[%s] Update status: %s", user, resp.Status)
}

func randomString(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "fallback"
	}
	return hex.EncodeToString(bytes)
}

func TestUpdatePost(t *testing.T) {
	var wg sync.WaitGroup

	// Assuming the post ID to update is 13
	postID := 13

	// Simulate User A and User B updating the same post concurrently
	wg.Add(2)

	// Generate dynamic/random title and content for each test run
	title := fmt.Sprintf("TITLE FROM USER A [%s]", randomString(4))
	content := fmt.Sprintf("CONTENT FROM USER B [%s]", randomString(6))

	// User A updates only the Title
	go updatePost("USER A", postID, UpdatePostPayloadTest{Title: &title}, &wg, t)

	// User B updates only the Content
	go updatePost("USER B", postID, UpdatePostPayloadTest{Content: &content}, &wg, t)

	wg.Wait()
}
