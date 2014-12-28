// Package nnz defines variants of primitive types where the zero value
// represents null when (de)serializing with encoding/json and database/sql.
package nnz

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Bool bool

func (b *Bool) Scan(v interface{}) error {
	if v == nil {
		*b = false
		return nil
	}
	switch v := v.(type) {
	case bool:
		*b = Bool(v)
	default:
		return fmt.Errorf("nnz: scanning %T, got %T", b, v)
	}
	return nil
}

func (b Bool) Value() (driver.Value, error) {
	if b == false {
		return nil, nil
	}
	return bool(b), nil
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b == false {
		return json.Marshal(nil)
	}
	return json.Marshal(bool(b))
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v == nil {
		*b = false
	} else if v, ok := v.(bool); ok {
		*b = Bool(v)
	} else {
		return fmt.Errorf("nnz: unmarshaling %T, got %T", b, v)
	}
	return nil
}

// Int is a wrapper around int where Go int(0) serializes to SQL/JSON null, and
// SQL/JSON null deserializes to Go int(0).
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

// Int64 is a wrapper around int64 where Go int64(0) serializes to SQL/JSON
// null, and SQL/JSON null deserializes to Go int64(0).
type Int64 int64

// Scan implements the database/sql/driver.Scanner interface.
func (i *Int64) Scan(v interface{}) error {
	if v == nil {
		*i = 0
		return nil
	}
	switch v := v.(type) {
	case int64:
		*i = Int64(v)
	default:
		return fmt.Errorf("nnz: scanning %T, got %T", i, v)
	}
	return nil
}

// Value implements the database/sql/driver.Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	if i == 0 {
		return nil, nil
	}
	return int64(i), nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (i Int64) MarshalJSON() ([]byte, error) {
	if i == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(int64(i))
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v == nil {
		*i = 0
	} else if v, ok := v.(float64); ok {
		*i = Int64(v)
	} else {
		return fmt.Errorf("nnz: unmarshaling %T, got %T", i, v)
	}
	return nil
}

// Float64 is a wrapper around float64 where Go float64(0) serializes to SQL/JSON
// null, and SQL/JSON null deserializes to Go float64(0).
type Float64 float64

// Scan implements the database/sql/driver.Scanner interface.
func (f *Float64) Scan(v interface{}) error {
	if v == nil {
		*f = 0
		return nil
	}
	switch v := v.(type) {
	case float64:
		*f = Float64(v)
	default:
		return fmt.Errorf("nnz: scanning %T, got %T", f, v)
	}
	return nil
}

// Value implements the database/sql/driver.Valuer interface.
func (f Float64) Value() (driver.Value, error) {
	if f == 0 {
		return nil, nil
	}
	return float64(f), nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	if f == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(float64(f))
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v == nil {
		*f = 0
	} else if v, ok := v.(float64); ok {
		*f = Float64(v)
	} else {
		return fmt.Errorf("nnz: unmarshaling %T, got %T", f, v)
	}
	return nil
}

// String is a wrapper around string where Go "" serializes to SQL/JSON null,
// and SQL/JSON null deserializes to Go "".
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

//  Time is a wrapper around time.Time where Go time.Time(0) serializes to SQL/JSON
// null, and SQL/JSON null deserializes to Go time.Time(0).
type Time time.Time

var zt time.Time // the zero time.

// Scan implements the database/sql/driver.Scanner interface.
func (t *Time) Scan(v interface{}) error {
	if v == nil {
		*t = Time(zt)
		return nil
	}
	switch v := v.(type) {
	case time.Time:
		*t = Time(v)
	default:
		return fmt.Errorf("nnz: scanning %T, got %T", t, v)
	}
	return nil
}

// Value implements the database/sql/driver.Valuer interface.
func (t Time) Value() (driver.Value, error) {
	tm := time.Time(t)

	if tm.IsZero() {
		return nil, nil
	}

	return tm, nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	tm := time.Time(t)

	if tm.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(time.Time(t))
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v == nil {
		*t = Time(zt)
	} else if v, ok := v.(time.Time); ok {
		*t = Time(v)
	} else {
		return fmt.Errorf("nnz: unmarshaling %T, got %T", t, v)
	}
	return nil
}

// UnmarshalText implements the encoding/TextUnmarshaler interface.
func (t *Time) UnmarshalText(data []byte) error {
	if string(data) == "" {
		*t = Time(zt)
		return nil
	}

	nt, err := time.Parse(time.RFC3339, string(data))

	if err != nil {
		return err
	}

	*t = Time(nt)

	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (t Time) GobEncode() ([]byte, error) {
	return time.Time(t).GobEncode()
}

// GobDecode implements the gob.GobDecoder interface.
func (t *Time) GobDecode(data []byte) error {
	var tm time.Time

	if err := tm.UnmarshalBinary(data); err != nil {
		return err
	}

	*t = Time(tm)

	return nil
}
