package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"strconv"
)

func ServerCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	server := ctx.GetGuild()
	owner, err := session.User(server.OwnerID)
	ownerName := "None"
	if err == nil {
		ownerName = owner.Username
	}

	textChannels := 0
	voiceChannels := 0
	categories := 0
	for _, channel := range server.Channels {
		switch channel.Type {
		case discordgo.ChannelTypeGuildText:
			textChannels++
		case discordgo.ChannelTypeGuildVoice:
			voiceChannels++
		case discordgo.ChannelTypeGuildCategory:
			categories++
		}
	}
	_, err = SendEmbed(ctx, session, &discordgo.MessageEmbed{
		Title: server.Name,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Owner", Value: ownerName},
			{Name: "Total Channels", Value: strconv.Itoa(len(server.Channels))},
			{Name: "Total Categories", Value: strconv.Itoa(categories)},
			{Name: "Text Channels", Value: strconv.Itoa(textChannels)},
			{Name: "Voice Channels", Value: strconv.Itoa(voiceChannels)},
			{Name: "Members", Value: strconv.Itoa(server.MemberCount)},
			{Name: "Roles", Value: strconv.Itoa(len(server.Roles))},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: server.IconURL(),
		},
		Color: utils.Pink,
	})
	return err
}
