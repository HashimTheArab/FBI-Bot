package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
)

func QueryCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) > 0 {
		return utils.Query(ctx.GetArgs()[0], "19132", ctx, session, utils.LongQuery)
	}
	return GetCommand("query").SendUsage(ctx, session)
}
