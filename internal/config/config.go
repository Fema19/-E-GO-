package config

import "os"

func JWTSecret() string {

	if v := os.Getenv("JWT_SECRET"); v != "" {
		return v
	}

	return "dev_super_secret_change_me"
}
