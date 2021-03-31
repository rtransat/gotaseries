package gotaseries

import (
	"bytes"
	"errors"
	"fmt"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"text"`
}
type Errors []Error

func (errs Errors) Err() error {
	if len(errs) == 0 {
		return nil
	}

	errS := bytes.NewBuffer(nil)
	for _, e := range errs {
		_, _ = fmt.Fprintf(errS, "Code: %v Message: %v\n", e.Code, e.Message)
	}
	return errors.New(errS.String())
}
