package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "http://localhost:8080"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

func main() {
	client := &http.Client{}

	// 1️⃣ Sign up a user
	user := User{Email: "test@example.com", Password: "password123"}
	fmt.Println("📝 Signing up user...")
	postJSON(client, "/signup", user, "")

	// 2️⃣ Log in and get JWT token
	fmt.Println("🔑 Logging in...")
	resp, _ := postJSON(client, "/login", user, "")
	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var result map[string]string
	json.Unmarshal(data, &result)
	token := result["token"]
	fmt.Println("✅ Token received:", token)

	// 3️⃣ Create tasks
	tasksToCreate := []string{"Learn Go", "Build API", "Test API"}
	for _, title := range tasksToCreate {
		task := Task{Title: title}
		postJSON(client, "/tasks", task, token)
	}

	// 4️⃣ List all tasks
	fmt.Println("\n📋 Listing all tasks:")
	getJSON(client, "/tasks", token)

	// 5️⃣ Update first task
	update := Task{Completed: true}
	putJSON(client, "/tasks/1", update, token)

	// 6️⃣ Delete second task
	deleteReq(client, "/tasks/2", token)

	// 7️⃣ Final list of tasks
	fmt.Println("\n📦 Final list of tasks:")
	getJSON(client, "/tasks", token)
}

func postJSON(client *http.Client, path string, body interface{}, token string) (*http.Response, error) {
	url := baseURL + path
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	respData, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respData))
	resp.Body.Close()
	return resp, err
}

func getJSON(client *http.Client, path, token string) {
	url := baseURL + path
	req, _ := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))
}

func putJSON(client *http.Client, path string, body interface{}, token string) {
	url := baseURL + path
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	respData, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(respData))
}

func deleteReq(client *http.Client, path, token string) {
	url := baseURL + path
	req, _ := http.NewRequest("DELETE", url, nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deleted task:", path)
	resp.Body.Close()
}
