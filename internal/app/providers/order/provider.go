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
		log.Fatalf("Error while requesting order: %s", err)
		return providers.Response{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия, падаем")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения ответа: %s", err)
	}

	var response providers.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error while unmarshalling orders: %s", err)
		return response, err
	}
	return response, nil
}
