package flights

import (
	"fmt"
	"time"
)

type Flight struct {
	Code    string
	Origin  string
	DueTime time.Time
}

func (f Flight) String() string {
	return fmt.Sprintf("Flight %s from %s is expected at %s", f.Code, f.Origin, f.DueTime.Format("15:04"))
}
