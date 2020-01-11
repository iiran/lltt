package db

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/iiran/lltt/pkg/core/errors"
)

type JsonMap map[string]interface{}

func NewJsonMapFromStringInt64(from map[string]int64) JsonMap {
	b := make(JsonMap)
	for k, v := range from {
		b[k] = v
	}
	return b
}

func NewJsonMapFromStringFloat64(from map[string]float64) JsonMap {
	b := make(JsonMap)
	for k, v := range from {
		b[k] = v
	}
	return b
}

func NewJsonMapFromStrStringString(from map[string]string) JsonMap {
	b := make(JsonMap)
	for k, v := range from {
		b[k] = v
	}
	return b
}

func (p JsonMap) MapStringInt64() (res map[string]int64, err error) {
	res = make(map[string]int64, 0)
	for k, v := range p {
		i, ok := v.(float64)
		if !ok {
			return nil, errors.GetErr(errors.TYPE_ASSERT_ERR)
		}
		res[k] = int64(i)
	}
	return res, nil
}

func (p JsonMap) MapStringString() (res map[string]string, err error) {
	res = make(map[string]string, 0)
	for k, v := range p {
		i, ok := v.(string)
		if !ok {
			return nil, errors.GetErr(errors.TYPE_ASSERT_ERR)
		}
		res[k] = i
	}
	return res, nil
}

// Types implementing Valuer interface are able to convert
// themselves to a driver Value.
func (p JsonMap) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan assigns a value from a database driver.
func (p *JsonMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.GetErr(errors.TYPE_ASSERT_ERR)
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.GetErr(errors.TYPE_ASSERT_ERR)
	}

	return nil
}
