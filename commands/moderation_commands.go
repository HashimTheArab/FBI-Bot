package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
)

func BanCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if !hasPermission(ctx, session, discordgo.PermissionBanMembers) {
		return nil
	}
	if len(ctx.GetArgs()) > 1 {
		var user *discordgo.User
		if len(ctx.GetMessage().Mentions) > 0 {
			user = ctx.GetMessage().Mentions[0]
		} else {
			var err error
			user, err = session.User(ctx.GetArgs()[0])
			if err != nil {
				return SendError(ctx, session, "That user does not exist!")
			}
		}
		if err := session.GuildBanCreateWithReason(ctx.GetGuild().ID, user.ID, ctx.GetArgs()[1], 0); err != nil {
			return SendError(ctx, session, err.Error())
		}
		_, _ = SendEmbed(ctx, session, &discordgo.MessageEmbed{
			URL:         "",
			Type:        "",
			Title:       "",
			Description: "",
			Timestamp:   "",
			Color:       0,
			Footer:      nil,
			Image:       nil,
			Thumbnail:   nil,
			Video:       nil,
			Provider:    nil,
			Author:      nil,
			Fields:      nil,
		})
		return nil
	}
	_, _ = session.ChannelMessageSend(ctx.GetChannel().ID, GetCommand("ban").Usage)
	return nil
}
