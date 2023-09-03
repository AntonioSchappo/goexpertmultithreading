package entity

import "fmt"

const (
	Viacep            string = "viacep"
	Apicep            string = "apicep"
	ResponseFormatCli string = "Cep: %s;\nAddress: %s;\nDistrict: %s;\nCity: %s;\nState: %s"
	ReponseFormatJson string = "Cep: %s; Address: %s; District: %s; City: %s; State: %s"
	Timeout           string = "timeout"
	TimeoutMessage    string = "the calls to ViaCep and ApiCep for the Cep code %s timed out"
)

type ApiCepOutput struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCepOutput struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func FormatApiCepString(format string, a ApiCepOutput) string {
	return fmt.Sprintf(format, a.Cep, a.Street, a.Neighborhood, a.City, a.State)
}

func FormatViaCepString(format string, v ViaCepOutput) string {
	return fmt.Sprintf(format, v.Cep, v.Logradouro, v.Bairro, v.Localidade, v.Uf)
}

func FormatTimeoutString(c string) string {
	return fmt.Sprintf(TimeoutMessage, c)
}
