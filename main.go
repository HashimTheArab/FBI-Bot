package main

import (
	"FBI/commands"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main(){
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		panic(err)
	}

	dg.AddHandler(onMessageDelete)
	dg.AddHandler(onMessageUpdate)
	dg.AddHandler(onMessageCreate)

	commands.RegisterAll(dg)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	if err := dg.Open(); err != nil {
		panic(err)
	}

	_ = dg.UpdateListeningStatus("moans")

	fmt.Println("FBI Bot is now running!")

	select{}
}

func onMessageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if _, ok := commands.Snipes[msg.ChannelID]; !ok {
		//commands.Snipes[msg.ChannelID] = make([]*discordgo.Message, 15)
	}

	var attachment *discordgo.MessageAttachment
	if len(msg.Attachments) > 0 {
		attachment = msg.Attachments[0]
	}

	commands.Messages[msg.ID] = &commands.SnipedMessage{Content: msg.Content, Author: msg.Author, ChannelID: msg.ChannelID, ID: msg.ID, Timestamp: msg.Timestamp, Attachment: attachment}
}

func onMessageDelete(session *discordgo.Session, msg *discordgo.MessageDelete) {
	if m, ok := commands.Messages[msg.ID]; ok {
		var list []*commands.SnipedMessage
		list = append(list, m)
		for _, value := range commands.Snipes[msg.ChannelID] {
			list = append(list, value)
		}
		commands.Snipes[msg.ChannelID] = list
		delete(commands.Messages, msg.ID)
	}
}

func onMessageUpdate(session *discordgo.Session, msg *discordgo.MessageUpdate) {
	if m, ok := commands.Messages[msg.ID]; ok {
		m.NewContent = msg.Content
		var list []*commands.SnipedMessage
		list = append(list, m)
		for _, value := range commands.EditSnipes[msg.ChannelID] {
			list = append(list, value)
		}
		commands.EditSnipes[msg.ChannelID] = list
		delete(commands.Messages, msg.ID)
	}
}