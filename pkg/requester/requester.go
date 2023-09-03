package requester

import (
	"time"

	"github.com/AntonioSchappo/desafiomultithreading/pkg/entity"
)

type ApiClient interface {
	GetApiCep(cep string) (entity.ApiCepOutput, error)
	GetViaCep(cep string) (entity.ViaCepOutput, error)
}

type Requester struct {
	client ApiClient
}

func NewRequester(c ApiClient) *Requester {
	return &Requester{client: c}
}

func (r *Requester) GetCep(cep string) (entity.Message, error) {
	c1 := make(chan entity.ApiCepOutput)
	c2 := make(chan entity.ViaCepOutput)
	c3 := make(chan error)
	c4 := make(chan error)

	// ApiCep
	go func(c string) {
		for {
			response, err := r.client.GetApiCep(cep)
			if err != nil {
				c3 <- err
			} else {
				c1 <- response
			}
		}
	}(cep)

	//ViaCep
	go func(c string) {
		for {
			response, err := r.client.GetViaCep(cep)
			if err != nil {
				c4 <- err
			} else {
				c2 <- response
			}
		}
	}(cep)

	for {
		select {
		case output := <-c1:
			return entity.Message{WinnerAPI: entity.Apicep, CepResponse: entity.FormatApiCepString(entity.ReponseFormatJson, output)}, nil
		case output := <-c2:
			return entity.Message{WinnerAPI: entity.Viacep, CepResponse: entity.FormatViaCepString(entity.ReponseFormatJson, output)}, nil
		case err := <-c3:
			return entity.Message{}, err
		case err := <-c4:
			return entity.Message{}, err
		case <-time.After(time.Second * 1):
			return entity.Message{WinnerAPI: entity.Timeout, CepResponse: entity.FormatTimeoutString(cep)}, nil
		}
	}
}
