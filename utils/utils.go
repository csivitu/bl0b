package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Status specifies the status of events in the database
type Status string

const (
	// Upcoming status
	Upcoming Status = "upcoming"
	// Ongoing status
	Ongoing Status = "ongoing"
	// Over staatus
	Over Status = "over"
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

// ComputeStatus finds if a CTF is upcoming, ongoing or over
func ComputeStatus(start time.Time, finish time.Time) Status {
	t := time.Now()

	if t.After(start) && t.Before(finish) {
		return Ongoing
	}

	if t.Before(start) && t.Before(finish) {
		return Upcoming
	}

	return Over
}