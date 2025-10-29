package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name          string
		password      string
		hash          string
		wantErr       bool
		matchPassword bool
	}{
		{
			name:          "Correct password",
			password:      password1,
			hash:          hash1,
			wantErr:       false,
			matchPassword: true,
		},
		{
			name:          "Incorrect password",
			password:      "wrongPassword",
			hash:          hash1,
			wantErr:       false,
			matchPassword: false,
		},
		{
			name:          "Password doesn't match different hash",
			password:      password1,
			hash:          hash2,
			wantErr:       false,
			matchPassword: false,
		},
		{
			name:          "Empty password",
			password:      "",
			hash:          hash1,
			wantErr:       false,
			matchPassword: false,
		},
		{
			name:          "Invalid hash",
			password:      password1,
			hash:          "invalidhash",
			wantErr:       true,
			matchPassword: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && match != tt.matchPassword {
				t.Errorf("CheckPasswordHash() expects %v, got %v", tt.matchPassword, match)
			}
		})
	}
}

func TestMakeAndValidateJWT(t *testing.T) {
	tokenSecret1 := "my-super-secret-key"
	tokenSecret2 := "another-secret-key"
	userID := uuid.New()

	t.Run("Valid token", func(t *testing.T) {
		tokenString, err := MakeJWT(userID, tokenSecret1, time.Hour)
		if err != nil {
			t.Fatalf("MakeJWT() returned an unexpected error: %v", err)
		}

		validatedUserID, err := ValidateJWT(tokenString, tokenSecret1)
		if err != nil {
			t.Fatalf("ValidateJWT() returned an unexpected error: %v", err)
		}

		if validatedUserID != userID {
			t.Errorf("ValidateJWT() returned wrong userID. got %v, want %v", validatedUserID, userID)
		}
	})

	t.Run("Invalid signature", func(t *testing.T) {
		tokenString, err := MakeJWT(userID, tokenSecret1, time.Hour)
		if err != nil {
			t.Fatalf("MakeJWT() returned an unexpected error: %v", err)
		}

		_, err = ValidateJWT(tokenString, tokenSecret2)
		if err == nil {
			t.Error("ValidateJWT() expected an error for invalid signature, but got nil")
		}
	})

	t.Run("Expired token", func(t *testing.T) {
		// Create a token that expired an hour ago
		tokenString, err := MakeJWT(userID, tokenSecret1, -time.Hour)
		if err != nil {
			t.Fatalf("MakeJWT() returned an unexpected error: %v", err)
		}

		_, err = ValidateJWT(tokenString, tokenSecret1)
		if err == nil {
			t.Error("ValidateJWT() expected an error for expired token, but got nil")
		}
	})

	t.Run("Malformed token", func(t *testing.T) {
		_, err := ValidateJWT("this.is.not.a.jwt", tokenSecret1)
		if err == nil {
			t.Error("ValidateJWT() expected an error for malformed token, but got nil")
		}
	})
}
