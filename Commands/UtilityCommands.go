package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func NukeCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	p, err := session.State.UserChannelPermissions(ctx.GetAuthor().ID, ctx.GetChannel().ID)
	if err != nil {
		SendError(ctx, session, "Failed to retrieve user permissions! Error: " + err.Error())
		return nil
	}
	if (p & discordgo.PermissionManageChannels) == 0 {
		SendError(ctx, session, "In order to use this command, you need the \"Manage Channels\" permission!")
		return nil
	}
	channel := ctx.GetChannel()
	if _, err := session.ChannelDelete(channel.ID); err != nil {
		SendError(ctx, session, "Failed to nuke the channel, make sure I have perms!")
		return nil
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
		Title:       "Kaboom! Pussy Nuked!",
		Color:       0x6b0000,
		Image:       &discordgo.MessageEmbedImage{
			URL:      "https://media.discordapp.net/attachments/814542881594671155/858199443056623656/nuke_picture.png",
			Width:    631,
			Height:   473,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Ran by: " + ctx.GetAuthor().String(),
		},
	})
	return nil
}

func PurgeCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) < 1 {
		SendInvalidUsage(ctx, session, "purge <amount>")
		return nil
	}
	num, err := strconv.Atoi(ctx.GetArgs()[0])
	if err != nil {
		SendError(ctx, session, "Error: " + err.Error())
		return nil
	}
	if num < 0 {
		SendError(ctx, session, "The accepted amount range is 1-100!")
		return nil
	}
	p, err := session.State.UserChannelPermissions(ctx.GetAuthor().ID, ctx.GetChannel().ID)
	if err != nil {
		SendError(ctx, session, "Failed to retrieve user permissions! Error: " + err.Error())
		return nil
	}
	if (p & discordgo.PermissionManageMessages) == 0 {
		SendError(ctx, session, "In order to use this command, you need the \"Manage Messages\" permission!")
		return nil
	}
	if ctx.GetMessage() != nil {
		if err := session.ChannelMessageDelete(ctx.GetChannel().ID, ctx.GetMessage().ID); err != nil {
			SendError(ctx, session, "Failed to purge, error: " + err.Error())
			return nil
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
		SendError(ctx, session, "Failed to fetch messages, error: " + err.Error())
		return nil
	}

	if err := session.ChannelMessagesBulkDelete(ctx.GetChannel().ID, messages); err != nil {
		SendError(ctx, session, "Failed to purge, error: " + err.Error())
		return nil
	}
	_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Title: "Purge Finished!",
		Description: strconv.Itoa(len(msgs)) + " messages purged!",
		Color: 0xACF9DC,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Ran by: " + ctx.GetAuthor().String(),
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