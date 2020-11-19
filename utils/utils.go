package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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

// SetInterval runs a function repeatedly in a goroutine
// and returns a channel `done` to stop the routine
func SetInterval(f func(time.Time), t time.Duration) chan bool {
	ticker := time.NewTicker(t)

	done := make(chan bool, 1)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				f(t)
			case <-done:
				return
			}
		}
	}()
	return done
}
