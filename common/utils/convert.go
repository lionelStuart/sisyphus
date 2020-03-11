package utils

import (
	"github.com/mitchellh/mapstructure"
)

func MapToStruct(from map[string]interface{}, to interface{}) error {
	return mapstructure.Decode(from, to)
}
