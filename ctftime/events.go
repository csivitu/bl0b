package ctftime

import (
	"fmt"
	"time"

	"github.com/csivitu/bl0b/utils"
)

// GetEvents returns events between startTime and endTime. The number
// of events is defined by the value of limit.
func (ctf *CTFtime) GetEvents(limit int, startTime time.Time, finishTime time.Time) (Events, error) {
	start := startTime.Unix()
	finish := finishTime.Unix()

	endpoint := fmt.Sprintf(
		"/events/?limit=%d&start=%d&finish=%d",
		limit,
		start,
		finish,
	)

	res, err := ctf.Get(endpoint)

	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	var events Events

	if err := utils.HTTPResponseToStruct(res, &events); err != nil {
		return nil, err
	}

	// TODO: Temporary hack, later on add Organizers in separate table
	for i := range events {
		events[i].Organizer = events[i].Organizers[0].Name
	}

	return events, nil
}
