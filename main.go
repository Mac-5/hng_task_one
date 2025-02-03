package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// NumberInfo represents the JSON response structure
type NumberInfo struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPerfect checks if a number is a perfect number
func isPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

// isArmstrong checks if a number is an Armstrong number
func isArmstrong(n int) bool {
	sum, temp, digits := 0, n, len(strconv.Itoa(n))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// sumOfDigits calculates the sum of a number's digits
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// getFunFact fetches a fun fact from Numbers API
func getFunFact(n int) string {
	resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d?json", n))
	if err != nil {
		return "Could not fetch fun fact"
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "Error parsing fun fact"
	}
	if text, ok := result["text"].(string); ok {
		return text
	}
	return "No fun fact available"
}

func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Get number from query parameters
	queryValues := r.URL.Query()
	numStr := queryValues.Get("number")

	if numStr == "" {
		http.Error(w, "Missing 'number' parameter", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	properties := []string{}
	if isPrime(num) {
		properties = append(properties, "prime")
	}
	if isPerfect(num) {
		properties = append(properties, "perfect")
	}
	if isArmstrong(num) {
		properties = append(properties, "armstrong")
	}
	if num%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	response := NumberInfo{
		Number:     num,
		IsPrime:    isPrime(num),
		IsPerfect:  isPerfect(num),
		Properties: properties,
		DigitSum:   sumOfDigits(num),
		FunFact:    getFunFact(num),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/api/classify-number", classifyNumberHandler)
	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
