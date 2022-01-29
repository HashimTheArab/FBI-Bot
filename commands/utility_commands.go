package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func NukeCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	p, err := session.State.UserChannelPermissions(ctx.GetAuthor().ID, ctx.GetChannel().ID)
	if err != nil {
		return SendError(ctx, session, "Failed to retrieve user permissions! Error: "+err.Error())
	}
	if (p & discordgo.PermissionManageChannels) == 0 {
		return SendError(ctx, session, "In order to use this command, you need the \"Manage Channels\" permission!")
	}
	channel := ctx.GetChannel()
	if _, err := session.ChannelDelete(channel.ID); err != nil {
		return SendError(ctx, session, "Failed to nuke the channel, make sure I have perms!")
	}
	if channel == nil {
		return nil
	}
	ch, _ := session.GuildChannelCreateComplex(ctx.GetGuild().ID, discordgo.GuildChannelCreateData{
		Name:                 channel.Name,
		Type:                 channel.Type,
		Topic:                channel.Topic,
		Bitrate:              channel.Bitrate,
		UserLimit:            channel.UserLimit,
		RateLimitPerUser:     channel.RateLimitPerUser,
		Position:             channel.Position,
		PermissionOverwrites: channel.PermissionOverwrites,
		ParentID:             channel.ParentID,
		NSFW:                 channel.NSFW,
	})
	_, _ = session.ChannelMessageSendEmbed(ch.ID, &discordgo.MessageEmbed{
		Title: "Channel Nuked Successfully.",
		Color: 0x6b0000,
		Image: &discordgo.MessageEmbedImage{
			URL:    "https://imgur.com/GsxVDi7.gif",
			Width:  430,
			Height: 225,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Nuked by: " + ctx.GetAuthor().String(),
		},
	})
	return nil
}

func PurgeCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) < 1 {
		return GetCommand("purge").SendUsage(ctx, session)
	}
	num, err := strconv.Atoi(ctx.GetArgs()[0])
	if err != nil {
		return SendError(ctx, session, "Error: "+err.Error())
	}
	if num < 0 {
		return SendError(ctx, session, "Invalid purge amount. (1-100)")
	}
	if !hasPermission(ctx, session) {
		return nil
	}
	if ctx.GetMessage() != nil {
		if err := session.ChannelMessageDelete(ctx.GetChannel().ID, ctx.GetMessage().ID); err != nil {
			return SendError(ctx, session, "Failed to purge, error: "+err.Error())
		}
	}

	msgs, err := session.ChannelMessages(ctx.GetChannel().ID, num, "", "", "")
	var messages []string
	for _, value := range msgs {
		if !value.Pinned {
			messages = append(messages, value.ID)
		}
	}
	if err != nil {
		return SendError(ctx, session, "Failed to fetch messages, error: "+err.Error())
	}

	if err := session.ChannelMessagesBulkDelete(ctx.GetChannel().ID, messages); err != nil {
		return SendError(ctx, session, "Failed to purge, error: "+err.Error())
	}
	_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Title:       "Command ran successfully.",
		Description: strconv.Itoa(len(msgs)) + " messages removed!",
		Color:       0xACF9DC,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Purged by: " + ctx.GetAuthor().String(),
		},
	})
	return nil
}

func BackupCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	// todo
	return nil
}

func LoadCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	// todo
	return nil
}
