package amqprpc

import (
	"go-scaffold/internal/usecase"
	"go-scaffold/pkg/rabbitmq/rmq_rpc/server"
)

func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}
	return routes
}
