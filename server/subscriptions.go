package main

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/pkg/errors"
)

const (
	subscriptionsKVKey = "subscriptions"
)

type Subscriptions struct {
	Repositories map[string][]*Subscription
}

func (p *Plugin) GetSubscriptionsByChannel(channelID string) ([]*Subscription, error) {
	var filteredSubs []*Subscription
	subs, err := p.GetSubscriptions()
	if err != nil {
		return nil, errors.Wrap(err, "could not get subscriptions")
	}

	for repo, v := range subs.Repositories {
		for _, s := range v {
			if s.ChannelID == channelID {
				// this is needed to be backwards compatible
				if len(s.Repository) == 0 {
					s.Repository = repo
				}
				filteredSubs = append(filteredSubs, s)
			}
		}
	}

	sort.Slice(filteredSubs, func(i, j int) bool {
		return filteredSubs[i].Repository < filteredSubs[j].Repository
	})

	return filteredSubs, nil
}

func (p *Plugin) GetSubscriptions() (*Subscriptions, error) {
	var subscriptions *Subscriptions

	value, appErr := p.API.KVGet(subscriptionsKVKey)
	if appErr != nil {
		return nil, errors.Wrap(appErr, "could not get subscriptions from KVStore")
	}

	if value == nil {
		return &Subscriptions{Repositories: map[string][]*Subscription{}}, nil
	}

	err := json.NewDecoder(bytes.NewReader(value)).Decode(&subscriptions)
	if err != nil {
		return nil, errors.Wrap(err, "could not properly decode subscriptions key")
	}

	return subscriptions, nil
}

func (p *Plugin) AddSubscription(repo string, sub *Subscription) error {
	subs, err := p.GetSubscriptions()
	if err != nil {
		return errors.Wrap(err, "could not get subscriptions")
	}

	repoSubs := subs.Repositories[repo]
	if repoSubs == nil {
		repoSubs = []*Subscription{sub}
	} else {
		exists := false
		for index, s := range repoSubs {
			if s.ChannelID == sub.ChannelID {
				repoSubs[index] = sub
				exists = true
				break
			}
		}

		if !exists {
			repoSubs = append(repoSubs, sub)
		}
	}

	subs.Repositories[repo] = repoSubs

	err = p.StoreSubscriptions(subs)
	if err != nil {
		return errors.Wrap(err, "could not store subscriptions")
	}

	return nil
}

func (p *Plugin) RemoveSubscription(channelID string, repo string) error {
	// owner, repo := parseOwnerAndRepo(repo, p.getBaseURL())
	// if owner == "" && repo == "" {
	// 	return errors.New("invalid repository")
	// }
	// repoWithOwner := fmt.Sprintf("%s/%s", owner, repo)

	// subs, err := p.GetSubscriptions()
	// if err != nil {
	// 	return errors.Wrap(err, "could not get subscriptions")
	// }

	// repoSubs := subs.Repositories[repoWithOwner]
	// if repoSubs == nil {
	// 	return nil
	// }

	// removed := false
	// for index, sub := range repoSubs {
	// 	if sub.ChannelID == channelID {
	// 		repoSubs = append(repoSubs[:index], repoSubs[index+1:]...)
	// 		removed = true
	// 		break
	// 	}
	// }

	// if removed {
	// 	subs.Repositories[repoWithOwner] = repoSubs
	// 	if err := p.StoreSubscriptions(subs); err != nil {
	// 		return errors.Wrap(err, "could not store subscriptions")
	// 	}
	// }

	return nil
}

func (p *Plugin) StoreSubscriptions(s *Subscriptions) error {
	b, err := json.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "error while converting subscriptions map to json")
	}

	if appErr := p.API.KVSet(subscriptionsKVKey, b); appErr != nil {
		return errors.Wrap(appErr, "could not store subscriptions in KV store")
	}

	return nil
}

func (p *Plugin) GetSubscribedChannelsForRepository(owner, repository string) []*Subscription {
	// name := repo.GetFullName()
	// org := strings.Split(name, "/")[0]
	// subs, err := p.GetSubscriptions()
	// if err != nil {
	// 	return nil
	// }

	// // Add subscriptions for the specific repo
	// subsForRepo := []*Subscription{}
	// if subs.Repositories[name] != nil {
	// 	subsForRepo = append(subsForRepo, subs.Repositories[name]...)
	// }

	// // Add subscriptions for the organization
	// orgKey := fullNameFromOwnerAndRepo(org, "")
	// if subs.Repositories[orgKey] != nil {
	// 	subsForRepo = append(subsForRepo, subs.Repositories[orgKey]...)
	// }

	// if len(subsForRepo) == 0 {
	// 	return nil
	// }

	subsToReturn := []*Subscription{}

	// for _, sub := range subsForRepo {
	// 	if repo.GetPrivate() && !p.permissionToRepo(sub.CreatorID, name) {
	// 		continue
	// 	}
	// 	subsToReturn = append(subsToReturn, sub)
	// }

	return subsToReturn
}
