package commands

import (
	"encoding/json"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"github.com/prim69/fbi-bot/utils/settings"
	"github.com/prim69/wolframgo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AskCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if settings.Data.WolframAppID == "" {
		return SendError(ctx, session, "No WolframAlpha AppID is specified in the settings file. Please contact the server owner.")
	}
	if len(ctx.GetArgs()) < 1 {
		return GetCommand("ask").SendUsage(ctx, session)
	}
	answer, err := wolframgo.GetSimpleResult(strings.Join(ctx.GetArgs(), " "))
	if err != nil {
		return SendError(ctx, session, err.Error())
	}
	_, err = session.ChannelMessageSend(ctx.GetChannel().ID, answer)
	return err
}

func LookupCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if settings.Data.XblApiToken == "" {
		return SendError(ctx, session, "There is no XblApiToken is specified in the settings file. Sign up to get one at https://xbl.io/.")
	}
	if len(ctx.GetArgs()) < 1 {
		return GetCommand("lookup").SendUsage(ctx, session)
	}
	start := time.Now()
	request, err := http.NewRequest("GET", "https://xbl.io/api/v2/friends/search?gt="+ctx.GetArgs()[0], nil)
	if err != nil {
		return err
	}
	request.Header.Add("X-Authorization", settings.Data.XblApiToken)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	rsp := utils.Response{}
	if err := json.Unmarshal(data, &rsp); err != nil {
		return err
	}
	if len(rsp.ProfileUsers) < 1 {
		_, err := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
			Title:       "User Not Found!",
			Description: "A user with that gamertag was not found.",
			Color:       utils.Green,
		})
		return err
	}
	t := time.Now().Sub(start).Seconds()
	_, err = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Title: "Xbox Account Info",
		Color: 0x33FF33,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Ran by " + ctx.GetAuthor().String() + " | Time: " + strconv.FormatFloat(t, 'f', 3, 64) + "s",
			IconURL: ctx.GetAuthor().AvatarURL(""),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: rsp.ProfileUsers[0].Settings[0].Value},
		Fields: []*discordgo.MessageEmbedField{{Name: "Gamertag", Value: rsp.ProfileUsers[0].Settings[2].Value},
			{Name: "XUID", Value: rsp.ProfileUsers[0].ID},
			{Name: "Account Tier", Value: rsp.ProfileUsers[0].Settings[3].Value},
			{Name: "Gamer Score", Value: rsp.ProfileUsers[0].Settings[1].Value},
			{Name: "Xbox One Rep", Value: rsp.ProfileUsers[0].Settings[4].Value},
			{Name: "Real Name", Value: rsp.Name()},
			{Name: "Location", Value: rsp.Location()},
			{Name: "Bio", Value: rsp.Bio()}},
	})
	return err
}
