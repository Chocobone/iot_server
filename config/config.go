package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env  string `env:"ENV" default:"dev"`
	Port int    `env:"PORT" default:"8080"`

	// Home Assistant
	Home_host string `env:"HOME_URL" default:"127.0.0.1`
	Home_port int    `env:"HOME_PORT" default:"8123"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

/* func main() {
    url := "http://localhost:8123/api/services/vacuum/return_to_base"
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJkYjYwYWJkOTRlN2M0YTZjODkyMzQ3Y2JjOTgzZWUxYSIsImlhdCI6MTc0NzAyMTI5NCwiZXhwIjoyMDYyMzgxMjk0fQ.7mybkEqIh7coIRrVxkno8I1iTXCDz5wipB9rpomVUB0"
    payload := []byte(`{"entity_id": "vacuum.robosceongsogi"}`)

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
} */
