package json

import (
	"github.com/bytedance/sonic"
)

func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(buf []byte, v interface{}) error {
	return sonic.Unmarshal(buf, v)
}
