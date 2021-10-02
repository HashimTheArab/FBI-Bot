package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
)

func AvatarCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	user := ctx.GetAuthor()
	if len(ctx.GetArgs()) > 0 {
		if len(ctx.GetMessage().Mentions) > 0 {
			user = ctx.GetMessage().Mentions[0]
		} else {
			var err error
			user, err = session.User(ctx.GetArgs()[0])
			if err != nil {
				return SendError(ctx, session, "That user does not exist!")
			}
		}
	}
	_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
		Title: user.Username + "'s avatar",
		Image: &discordgo.MessageEmbedImage{
			URL: user.AvatarURL("2048"),
		},
		Color: utils.Pink,
	})
	return err
}

func WhoIsCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	user := ctx.GetAuthor()
	if len(ctx.GetArgs()) > 0 {
		if len(ctx.GetMessage().Mentions) > 0 {
			user = ctx.GetMessage().Mentions[0]
		} else {
			var err error
			user, err = session.User(ctx.GetArgs()[0])
			if err != nil {
				return SendError(ctx, session, "That user does not exist!")
			}
		}
	}
	_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
		Title: user.Username + "'s avatar",
		Image: &discordgo.MessageEmbedImage{
			URL: user.AvatarURL("2048"),
		},
		Color: utils.Pink,
	})
	return err
}