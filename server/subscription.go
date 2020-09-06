package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	flagOnlyFailedBuilds = "only-failed"
)

type SubscriptionFlags struct {
	onlyFailedBuilds bool
}

func (s *SubscriptionFlags) AddFlag(flag string) {
	switch flag { // nolint:gocritic // It's expected that more flags get added.
	case flagOnlyFailedBuilds:
		s.onlyFailedBuilds = true
	}
}

func (s SubscriptionFlags) String() string {
	flags := []string{}

	if s.onlyFailedBuilds {
		flag := "--" + flagOnlyFailedBuilds
		flags = append(flags, flag)
	}

	return strings.Join(flags, ",")
}

type Subscription struct {
	ChannelID  string
	CreatorID  string
	Flags      SubscriptionFlags
	Repository string
}

func (s *Subscription) ToSlackAttachmentField() *model.SlackAttachmentField {
	return &model.SlackAttachmentField{
		Title: s.Repository,
		Short: true,
		Value: fmt.Sprintf(
			"Subscribed by: %s\nFlags: %s",
			s.CreatorID, // TODO id to username
			s.Flags.String(),
		),
	}
}
