package decoder

import (
	"encoding/json"
	"go-api/errors"
	"io"
)

// DecodeJSON reads JSON data from reader and decodes it
// into the value pointed to by v.
func DecodeJSON(r io.Reader, v interface{}) error  {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return errors.Error(err.Error())
	}
	return nil
}
