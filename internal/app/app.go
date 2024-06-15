package app

import (
	"fmt"
	"log"
	"os"

	"github.com/st/internal/handler/ea"
	"github.com/st/internal/handler/har"
	"github.com/st/internal/handler/hosts"
	"github.com/st/internal/handler/info"
	"github.com/st/internal/handler/ptr"
	"github.com/st/internal/handler/rps"
	"github.com/st/internal/handler/target"
)

const (
	INFO   = "info"
	PTR    = "ptr"
	RPS    = "rps"
	TARGET = "target"
	HOSTS  = "hosts"
	T443   = "443"
	T80    = "80"
	HAR    = "har"
	EA     = "ea"
)

type Handler interface {
	Handle() (string, error)
}

func Run() {
	// Handlers
	handlers := map[string]Handler{
		INFO:   info.NewHandler(),
		PTR:    ptr.NewHandler(),
		RPS:    rps.NewHandler(),
		TARGET: target.NewHandler(),
		HOSTS:  hosts.NewHandler(),
		HAR:    har.NewHandler(),
		EA:     ea.NewHandler(),
	}

	// Get action
	action := getAction()

	// Get handler depending on action
	var h Handler
	var ok bool
	if h, ok = handlers[action]; !ok {
		log.Fatalf("unknown action [%s]", action)
	}

	// Get result and print result
	if text, err := h.Handle(); err == nil {
		fmt.Println(text)
	} else {
		log.Println(err)
	}
}

func getAction() (action string) {
	if len(os.Args) == 1 { // only binary
		action = INFO
	} else { // the action exists
		value := os.Args[1]
		if value == T443 || value == T80 {
			action = TARGET
			return
		}
		action = value
	}
	return
}
