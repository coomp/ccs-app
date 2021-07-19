package handler

import (
	"fmt"

	ccssdk "github.com/coomp/ccs-sdk"
)

func NewExampleHandler() ccssdk.HandleFunc {
	return func(c ccssdk.MessageContext) error {
		fmt.Printf("ExampleHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}
