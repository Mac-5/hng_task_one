package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"unicode"
)

// NumberInfo represents the JSON response structure for valid numbers
type NumberInfo struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

// ErrorResponse represents the JSON response structure for errors
type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
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
func isAlphabetic(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

// isNumeric checks if a string is a valid integer
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}

// getFunFact fetches a fun fact from Numbers API
func getFunFact(n int) string {
	resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d/math?json", n))
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

// respondWithError sends a JSON error response
func respondWithError(w http.ResponseWriter, numString string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(ErrorResponse{Number: numString, Error: true}); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

// classifyNumberHandler handles requests to /api/classify-number
func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the 'number' query parameter
	numStr := r.URL.Query().Get("number")

	if numStr == "" {
		respondWithError(w, "missing")
		return
	}

	if isAlphabetic(numStr) {
		respondWithError(w, "alphabet")
		return
	}

	if !isNumeric(numStr) {
		respondWithError(w, "invalid")
		return
	}

	num, _ := strconv.Atoi(numStr)

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
