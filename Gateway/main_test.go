package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

// Student represents the student data structure.
type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	College struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"college"`
	Email []string `json:"email"`
}

// aUrlRegister is the URL for registering a student.
const aUrlRegister = "http://127.0.0.1:8888/gateway/StudentService/Register"
const aUrlEcho = "http://localhost:8888/gateway/EchoService/Echo"

// generateRandomStudent generates a random student for testing.
func generateRandomStudent() Student {
	st := Student{
		ID:   rand.Intn(1000) + 1 + 32,
		Name: fmt.Sprintf("Student%d", rand.Intn(100)),
		College: struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		}{
			Name:    fmt.Sprintf("College%d", rand.Intn(10)),
			Address: fmt.Sprintf("Address%d", rand.Intn(100)),
		},
	}
	emailCount := rand.Intn(5) + 1 // Generate a random number between 1 and 5
	st.Email = make([]string, emailCount)
	for i := 0; i < emailCount; i++ {
		st.Email[i] = fmt.Sprintf("email%d@test.com", rand.Intn(100))
	}

	return st
}

type RegisterResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// BenchmarkStudentServiceRegister tests the performance of the StudentService Register function.
func BenchmarkStudentServiceRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		student := generateRandomStudent()

		bytesData, err := json.Marshal(student)
		if err != nil {
			b.Fatalf("Error marshalling data: %v", err)
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", aUrlRegister, reader)

		if err != nil {
			b.Fatalf("Error creating request: %v", err)
		}
		defer request.Body.Close()

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		startTime := time.Now()
		resp, err := client.Do(request)
		if err != nil {
			b.Fatalf("Error sending request: %v", err)
		}
		elapsed := time.Since(startTime)

		defer resp.Body.Close()

		var response RegisterResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			b.Fatalf("Error decoding response: %v", err)
		}

		if !response.Success {
			b.Fatalf("Registration failed for student: %+v", student)
		}

		b.Logf("Student registration successful: %+v (Elapsed time: %v)", student, elapsed)
	}
}

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message string `json:"message"`
}

// BenchmarkEchoService tests the performance of the EchoService Echo function.
func BenchmarkEchoService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Create a random message for testing
		message := fmt.Sprintf("hello%d, Cloudwego!", rand.Intn(100))

		// Create the request payload
		request := EchoRequest{
			Message: message,
		}

		bytesData, err := json.Marshal(request)
		if err != nil {
			b.Fatalf("Error marshalling data: %v", err)
		}

		reader := bytes.NewReader(bytesData)
		request_2, err := http.NewRequest("POST", aUrlEcho, reader)
		if err != nil {
			b.Fatalf("Error creating request: %v", err)
		}
		defer request_2.Body.Close()

		request_2.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		startTime := time.Now()
		resp, err := client.Do(request_2)
		if err != nil {
			b.Fatalf("Error sending request: %v", err)
		}
		elapsed := time.Since(startTime)

		defer resp.Body.Close()

		var response EchoResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			b.Fatalf("Error decoding response: %v", err)
		}

		if response.Message != message {
			b.Fatalf("Unexpected response: expected %q, got %q", message, response.Message)
		}

		b.Logf("EchoService successful: Request=%q, Response=%q (Elapsed time: %v)", message, response.Message, elapsed)
	}
}
