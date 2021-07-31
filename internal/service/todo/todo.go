package todo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/squaaat/playground/internal/db"
	"github.com/squaaat/playground/internal/server"
)

type Service struct {
	DB *db.Client
}

func New(db *db.Client) *Service {
	return &Service{
		DB: db,
	}
}


type Body struct {
	Todo db.Todo `json:"todo"`
}
func (srv *Service) RouteForHTTP(s *server.Server) {
	s.HTTP.Put("/create", func(c *fiber.Ctx) error {

		body := new(Body)

		if err := c.BodyParser((body)); err != nil {
			return err
		}


		err := srv.DB.Create(body.Todo)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	s.HTTP.Get("/todo/:id", func(c *fiber.Ctx) error {
		params := c.Params("id")
		var todo db.Todo
		_id, err := strconv.ParseInt(params,10,64)
		todo.ID = _id
		if err !=nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		res, err := srv.DB.GetById(todo)
		if err != nil {
			fmt.Println(err)
		}
		if res == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		encjson, _ := json.Marshal(res)

		return c.SendString(string(encjson))
	})

	s.HTTP.Get("/todos", func(c *fiber.Ctx) error {
		res, err := srv.DB.GetAll()

		if err != nil {
			fmt.Println(err)
		}
		encjson, _ := json.Marshal(res)

		return c.SendString(string(encjson))
	})
}
