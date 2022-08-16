package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VrMolodyakov/url-shortener/internal/config"
	"github.com/VrMolodyakov/url-shortener/internal/service/shortener"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := config.GetConfig()
	str := shortener.Encode(15151616161)
	n := shortener.Decode(str)
	fmt.Println(n)
	fmt.Println(cfg)
}
