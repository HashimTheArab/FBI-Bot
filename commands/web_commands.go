package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils/settings"
	"github.com/prim69/wolframgo"
	"strings"
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
