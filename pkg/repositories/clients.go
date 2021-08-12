package repositories

import (
	"context"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func Setup(ctx context.Context, token string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}
