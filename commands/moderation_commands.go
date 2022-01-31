package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"strings"
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
		reason := strings.Join(ctx.GetArgs()[1:], " ")
		if err := session.GuildBanCreateWithReason(ctx.GetGuild().ID, user.ID, reason, 0); err != nil {
			return SendError(ctx, session, err.Error())
		}
		_, _ = SendEmbed(ctx, session, &discordgo.MessageEmbed{
			Description: "**Reason:** " + reason,
			Color:       utils.Purple,
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Banned by: " + ctx.GetAuthor().String(),
				IconURL: ctx.GetAuthor().AvatarURL(""),
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name: "User Banned!",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: user.AvatarURL(""),
			},
		})
		return nil
	}
	_, _ = session.ChannelMessageSend(ctx.GetChannel().ID, GetCommand("ban").Usage)
	return nil
}
