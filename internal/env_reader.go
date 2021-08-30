package internal

import "github.com/joho/godotenv"

func GetEnvVariableByName(name string) (string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	if err != nil {
		return "", err
	}

	return myEnv[name], nil
}
