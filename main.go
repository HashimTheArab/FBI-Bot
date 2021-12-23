package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/commands"
	"github.com/prim69/fbi-bot/utils/settings"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dg, err := discordgo.New("Bot " + settings.Data.Token)

	if err != nil {
		panic(err)
	}

	dg.AddHandler(onMessageDelete)
	dg.AddHandler(onMessageUpdate)

	commands.RegisterAll(dg)

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	if err := dg.Open(); err != nil {
		panic(err)
	}

	dg.State.MaxMessageCount = 30000

	_ = dg.UpdateListeningStatus("moans")

	fmt.Println("FBI Bot is now running!")

	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		settings.Data.DisabledCommands = commands.GetHandler().GetCommandMap().GetDisabled()
		settings.Save()
		os.Exit(0)
	}()

	select {}
}

func onMessageDelete(_ *discordgo.Session, msg *discordgo.MessageDelete) {
	b := msg.BeforeDelete
	var attachment *discordgo.MessageAttachment
	if len(b.Attachments) > 0 {
		attachment = b.Attachments[0]
	}
	var list []*commands.SnipedMessage
	list = append(list, &commands.SnipedMessage{
		Content:    b.Content,
		Author:     b.Author,
		ChannelID:  b.ChannelID,
		ID:         b.ID,
		Timestamp:  b.Timestamp,
		Attachment: attachment,
	})
	for _, value := range commands.Snipes[b.ChannelID] {
		list = append(list, value)
	}
	commands.Snipes[b.ChannelID] = list
}

func onMessageUpdate(_ *discordgo.Session, msg *discordgo.MessageUpdate) {
	b := msg.BeforeUpdate
	var attachment *discordgo.MessageAttachment
	if len(b.Attachments) > 0 {
		attachment = b.Attachments[0]
	}
	var list []*commands.SnipedMessage
	list = append(list, &commands.SnipedMessage{
		Content:    b.Content,
		NewContent: msg.Content,
		Author:     b.Author,
		ChannelID:  b.ChannelID,
		ID:         b.ID,
		Timestamp:  b.Timestamp,
		Attachment: attachment,
	})
	for _, value := range commands.EditSnipes[b.ChannelID] {
		list = append(list, value)
	}
	commands.EditSnipes[b.ChannelID] = list
}
