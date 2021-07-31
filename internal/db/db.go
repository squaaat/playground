package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	DB *sql.DB
}

type Todo struct {
	ID   int64	`json:"id"`
	Text string	`json:"text"`
}

func New(host string, port string, username string, password string, schema string) (*Client, error) {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, schema))
	if err != nil {
		return nil, err
	}
	return &Client{
		DB: db,
	}, nil
}

func (c *Client) InitDb() error {
	_, err := c.DB.Exec(`
		CREATE TABLE
		IF NOT EXISTS mytable
		(id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			todo TEXT NOT NULL)
	`)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Create(todo Todo) error {
	// Create
	res, err := c.DB.Exec("INSERT INTO mytable (todo) VALUES (?)", todo.Text)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

func (c *Client) GetById(todo Todo) (*Todo, error) {
	res, err := c.DB.Query(`SELECT id, todo FROM mytable WHERE id = ? LIMIT 1`, todo.ID)
	if err != nil{
		return nil, err
	}


	var t *Todo
	for res.Next() {
		t := &Todo{}
		err = res.Scan(&t.ID, &t.Text)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (c *Client) GetAll() ([]Todo, error) {
	res, err := c.DB.Query(`
		SELECT * FROM mytable
	`)
	if err != nil {
		return nil, err
	}

	fmt.Println((res))
	var todos []Todo
	for res.Next() {
		todo := Todo{}
		err = res.Scan(&todo.ID, &todo.Text)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
