// Package main is the entry point for the application.
package main

import (
	"github.com/anotherhope/FizzBuzz/project/config"
	v1 "github.com/anotherhope/FizzBuzz/project/internal/api/v1"
	"github.com/anotherhope/FizzBuzz/project/internal/lib"
	"github.com/anotherhope/FizzBuzz/project/pkg/server"
	"github.com/spf13/cobra"
)

func main() {
	config.Helper.Run = func(cmd *cobra.Command, args []string) {
		srv := server.Create()
		srv.API(v1.Status)
		srv.API(v1.Fizzbuzz)
		srv.Start()
	}

	if err := config.Helper.Execute(); err != nil {
		lib.WithCriticalError(err)
	}
}
