package handler

import "amartha-loan-svc/usecase"

type AmarthaHttpServer struct {
	app usecase.AmarthaUsecaseInterface
}

func NewAmarthaHttpServer(app usecase.AmarthaUsecaseInterface) AmarthaHttpServer {
	return AmarthaHttpServer{app: app}
}
