package initializers

import "github.com/joho/godotenv"

func LoadEnvVars() {
	godotenv.Load()
}
