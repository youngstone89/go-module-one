package kafka

import "fmt"

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Init() {
	fmt.Println("Initializing kafka client...")
}
