package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type SnipedMessage struct {
	Content string
	NewContent string // only applies to edited messages
	Author *discordgo.User
	ChannelID string
	ID string
	Timestamp discordgo.Timestamp
	Attachment *discordgo.MessageAttachment
}

var Messages = map[string]*SnipedMessage{}
var Snipes = map[string][]*SnipedMessage{}
var EditSnipes = map[string][]*SnipedMessage{}

func SnipeCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(Snipes[ctx.GetChannel().ID]) < 1 {
		_, _ = session.ChannelMessageSend(ctx.GetChannel().ID, "There's nothing to snipe! *uwu*")
		return nil
	}

	var num int
	if len(ctx.GetArgs()) > 0 {
		if n, err := strconv.Atoi(ctx.GetArgs()[0]); err == nil {
			num = n - 1
		}
	}

	if num < 0 || (num + 1) > len(Snipes[ctx.GetChannel().ID]) {
		num = 0
	}

	msg := Snipes[ctx.GetChannel().ID][num]

	var image *discordgo.MessageEmbedImage
	if msg.Attachment != nil {
		image = &discordgo.MessageEmbedImage{
			URL:      msg.Attachment.URL,
			ProxyURL: msg.Attachment.ProxyURL,
			Width:    msg.Attachment.Width,
			Height:   msg.Attachment.Height,
		}
	}

	var time string
	t, err := msg.Timestamp.Parse()
	if err == nil {
		time = t.Format("January-02-2006 3:04:05 PM MST")
	} else {
		time = "Unavailable"
	}

	_, _ = SendEmbed(ctx, session, discordgo.MessageEmbed{
		Description: msg.Content,
		Color: 0xD8CDE9,
		Footer: &discordgo.MessageEmbedFooter{
			Text: strconv.Itoa(num + 1) + "/" + strconv.Itoa(len(Snipes[ctx.GetChannel().ID])) + " | " + time,
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: msg.Author.Username + "#" + msg.Author.Discriminator,
			IconURL: msg.Author.AvatarURL(""),
		},
		Image: image,
	})
	return nil
}

func EditSnipeCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(EditSnipes[ctx.GetChannel().ID]) < 1 {
		_, _ = session.ChannelMessageSend(ctx.GetChannel().ID, "There's nothing to snipe! *uwu*")
		return nil
	}

	var num int
	if len(ctx.GetArgs()) > 0 {
		if n, err := strconv.Atoi(ctx.GetArgs()[0]); err == nil {
			num = n - 1
		}
	}

	if num < 0 || (num + 1) > len(EditSnipes[ctx.GetChannel().ID]) {
		num = 0
	}

	msg := EditSnipes[ctx.GetChannel().ID][num]

	var time string
	t, err := msg.Timestamp.Parse()
	if err == nil {
		time = t.Format("January-02-2006 3:04:05 PM MST")
	} else {
		time = "Unavailable"
	}

	var image *discordgo.MessageEmbedImage
	if msg.Attachment != nil {
		image = &discordgo.MessageEmbedImage{
			URL:      msg.Attachment.URL,
			ProxyURL: msg.Attachment.ProxyURL,
			Width:    msg.Attachment.Width,
			Height:   msg.Attachment.Height,
		}
	}

	_, _ = SendEmbed(ctx, session, discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Before",
				Value: msg.Content,
			},
			{
				Name: "After",
				Value: msg.NewContent,
			},
		},
		Color: 0xD8CDE9,
		Footer: &discordgo.MessageEmbedFooter{
			Text: strconv.Itoa(num + 1) + "/" + strconv.Itoa(len(EditSnipes[ctx.GetChannel().ID])) + " | " + time,
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: msg.Author.Username + "#" + msg.Author.Discriminator,
			IconURL: msg.Author.AvatarURL(""),
		},
		Image: image,
	})
	return nil
}