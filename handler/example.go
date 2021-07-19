package handler

import (
	"fmt"

	ccssdk "github.com/coomp/ccs-sdk"
)

func NewExampleRequestHandler() ccssdk.RequestHandleFunc {
	return func(c ccssdk.RequestMessageContext) error {
		fmt.Printf("NewExampleRequestHandler got message at %v\n", c.GetTimestamp())
		return nil
	}
}

func NewExampleResponseHandler() ccssdk.ResponseHandleFunc {
	return func(c ccssdk.ResponseMessageContext) error {
		fmt.Printf("NewExampleResponseHandler got message at %v\n", c.GetTimestamp())
		return nil
	}
}
