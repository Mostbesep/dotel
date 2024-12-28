package main

import (
	"bytes"
	"fmt"
	commands "github.com/Mostbesep/dotet/commands"
	"github.com/tidwall/resp"
	"io"
	"log"
	"strings"
)

const (
	CommandSET  = "set"
	CommandGET  = "get"
	CommandMGET = "mget"
)

func ParseCommand(msg string) ([]commands.Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(msg))
	commandList := []commands.Command{}
	for {

		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var cmd commands.Command

		if v.Type() == resp.Array {
			cmd, err = handleArray(v)
			if err != nil {
				return nil, err
			}
		}

		commandList = append(commandList, cmd)

	}
	return commandList, nil
}

func handleArray(v resp.Value) (commands.Command, error) {
	values := v.Array()
	var cmd commands.Command

	switch strings.ToLower(values[0].String()) {
	case CommandSET:
		if len(values) < 3 {
			return nil, fmt.Errorf("invalid SET command, , where %s", v.String())
		}
		cmd = commands.CommandSET{Key: values[1].String(), Value: values[2].String()}
	case CommandGET:
		if len(values) < 2 {
			return nil, fmt.Errorf("invalid GET command, where %s", v.String())
		}
		cmd = commands.CommandGET{Key: values[1].String()}
	case CommandMGET:
		keys := []string{}
		for _, key := range values[1:] {
			keys = append(keys, key.String())
		}
		cmd = commands.CommandMGET{
			Keys: keys,
		}
	default:
		return nil, fmt.Errorf("invalid command, or unknow command, where %s", v.String())
	}
	return cmd, nil

}
