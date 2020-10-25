package main

type Client struct {
	name string
	sex  int8
}

func NewClient(name string) *Client {
	return &Client{
		name: name,
		sex:  0,
	}
}
