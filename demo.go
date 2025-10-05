package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "http://localhost:8080/tasks"

type Task struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

func main() {
	client := &http.Client{}

	// 1️⃣ Create tasks
	tasksToCreate := []string{"Learn Go", "Build API", "Test API"}
	for _, title := range tasksToCreate {
		task := Task{Title: title}
		body, _ := json.Marshal(task)
		resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Error creating task:", err)
			continue
		}
		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("Created task:", string(data))
	}

	// 2️⃣ List all tasks
	fmt.Println("\nListing all tasks:")
	resp, _ := http.Get(baseURL)
	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))

	// 3️⃣ Update first task
	update := Task{Completed: true}
	body, _ := json.Marshal(update)
	req, _ := http.NewRequest(http.MethodPut, baseURL+"/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	data, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("\nUpdated task 1:", string(data))

	// 4️⃣ Delete second task
	req, _ = http.NewRequest(http.MethodDelete, baseURL+"/2", nil)
	resp, _ = client.Do(req)
	resp.Body.Close()
	fmt.Println("\nDeleted task 2")

	// 5️⃣ Final list of tasks
	fmt.Println("\nFinal list of tasks:")
	resp, _ = http.Get(baseURL)
	data, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))
}
