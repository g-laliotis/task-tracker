package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Task struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

func main() {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/tasks" // default for local
	}

	client := &http.Client{}

	// 1️⃣ Create tasks
	tasksToCreate := []string{"Learn Go", "Build API", "Test API"}
	fmt.Println("🪄 Creating tasks...")
	for _, title := range tasksToCreate {
		task := Task{Title: title}
		body, _ := json.Marshal(task)

		resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("❌ Error creating task:", err)
			continue
		}
		handleResponse(resp)
	}

	// 2️⃣ List all tasks
	fmt.Println("\n📋 Listing all tasks:")
	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("❌ Error getting tasks:", err)
		return
	}
	handleResponse(resp)

	// 3️⃣ Update first task
	fmt.Println("\n🔄 Updating task 1 as completed...")
	update := Task{Completed: true}
	body, _ := json.Marshal(update)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/1", baseURL), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("❌ Error updating task:", err)
		return
	}
	handleResponse(resp)

	// 4️⃣ Delete second task
	fmt.Println("\n🗑️ Deleting task 2...")
	req, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/2", baseURL), nil)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("❌ Error deleting task:", err)
		return
	}
	handleResponse(resp)

	// 5️⃣ Final list of tasks
	fmt.Println("\n📦 Final list of tasks:")
	resp, err = http.Get(baseURL)
	if err != nil {
		fmt.Println("❌ Error getting tasks:", err)
		return
	}
	handleResponse(resp)
}

func handleResponse(resp *http.Response) {
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println(string(data))
	} else {
		fmt.Printf("⚠️  Status %d: %s\n", resp.StatusCode, string(data))
	}
}
