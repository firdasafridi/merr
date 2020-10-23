package main

import (
	"errors"
	"log"

	"github.com/firdasafridi/merr"
)

func main() {
	var mulerr merr.Error

	if err := sampleLogic1(); err != nil {
		mulerr.Set(err)
	}

	if err := sampleLogic2(); err != nil {
		mulerr.SetPrefix("Sample Prefix", err)
	}

	if err := mulerr.IsError(); err != nil {
		log.Println("[Count]", mulerr.Len())
		log.Println("[ERR]", mulerr.Error())
		return
	}
	log.Println("It will not executed")
}

func sampleLogic1() error {
	return errors.New("Failed to insert data")
}

func sampleLogic2() error {
	return errors.New("Failed do authorization")
}
