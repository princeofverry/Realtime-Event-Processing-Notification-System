package config

import "os"

type Config struct {
	RedisAddr string
}

// heap config
// ambil isi config
func Load() *Config {
	// ambil alamat redis
	return &Config{
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	} 
	return fallback
}