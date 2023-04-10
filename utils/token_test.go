package utils_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/anyingiit/GoReactResourceManagement/utils"
)

func TestGenerateTokenAndParseToken(t *testing.T) {
	jsonData, err := json.Marshal(map[string]string{"id": "1234", "username": "testuser"})
	if err != nil {
		t.Fatalf("Failed to marshal json data: %v", err)
	}

	// Generate a token
	customData := string(jsonData)
	issuer := "myIssuer"
	validDuration := 10 * time.Minute
	signingKey := "mySigningKey"
	tokenString, err := utils.GenerateToken(customData, issuer, validDuration, signingKey)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Parse the token
	d, err := utils.ParseToken(tokenString, signingKey)
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	parsedCustomData, ok := d.(string)
	if !ok {
		t.Fatalf("Failed to parse custom data")
	}

	// Check if the parsed custom data matches the original custom data
	if !reflect.DeepEqual(parsedCustomData, customData) {
		t.Errorf("Parsed custom data %v does not equal expected custom data %v", parsedCustomData, customData)
	}
}
