package main

import (
	hello "example/hellopkg" // with alias
	"github.com/rs/zerolog/log"
)

func main() {
	// hellopkg.Hello() // without alias
	hello.Hello()
	log.Info().Msg("Hello, World!")
}
