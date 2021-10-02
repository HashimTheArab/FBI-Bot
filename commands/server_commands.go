package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"strconv"
)

func ServerCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	server := ctx.GetGuild()
	fields := make([]*discordgo.MessageEmbedField, 7)
	owner, err := session.User(server.OwnerID)
	ownerName := "None"
	if err == nil {
		ownerName = owner.Username
	}
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Owner", Value: ownerName})
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Total Channels", Value: strconv.Itoa(len(server.Channels))})
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
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Total Categories", Value: strconv.Itoa(categories)})
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Text Channels", Value: strconv.Itoa(textChannels)})
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Voice Channels", Value: strconv.Itoa(voiceChannels)})
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Members", Value: strconv.Itoa(server.MemberCount)})
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Roles", Value: strconv.Itoa(len(server.Roles))})
	_, err = SendEmbed(ctx, session, &discordgo.MessageEmbed{
		Title:  server.Name,
		Fields: fields,
		Image: &discordgo.MessageEmbedImage{
			URL: server.IconURL(),
		},
		Color: utils.Pink,
	})
	return err
}
