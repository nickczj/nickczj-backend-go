package config

import (
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	log "github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// GetSecret retrieves a secret, first checking environment variables,
// then falling back to GCP Secret Manager.
// For env vars, extracts the secret name from the GCP path (e.g., "db_password" from ".../secrets/db_password/...")
func GetSecret(name string) (*string, error) {
	// Extract secret name from GCP path for env var lookup
	// e.g., "projects/123/secrets/db_password/versions/latest" -> "DB_PASSWORD"
	parts := strings.Split(name, "/")
	var envKey string
	for i, part := range parts {
		if part == "secrets" && i+1 < len(parts) {
			envKey = strings.ToUpper(parts[i+1])
			break
		}
	}

	// Check environment variable first
	if envKey != "" {
		if val := os.Getenv(envKey); val != "" {
			log.Info("Using secret from environment: ", envKey)
			return &val, nil
		}
	}

	// Fall back to GCP Secret Manager
	return AccessSecretVersion(name)
}

// AccessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
func AccessSecretVersion(name string) (*string, error) {
	// name := "projects/my-project/secrets/my-secret/versions/5"
	// name := "projects/my-project/secrets/my-secret/versions/latest"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, err
	}

	secret := string(result.Payload.Data)
	return &secret, nil
}
