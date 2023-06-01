package order

import (
	"TestTask/internal/app/providers"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Provider struct {
	url string
}

func New(url string) *Provider {
	return &Provider{url: url}
}

func (p Provider) GetOrder() (providers.Response, error) {
	resp, err := http.Get(p.url)
	if err != nil {
		log.Println("Error while requesting order:", err)
		return providers.Response{}, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("Ошибка закрытия:", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return providers.Response{}, err
	}

	var response providers.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error while unmarshalling orders:", err)
		return response, err
	}
	return response, nil
}
