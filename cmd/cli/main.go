package main

import (
	"fmt"
	"os"

	"github.com/AntonioSchappo/desafiomultithreading/pkg/client"
	"github.com/AntonioSchappo/desafiomultithreading/pkg/entity"
	"github.com/AntonioSchappo/desafiomultithreading/pkg/requester"
)

func main() {
	for _, cep := range os.Args[1:] {
		apiClient := client.NewClient()
		requester := requester.NewRequester(apiClient)
		msg, err := requester.GetCep(cep)
		if err != nil {
			fmt.Println("There was an error retrieving the data for the cep informed. Please check if it was typed correctly and try again.")
		} else {
			fmt.Print(entity.FormatCliMessage(msg, cep))
		}
	}
}
