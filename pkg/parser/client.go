package parser

import "encoding/json"

type commonRes[T any] struct {
	Message         string      `json:"string"`
	Data            T           `json:"data"`
	ValidationError interface{} `json:"validation error"`
}

func ParseResp[T any](res []byte) commonRes[T] {
	cr := commonRes[T]{}
	json.Unmarshal(res, &cr)

	return cr
}
