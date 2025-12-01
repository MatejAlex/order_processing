package cfg

import "os"

type Config struct {
	Port     string
	AuthURL  string
	AuditURL string
	DBHost   string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() Config {
	return Config{
		Port:     getEnv("PORT", "8080"),
		AuthURL:  getEnv("AUTH_SERVICE_URL", "http://auth-service:8081"),
		AuditURL: getEnv("AUDIT_SERVICE_URL", "http://audit-service:8082"),
		DBHost:   getEnv("DB_HOST", "localhost"),
	}
}
