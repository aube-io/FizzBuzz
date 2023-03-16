package config

import "github.com/anotherhope/FizzBuzz/project/internal/docs"

func init() {
	docs.SwaggerInfo.Title = "FizzBuzz"
	docs.SwaggerInfo.Description = "FizzBuzz API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
