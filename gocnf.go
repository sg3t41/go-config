package gocnf

import (
	"errors"
	"reflect"

	"github.com/sg3t41/gocnf/config"
	"github.com/sg3t41/gocnf/strategy"
	"github.com/sg3t41/gocnf/strategy/json"
	"github.com/sg3t41/gocnf/strategy/yaml"
)

// ErrTypeIsPointer indicates that a pointer type was passed as a type argument
// where a struct type was expected.
var ErrTypeIsPointer = errors.New("gocnf: type argument must be a struct, not a pointer")

func init() {
	strategy.Register(".json", &json.JSONStrategy{})
	strategy.Register(".yaml", &yaml.YamlStrategy{})
	strategy.Register(".yml", &yaml.YamlStrategy{})
}

func Unmarshal[T any](filePath string) (*T, error) {
	if isTypePtr[T]() {
		return nil, ErrTypeIsPointer
	}

	s, err := strategy.Get(filePath)
	if err != nil {
		return nil, err
	}

	c := config.NewConfig()
	c.
		SetFilePath(filePath).
		SetStrategy(s)

	var out T
	if err := c.Unmarshal(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

// isTypePtr checks if the generic type T is a pointer.
func isTypePtr[T any]() bool {
	// reflect.TypeOf((*T)(nil)).Elem() gets the reflection type of T.
	// We then check if its Kind is a pointer.
	return reflect.TypeOf((*T)(nil)).Elem().Kind() == reflect.Ptr
}
