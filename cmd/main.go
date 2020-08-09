package main

import (
	"book-survey/internal/io"
	"book-survey/internal/services"
	"errors"
	"fmt"
	"log"
)

func run() error {
	reqParams := io.SetReqParams()

	resBody, err := services.SendRequest(reqParams)
	if err != nil {
		return err
	}

	decResBody, err := services.DecordResponse(resBody)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", decResBody)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(errors.Unwrap(err))
	}
}
