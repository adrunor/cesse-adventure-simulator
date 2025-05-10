package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Impl struct {
}

func (env *Impl) Get(key string) (string, error) {
	value := os.Getenv(key)

	return value, nil
}

func NewEnv() *Impl {
	load()
	return &Impl{}
}

func load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
