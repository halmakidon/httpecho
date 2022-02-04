package main

type Config struct {
	Port           int    `default:"8080"`
	OidcJwtService string `default:"localhost:8081"`
}
