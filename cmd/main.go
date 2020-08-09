package main

import (
	"errors"
	"job-survey/internal/io"
	"job-survey/internal/services"
	"log"
)

func run() error {
	reqParams := io.SetReqParams()

	resBody, err := services.SendRequest(reqParams)
	if err != nil {
		return err
	}

	fmtBody, err := services.FormatResponse(resBody)
	if err != nil {
		return err
	}

	println(fmtBody)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(errors.Unwrap(err))
	}
}
