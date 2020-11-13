package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HTTPResponseToStruct converts response from CTFtime into
// to a struct specified by the param v
func HTTPResponseToStruct(r *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	if valid, ok := v.(interface {
		OK() error
	}); ok {
		err = valid.OK()
		if err != nil {
			return err
		}
	}

	return nil
}
