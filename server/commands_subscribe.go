package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	subscribeTrigger  = "subscribe"
	subscribeHint     = "TODO"
	subscribeHelpText = "TODO"

	// TODO:  add subcommands Triggers, Hints and HelpTexts here
	subscribeListTrigger  = "list"
	subscribeListHint     = "TODO"
	subscribeListHelpText = "TODO"

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

	default:
		return p.sendIncorrectSubcommandResponse(context, subscribeTrigger)
	}
}

// TODO: implements the subcommands
func executeSubscribeList(p *Plugin, context *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// TODO
	return nil, nil
}
