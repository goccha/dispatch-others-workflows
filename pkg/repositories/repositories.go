package repositories

import (
	"context"
	"encoding/json"
	"github.com/goccha/dispatch-others-workflows/pkg/debug"
	"github.com/google/go-github/v37/github"
	"golang.org/x/xerrors"
	"strings"
)

type Request []Event

func (req Request) Dispatch(ctx context.Context, eventType string) error {
	for _, event := range req {
		if err := event.Repository.Dispatch(ctx, eventType, event.Payload); err != nil {
			return err
		}
	}
	return nil
}

func Parse(payloads map[string]interface{}) (req Request, err error) {
	req = make([]Event, 0, len(payloads))
	for k, v := range payloads {
		var payload Event
		if payload, err = newEvent(k, v); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		} else {
			req = append(req, payload)
		}
	}
	return
}

func newEvent(repositoryName string, v interface{}) (event Event, err error) {
	array := strings.Split(repositoryName, "/")
	if len(array) != 2 {
		return Event{}, xerrors.New("the format \"owner/repoName\" is required")
	} else {
		event = Event{
			Repository: Repository{
				Owner:      array[0],
				Repository: array[1],
			},
			Payload: v,
		}
	}
	return
}

type Event struct {
	Repository Repository
	Payload    interface{}
}

type Repository struct {
	Owner      string
	Repository string
}

func (rep *Repository) Dispatch(ctx context.Context, eventType string, payload interface{}) (err error) {
	var raw *json.RawMessage
	if payload != nil {
		var b json.RawMessage
		if b, err = json.Marshal(payload); err != nil {
			return xerrors.Errorf(": %w", err)
		}
		debug.Print("payload", string(b))
		raw = &b
	}
	var r *github.Repository
	if r, _, err = client.Repositories.Dispatch(ctx, rep.Owner, rep.Repository, github.DispatchRequestOptions{
		EventType:     eventType,
		ClientPayload: raw,
	}); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	debug.Print("repository", r.String())
	return nil
}
