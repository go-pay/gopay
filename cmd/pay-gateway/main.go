package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-pay/gopay/internal/paygateway"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", os.Getenv("PAY_GATEWAY_CONFIG"), "config file path (json)")
	flag.Parse()

	if configPath == "" {
		log.Fatal("missing --config (or PAY_GATEWAY_CONFIG)")
	}
	cfg, err := paygateway.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	srv, err := paygateway.NewServer(cfg)
	if err != nil {
		log.Fatalf("init server: %v", err)
	}
	log.Printf("pay-gateway listening on %s", cfg.Server.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
