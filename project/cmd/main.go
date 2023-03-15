package main

import (
	v1 "github.com/anotherhope/FizzBuzz/project/api/v1"
	"github.com/anotherhope/FizzBuzz/project/pkg/server"
)

func main() {
	srv := server.Create()
	srv.API(v1.STATUS)
	srv.API(v1.FIZZBUZZ)
	srv.Start()
}
