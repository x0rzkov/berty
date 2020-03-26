package jsonclient

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"

	"berty.tech/core/api/client"
)

type unaryCallback func(*client.Client, context.Context, []byte) (interface{}, metadata.MD, metadata.MD, error)

var unaryMap map[string]unaryCallback

func registerUnary(name string, endpoint unaryCallback) {
	if unaryMap == nil {
		unaryMap = make(map[string]unaryCallback)
	}
	unaryMap[name] = endpoint
}

func CallUnary(ctx context.Context, c *client.Client, endpoint string, jsonInput []byte) (interface{}, metadata.MD, metadata.MD, error) {
	if jsonInput == nil {
		jsonInput = []byte("{}")
	}

	for name, handler := range unaryMap {
		if strings.ToLower(name) == strings.ToLower(endpoint) {
			return handler(c, ctx, jsonInput)
		}
	}
	return nil, nil, nil, fmt.Errorf("unknown endpoint: %q", endpoint)
}

func Unaries() []string {
	names := []string{}
	for name := range unaryMap {
		names = append(names, name)
	}
	return names
}