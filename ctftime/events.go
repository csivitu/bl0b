package ctftime

import (
	"fmt"

	"github.com/csivitu/bl0b/utils"
)

// GetEvents returns events between startTime and endTime. The number
// of events is defined by the value of limit.
func (ctf *CTFtime) GetEvents(limit int, startTime int64, endTime int64) (Events, error) {
	endpoint := fmt.Sprintf(
		"/events/?limit=%d&start=%d&finish=%d",
		limit,
		startTime,
		endTime,
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

	return events, nil
}
