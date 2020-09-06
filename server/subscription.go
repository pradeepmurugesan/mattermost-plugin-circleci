package main

import (
	"strings"
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
	Features   string
	Flags      SubscriptionFlags
	Repository string
}

// func (s *Subscription) Pulls() bool {
// 	return strings.Contains(s.Features, featurePulls)
// }

// func (s *Subscription) Issues() bool {
// 	return strings.Contains(s.Features, featureIssues)
// }

func (s *Subscription) Pushes() bool {
	return strings.Contains(s.Features, "pushes")
}

func (s *Subscription) Creates() bool {
	return strings.Contains(s.Features, "creates")
}

func (s *Subscription) Deletes() bool {
	return strings.Contains(s.Features, "deletes")
}

func (s *Subscription) IssueComments() bool {
	return strings.Contains(s.Features, "issue_comments")
}

func (s *Subscription) PullReviews() bool {
	return strings.Contains(s.Features, "pull_reviews")
}

func (s *Subscription) Label() string {
	if !strings.Contains(s.Features, "label:") {
		return ""
	}

	labelSplit := strings.Split(s.Features, "\"")
	if len(labelSplit) < 3 {
		return ""
	}

	return labelSplit[1]
}

func (s *Subscription) ExcludeOrgMembers() bool {
	return s.Flags.onlyFailedBuilds
}
