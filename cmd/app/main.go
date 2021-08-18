package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/goccha/dispatch-others-workflows/pkg/debug"
	"github.com/goccha/dispatch-others-workflows/pkg/repositories"
	"os"
)

type Option struct {
	Token     string
	EventType string
	Payload   string
}

var opt = &Option{}

func init() {
	flag.StringVar(&opt.Token, "token", "", "github token")
	flag.StringVar(&opt.EventType, "event-type", "", "event type")
	flag.StringVar(&opt.Payload, "payload", "", "payload")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	repositories.Setup(ctx, opt.Token)
	payload := make(map[string]interface{})
	debug.Print("payload", opt.Payload)
	if err := json.Unmarshal([]byte(opt.Payload), &payload); err != nil {
		abort(err)
	}
	if req, err := repositories.Parse(payload); err != nil {
		abort(err)
	} else {
		debug.Print("event-type", opt.EventType)
		if err = req.Dispatch(ctx, opt.EventType); err != nil {
			abort(err)
		}
	}
}

func abort(err error) {
	fmt.Printf("%+v\n", err)
	os.Exit(1)
}
