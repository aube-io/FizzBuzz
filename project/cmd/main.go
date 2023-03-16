package main

import (
	"github.com/anotherhope/FizzBuzz/project/config"
	v1 "github.com/anotherhope/FizzBuzz/project/internal/api/v1"
	"github.com/anotherhope/FizzBuzz/project/internal/exit"
	"github.com/anotherhope/FizzBuzz/project/pkg/server"
	"github.com/spf13/cobra"
)

func main() {
	config.Helper.Run = func(cmd *cobra.Command, args []string) {
		srv := server.Create()
		srv.API(v1.STATUS)
		srv.API(v1.FIZZBUZZ)
		srv.Start()
	}

	if err := config.Helper.Execute(); err != nil {
		exit.WithCriticalError(err)
	}
}
