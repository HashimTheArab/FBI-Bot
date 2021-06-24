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

	commands.RegisterAll(dg)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	if err := dg.Open(); err != nil {
		panic(err)
	}

	_ = dg.UpdateListeningStatus("moans")

	fmt.Println("FBI Bot is now running!")

	select{}
}
