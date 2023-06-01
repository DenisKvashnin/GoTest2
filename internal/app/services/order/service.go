package order

import (
	"TestTask/internal/app/providers"
	OrderProvider "TestTask/internal/app/providers/order"
	OrderRepository "TestTask/internal/app/repositories/order"
	"log"
)

type Service struct {
	repo     *OrderRepository.Repository
	provider *OrderProvider.Provider
}

func New(repository *OrderRepository.Repository, provider *OrderProvider.Provider) *Service {
	return &Service{
		repo:     repository,
		provider: provider,
	}
}

func (o Service) GetAndSaveOrder() (providers.Response, error) {
	response, err := o.provider.GetOrder()
	if err != nil {
		return providers.Response{}, err
	}

	log.Println("Получено: ", response)

	o.repo.SaveAll(response.Content)

	return response, err
}
