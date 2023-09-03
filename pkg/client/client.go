package client

import (
	"encoding/json"
	"net/http"

	"github.com/AntonioSchappo/desafiomultithreading/pkg/entity"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetApiCep(cep string) (entity.ApiCepOutput, error) {
	var output entity.ApiCepOutput
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return output, err
	}
	return output, err
}

func (c *Client) GetViaCep(cep string) (entity.ViaCepOutput, error) {
	var output entity.ViaCepOutput
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return output, err
	}
	return output, err
}
