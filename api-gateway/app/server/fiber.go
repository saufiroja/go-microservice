package server

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
	*fiber.App
}

func NewFiberServer() *FiberServer {
	return &FiberServer{
		fiber.New(),
	}
}

func (s *FiberServer) Start(port string) error {
	return s.Listen(port)
}
