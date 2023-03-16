// Package config provides functions/var/const for loading and accessing configuration settings for the application.
package config

import (
	"github.com/spf13/cobra"
)

// Use Cobra package to create a CLI and give Args gesture
var Helper *cobra.Command = &cobra.Command{
	Use:                   "fizzbuzz",
	Short:                 "FizzBuzz API Server",
	DisableAutoGenTag:     true,
	DisableFlagsInUseLine: true,
}

var (
	ENABLE_HTTPS  = false
	TLS_PATH      = "/etc/ssl/certs"
	TLS_CERT_FILE = "/server.crt"
	TLS_KEY_FILE  = "/server.key"
)

// Initialize Helper
func init() {
	Helper.PersistentFlags().BoolVarP(&ENABLE_HTTPS, "enable-https", "S", ENABLE_HTTPS, "Enable HTTPS")
	Helper.PersistentFlags().StringVarP(&TLS_PATH, "tls", "t", TLS_PATH, "define certificats TLS path")
	Helper.PersistentFlags().StringVarP(&TLS_CERT_FILE, "cert", "c", TLS_CERT_FILE, "define certificats CERT name")
	Helper.PersistentFlags().StringVarP(&TLS_KEY_FILE, "key", "k", TLS_KEY_FILE, "define certificats KEY name")
}
