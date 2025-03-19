package flights

import (
	"fmt"
	"time"
)

type Flight struct {
	Code    string    `json:"code"`
	Origin  string    `json:"from"`
	DueTime time.Time `json:"scheduled_arrival"`
}

func (f Flight) String() string {
	return fmt.Sprintf("Flight %s from %s is expected at %s", f.Code, f.Origin, f.DueTime.Format("15:04"))
}
