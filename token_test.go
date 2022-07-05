package peg_go_client

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("failed to load .env.local")
	}
	email := os.Getenv("PEG_EMAIL")
	password := os.Getenv("PEG_PASSWORD")
	idToken, err := GetAuthToken(email, password)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("idToken: %v", idToken)
}
