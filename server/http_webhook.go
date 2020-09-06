package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

type BuildInfos struct {
	CircleProjectReponame string `json:"CircleProjectReponame"`
	CircleBuildNum        int    `json:"CircleBuildNum"`
}

// Convert the build info into a post attachment
func (bi *BuildInfos) toPostAttachments() []*model.SlackAttachment {
	const (
		circleStatusFailed = "failed"
	)

	attachment := &model.SlackAttachment{
		AuthorName: "CircleCI Integration",
		AuthorLink: "TODO URL TO BUILD FAILED",
		Fields: []*model.SlackAttachmentField{
			{
				Title: "Repo",
				Short: true,
				Value: bi.CircleProjectReponame,
			},
			{
				Title: "Build number",
				Short: true,
				Value: fmt.Sprintf("%d", bi.CircleBuildNum),
			},
		},
	}

	buildStatus := "success" // TODO handle this attribute correctly

	if buildStatus == circleStatusFailed {
		attachment.AuthorIcon = buildFailedURL
		attachment.Title = "TODO BUild failed"

		// TODO attachment.Color = red
	} else {
		attachment.AuthorIcon = buildGreenURL
		attachment.Title = "TODO BUild passed"
		// TODO attachment.Color = green
	}

	attachment.Fallback = attachment.Title

	return []*model.SlackAttachment{
		attachment,
	}
}

func httpHandleWebhook(p *Plugin, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		p.respondAndLogErr(w, http.StatusMethodNotAllowed, errors.New("method"+r.Method+"is not allowed, must be POST"))
		return
	}

	buildInfos := new(BuildInfos)
	if err := json.NewDecoder(r.Body).Decode(&buildInfos); err != nil {
		p.API.LogError("Unable to decode JSON for received webkook.", "Error", err.Error())
		return
	}

	channelToPost := "jyswqqfas3bk7db3eg3m1cy9xh" // TODO get the good channel

	post := &model.Post{
		ChannelId: channelToPost,
		UserId:    p.botUserID,
		Message:   "Webhook received",
	}
	post.AddProp("attachments", buildInfos.toPostAttachments())

	_, appErr := p.API.CreatePost(post)
	if appErr != nil {
		p.API.LogError("Failed to create Post", "appError", appErr)
	}
}
