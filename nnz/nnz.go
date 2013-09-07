// Package nnz defines variants of primitive types where the zero value
// represents null when (de)serializing with encoding/json and database/sql.
package nnz

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int int

// Scan implements the database/sql/driver.Scanner interface.
func (i *Int) Scan(v interface{}) error {
	if v == nil {
		*i = 0
		return nil
	}
	switch v := v.(type) {
	case int64:
		*i = Int(v)
	default:
		return fmt.Errorf("nnz: scanning %T, got %T", i, v)
	}
	return nil
}

// Value implements the database/sql/driver.Valuer interface.
func (i Int) Value() (driver.Value, error) {
	if i == 0 {
		return nil, nil
	}
	return int64(i), nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (i Int) MarshalJSON() ([]byte, error) {
	if i == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(int(i))
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v == nil {
		*i = 0
	} else if v, ok := v.(float64); ok {
		*i = Int(v)
	} else {
		return fmt.Errorf("nnz: unmarshaling %T, got %T", i, v)
	}
	return nil
}

type String string

// Scan implements the database/sql/driver.Scanner interface.
func (s *String) Scan(v interface{}) error {
	if v == nil {
		*s = ""
		return nil
	}
	switch v := v.(type) {
	case []byte:
		*s = String(v)
	case string:
		*s = String(v)
	default:
		return fmt.Errorf("nnz: scanning %T, got %T", s, v)
	}
	return nil
}

// Value implements the database/sql/driver.Valuer interface.
func (s String) Value() (driver.Value, error) {
	if s == "" {
		return nil, nil
	}
	return string(s), nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if s == "" {
		return json.Marshal(nil)
	}
	return json.Marshal(string(s))
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v == nil {
		*s = ""
	} else if v, ok := v.(string); ok {
		*s = String(v)
	} else {
		return fmt.Errorf("nnz: unmarshaling %T, got %T", s, v)
	}
	return nil
}
