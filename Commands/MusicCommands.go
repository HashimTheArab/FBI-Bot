package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
)

func PlayCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	/*conn, err := session.ChannelVoiceJoin(ctx.GetGuild().ID, "TODO", false, false)
	if err != nil {
		SendError(ctx, session, "Failed to connect! Error: " + err.Error())
		return nil
	}*/
	return nil
}
