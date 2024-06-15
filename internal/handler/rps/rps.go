package rps

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const ARGS = 5 // |0| binary |1| rps |2| requests |3| since |4| until
var err error

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (string, error) {
	if len(os.Args) != ARGS {
		return "", fmt.Errorf("args number error [expected: %d]", ARGS)
	}

	var requests float64
	if requests, err = strconv.ParseFloat(os.Args[2], 64); err != nil { // 12345
		return "", fmt.Errorf("parse requests error > %s", err)
	}

	var since time.Time
	if since, err = time.Parse(time.DateTime, os.Args[3]); err != nil { // "2023-03-30 11:00:00"
		return "", fmt.Errorf("parse time since error > %s", err)
	}

	var until time.Time
	if until, err = time.Parse(time.DateTime, os.Args[4]); err != nil { // "2023-03-30 11:07:00"
		return "", fmt.Errorf("parse time until error > %s", err)
	}

	if since.Unix() >= until.Unix() {
		return "", fmt.Errorf("until must be older than since")
	}

	seconds := until.Unix() - since.Unix()
	
	return fmt.Sprintf("%.2f rps", requests/float64(seconds)), nil
}
