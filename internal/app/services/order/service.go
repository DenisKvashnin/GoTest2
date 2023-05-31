package order

import (
	"TestTask/internal/app/providers"
	"TestTask/internal/app/providers/order"
	orderProvider "TestTask/internal/app/providers/order"
	_ "TestTask/internal/app/repositories/order"
	orderRepo "TestTask/internal/app/repositories/order"
	"log"
)

type Service struct {
	repo     *orderRepo.Repository
	provider *order.Provider
}

func New(repository *orderRepo.Repository, provider *orderProvider.Provider) *Service {
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
