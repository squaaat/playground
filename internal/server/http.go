package server

import (
	"net"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	HTTP *fiber.App
	Host string
	Port string

}


func New(host, port string) *Server {

	a := fiber.New()
	a.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})


	return &Server{
		HTTP: a,
		Host: host,
		Port: port,
	}
}





func (s *Server) Listen () error {
	return s.HTTP.Listen(net.JoinHostPort(s.Host, s.Port))
}