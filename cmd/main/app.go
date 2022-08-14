package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VrMolodyakov/url-shortener/internal/config"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := config.GetConfig()
	fmt.Println(cfg)
}
