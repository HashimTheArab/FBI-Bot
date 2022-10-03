package commands

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"

	_ "net/http"
)

type UrbanDictList struct {
	List []stuff
}
type stuff struct {
	Word       string
	Author     string
	Definition string
	Permalink       string
}

func UrbanCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) > 1 {
	_, _ = session.ChannelMessageSend(ctx.GetChannel().ID, "!urban <word>")
	return nil
	}
	api, err := http.Get("http://api.urbandictionary.com/v0/define?term=" + ctx.GetArgs()[0])
	if err != nil {
	fmt.Println("Could not make the request to the api!", err.Error())
	return nil
	}
	defer api.Body.Close()
	if api.StatusCode != 200 {
	fmt.Println("Could not get request from api")
	return nil
	}
	body, err := ioutil.ReadAll(api.Body)
	if err != nil {
	fmt.Println("Failed to read response body! ", err.Error())
	return nil
	}
	var found UrbanDictList
	err = json.Unmarshal(body, &found)
	if err != nil {
	fmt.Println("Error decoding json: ", err.Error())
	return nil
	}
	if len(found.List) < 1 {
	_, _ = session.ChannelMessageSend(ctx.GetChannel().ID, "Thats not a word!")
	fmt.Println("The requested word is not defined in urban dictionary!")

	}
	urban := found.List[0]

	_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
	Title:       "Urban Dictionary",
	Description: "Word: " + urban.Word + "\n" + "Author: " + urban.Author + "\n" + "Definition: " + urban.Definition + "\n" + "Link: " + urban.Permalink,
	})
	return nil
}
