package entity

import "fmt"

const (
	FormatCli string = "The winner api is: %s,\nThe data for cep code %s is as follows:\n%s"
)

type Message struct {
	WinnerAPI   string `json:"winner_api"`
	CepResponse string `json:"cep_response"`
}

func FormatCliMessage(m Message, c string) string {
	return fmt.Sprintf(FormatCli, m.WinnerAPI, c, m.CepResponse)
}
