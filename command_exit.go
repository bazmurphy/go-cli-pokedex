package main

import "os"

func commandExit(cfg *config, parameters []string) error {
	os.Exit(0)
	return nil
}
