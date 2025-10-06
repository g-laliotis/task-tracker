package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL  = "https://task-tracker-5cg1.onrender.com"
	email    = "test@example.com" // 👈 change this if needed
	password = "password123"      // 👈 change this if needed
)

type Task struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func main() {
	client := &http.Client{}

	fmt.Println("🚀 Starting Task Tracker live Render demo...")

	// 1️⃣ Login to get token
	fmt.Println("🔑 Logging in...")
	loginBody := map[string]string{"email": email, "password": password}
	body, _ := json.Marshal(loginBody)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Sprintf("❌ Login failed: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		panic(fmt.Sprintf("❌ Login error: %s", data))
	}

	var loginResp LoginResponse
	json.NewDecoder(resp.Body).Decode(&loginResp)
	token := loginResp.Token
	fmt.Println("✅ Logged in successfully.")
	fmt.Println("🔹 JWT token:", token)

	// 2️⃣ Create tasks
	fmt.Println("🧩 Creating tasks...")
	tasksToCreate := []string{"Learn Go", "Build API", "Deploy on Render"}
	for _, title := range tasksToCreate {
		task := Task{Title: title}
		body, _ := json.Marshal(task)
		req, _ := http.NewRequest(http.MethodPost, baseURL+"/tasks", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("❌ Error creating task:", err)
			continue
		}
		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("✅ Created:", string(data))
	}

	// 3️⃣ List all tasks
	fmt.Println("\n📋 Listing all tasks:")
	req, _ := http.NewRequest(http.MethodGet, baseURL+"/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, _ = client.Do(req)
	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))

	// 4️⃣ Update first task
	fmt.Println("\n🟢 Marking task 1 as completed...")
	update := Task{Completed: true}
	body, _ = json.Marshal(update)
	req, _ = http.NewRequest(http.MethodPut, baseURL+"/tasks/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, _ = client.Do(req)
	data, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ Updated:", string(data))

	// 5️⃣ Delete second task
	fmt.Println("\n🗑️ Deleting task 2...")
	req, _ = http.NewRequest(http.MethodDelete, baseURL+"/tasks/2", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, _ = client.Do(req)
	data, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ Deleted:", string(data))

	// 6️⃣ Final list
	fmt.Println("\n📦 Final list of tasks:")
	req, _ = http.NewRequest(http.MethodGet, baseURL+"/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, _ = client.Do(req)
	data, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))

	fmt.Println("\n🎉 Demo completed successfully!")
}
