package conf

import (
	"log"
	"os"
)

const UserIdHeader = "X-User-Identity"

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
var JwtFilterPort = os.Getenv("JWT_FILTER_PORT")
var DevMode = os.Getenv("DEV_MODE")

func init() {
	if DevMode == "prod" {
		if JwtSecretKey == "" {
			log.Fatalln("`JWT_SECRET_KEY` must be specified with non empty value!")
		}

		if JwtFilterPort == "" {
			log.Fatalln("`JWT_FILTER_PORT` must be specified with non empty value!")
		}
	} else if DevMode == "dev" {
		if JwtSecretKey == "" {
			JwtSecretKey = "kXpBmV^_|BFq#c.-\"\"B:cd#k6-/EuVp]"
			log.Println("`JWT_SECRET_KEY` must be specified with non empty value! default=" + JwtSecretKey)
		}

		if JwtFilterPort == "" {
			JwtFilterPort = "8086"
			log.Println("`JWT_FILTER_PORT` must be specified with non empty value! default=" + JwtFilterPort)
		}
	} else if DevMode == "test" {
		// Test tokens had been generated with specified secret key.
		JwtSecretKey = "kXpBmV^_|BFq#c.-\"\"B:cd#k6-/EuVp]"
		JwtFilterPort = "8086"
	}
}
