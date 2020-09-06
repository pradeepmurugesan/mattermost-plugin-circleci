package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	subscribeTrigger  = "subscription"
	subscribeHint     = "<" + subscribeListTrigger + "|" + subscribeChannelTrigger + ">"
	subscribeHelpText = "Manage your subscriptions"

	// TODO:  add subcommands Triggers, Hints and HelpTexts here
	subscribeListTrigger  = "list"
	subscribeListHint     = ""
	subscribeListHelpText = "List the CircleCI subscriptions for the current channel"

	subscribeChannelTrigger  = "channel"
	subscribeChannelHint     = "<username> <repository> [--flags]"
	subscribeChannelHelpText = "Subscribe the current channel to CircleCI notifications for a repository"

	// TODO: add theses subCommands in getAutocompleteData()
)

func (p *Plugin) executeSubscribe(context *model.CommandArgs, circleciToken string, split []string) (*model.CommandResponse, *model.AppError) {
	subcommand := commandHelpTrigger
	if len(split) > 0 {
		subcommand = split[0]
	}

	switch subcommand {
	case commandHelpTrigger:
		return p.sendHelpResponse(context, subscribeTrigger)

	case subscribeListTrigger:
		return executeSubscribeList(p, context)

	case subscribeChannelTrigger:
		return executeSubscribeChannel(p, context, split[1:])

	default:
		return p.sendIncorrectSubcommandResponse(context, subscribeTrigger)
	}
}

func executeSubscribeList(p *Plugin, context *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// TODO
	return nil, nil
}

func executeSubscribeChannel(p *Plugin, context *model.CommandArgs, args []string) (*model.CommandResponse, *model.AppError) {
	// TODO

	// if err := p.AddSubscription(fullNameFromOwnerAndRepo(owner, repo), sub); err != nil {
	// 	return errors.Wrap(err, "could not add subscription")
	// }

	return p.sendEphemeralResponse(context, "Successfully subscribed to TODO tell channel name"), nil
}
