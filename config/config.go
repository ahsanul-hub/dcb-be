package config

import "github.com/joho/godotenv"

// Config func to get env value
// func Config(key string) string {
// 	// load .env file
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Print("Error loading .env file")
// 	}
// 	return os.Getenv(key)
// }


var Env map[string]string

func Config(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := "../.env"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		panic(err)
	}

}