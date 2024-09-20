package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Name string `json:"Name"`
	Role string `json:"Role"`
	Seed string `json:"Seed"`
}

func isPrime(numStr string) bool {
	num, ok := new(big.Int).SetString(numStr, 10)
	if !ok {
		return false
	}
	return num.ProbablyPrime(0)
}

func validateJWT(tokenString string) bool {
	// Split the JWT into parts (header, payload, signature)
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return false
	}

	// Parse and validate the token
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return false
	}

	// Extract the claims
	claims := token.Claims.(jwt.MapClaims)

	// Check if there are exactly 3 claims
	if len(claims) != 3 {
		return false
	}

	// Validate Name
	name, ok := claims["Name"].(string)
	if !ok || len(name) > 256 || regexp.MustCompile(`[0-9]`).MatchString(name) {
		return false
	}

	// Validate Role
	role, ok := claims["Role"].(string)
	if !ok || (role != "Admin" && role != "Member" && role != "External") {
		return false
	}

	// Validate Seed
	seed, ok := claims["Seed"].(string)
	if !ok || !isPrime(seed) {
		return false
	}

	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.URL.Query().Get("token")
	if jwtToken == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	if validateJWT(jwtToken) {
		fmt.Fprintf(w, "true")
	} else {
		fmt.Fprintf(w, "false")
	}
}

func main() {
	http.HandleFunc("/validate", handler)
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
