package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	DBUsername    string `yaml:"db-username"`
	DBPassword    string `yaml:"db-password"`
	DBName        string `yaml:"db-name"`
	PaypalGateway string `yaml:"paypal-gateway"`
	EnvName       string `yaml:"env-name"`
}

var settings = Settings{
	DBUsername:    "dev-user",
	DBPassword:    "dev-password",
	DBName:        "dev-db",
	PaypalGateway: "sandbox",
	EnvName:       "dev",
}

func loadSettings() {
	data, err := os.ReadFile("./settings.yaml")
	if err != nil {
		fmt.Println("No settings.yaml file found. Running with Dev profile")
		return
	}

	var s Settings
	if err := yaml.Unmarshal(data, &s); err != nil {
		fmt.Println("Error reading settings.yaml file. Running with Dev profile")
		return
	}

	settings = s
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am a container running inside Kubernetes.\n")
	fmt.Fprintf(w, "db-username: %s\n", settings.DBUsername)
	fmt.Fprintf(w, "db-password: %s\n", settings.DBPassword)
	fmt.Fprintf(w, "db-name: %s\n", settings.DBName)
	fmt.Fprintf(w, "paypal-gateway: %s\n", settings.PaypalGateway)
	fmt.Fprintf(w, "env-name: %s\n", settings.EnvName)
}

func main() {
	loadSettings()
	fmt.Printf("Loaded settings: %+v\n", settings)

	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
