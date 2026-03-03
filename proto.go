package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

// we need to implement redis serializtion protocol which clients use to communicate with the redis server.

type Command interface {
}

type SetCommand struct {
	key, val string
}

func parseCommand(raw string) (Command, error) {

	rd := resp.NewReader(bytes.NewBufferString(raw))

	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read %s\n", v.Type())
		if v.Type() == resp.Array {
			for i, v := range v.Array() {
				fmt.Printf("  #%d %s, value: '%s'\n", i, v.Type(), v)
			}
		}
	}

	return "foo", nil
}
