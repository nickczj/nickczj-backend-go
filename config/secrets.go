package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

var vaultClient *api.Client

// InitVault initializes the Vault client if VAULT_ADDR is set
func InitVault() error {
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		log.Info("VAULT_ADDR not set, using environment variables for secrets")
		return nil
	}

	config := api.DefaultConfig()
	config.Address = vaultAddr

	client, err := api.NewClient(config)
	if err != nil {
		return fmt.Errorf("failed to create Vault client: %w", err)
	}

	// Set token from environment
	token := os.Getenv("VAULT_TOKEN")
	if token != "" {
		client.SetToken(token)
	}

	vaultClient = client
	log.Info("Vault client initialized: ", vaultAddr)
	return nil
}

// GetSecret retrieves a secret, checking sources in order:
// 1. Environment variable (e.g., DB_PASSWORD)
// 2. HashiCorp Vault (if configured)
func GetSecret(name string) (*string, error) {
	// Extract secret key name
	// Supports both simple names ("db_password") and paths ("secret/data/myapp/db_password")
	secretKey := extractSecretKey(name)
	envKey := strings.ToUpper(secretKey)

	// 1. Check environment variable first
	if val := os.Getenv(envKey); val != "" {
		log.Debug("Using secret from environment: ", envKey)
		return &val, nil
	}

	// 2. Try Vault if configured
	if vaultClient != nil {
		return getSecretFromVault(name, secretKey)
	}

	return nil, fmt.Errorf("secret %s not found in environment or Vault", secretKey)
}

// getSecretFromVault retrieves a secret from Vault
// Supports KV v2 secrets engine format: secret/data/path
func getSecretFromVault(path string, key string) (*string, error) {
	// Ensure path uses KV v2 format
	if !strings.Contains(path, "/data/") {
		// Convert "secret/myapp" to "secret/data/myapp"
		parts := strings.SplitN(path, "/", 2)
		if len(parts) == 2 {
			path = parts[0] + "/data/" + parts[1]
		}
	}

	secret, err := vaultClient.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret from Vault: %w", err)
	}

	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("secret not found at path: %s", path)
	}

	// KV v2 wraps data in a "data" key
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		// Try direct access for KV v1
		data = secret.Data
	}

	// Look for the specific key or common key names
	keysToTry := []string{key, "value", "password", "secret"}
	for _, k := range keysToTry {
		if val, ok := data[k].(string); ok {
			log.Debug("Retrieved secret from Vault: ", path)
			return &val, nil
		}
	}

	return nil, fmt.Errorf("key %s not found in secret at path: %s", key, path)
}

// extractSecretKey extracts the secret key from various path formats
func extractSecretKey(name string) string {
	// Handle GCP-style paths: "projects/123/secrets/db_password/versions/latest"
	if strings.Contains(name, "/secrets/") {
		parts := strings.Split(name, "/")
		for i, part := range parts {
			if part == "secrets" && i+1 < len(parts) {
				return parts[i+1]
			}
		}
	}

	// Handle Vault-style paths: "secret/data/myapp/db" -> "db"
	// or simple paths: "secret/myapp" -> "myapp"
	parts := strings.Split(name, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}

	return name
}
