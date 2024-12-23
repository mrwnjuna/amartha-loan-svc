package usecase

import "amartha-loan-svc/repository"

type AmarthaUsecase struct {
	amarthaRepo repository.AmarthaRepoInterface
}

type AmarthaUsecaseInterface interface {
	LoansUsecase
	MarketplaceUsecase
}

func NewAmarthaUsecase(amarthaRepo repository.AmarthaRepoInterface) AmarthaUsecaseInterface {
	return &AmarthaUsecase{
		amarthaRepo: amarthaRepo,
	}
}
