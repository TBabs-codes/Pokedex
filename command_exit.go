package main

import "os"

func callbackExit(nextPrev *config) error{
	os.Exit(0)
	return nil
}