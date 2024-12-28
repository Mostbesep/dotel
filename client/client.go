package client

import (
	"bytes"
	"context"
	"github.com/tidwall/resp"
	"net"
)

type Client struct {
	addr string
	conn net.Conn
}

func New(url string) (*Client, error) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (c *Client) Set(ctx context.Context, key string, value string) error {

	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	err := wr.WriteArray(
		[]resp.Value{resp.StringValue("SET"), resp.StringValue(key), resp.StringValue(value)},
	)
	if err != nil {
		return err
	}
	_, err = c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
