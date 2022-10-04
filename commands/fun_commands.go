package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"

	
)

type UrbanDictList struct {
	List []stuff
}
type stuff struct {
	Word       string
	Author     string
	Definition string
	Permalink  string
}

func UrbanCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) < 1 {
		return GetCommand("urban").SendUsage(ctx, session)
	}
	api, err := http.Get("http://api.urbandictionary.com/v0/define?term=" + ctx.GetArgs()[0])
	if err != nil {
		return SendError(ctx, session,"Could not make the request to the api!", err.Error())
	}
	defer api.Body.Close()
	if api.StatusCode != 200 {
		return SendError(ctx, session,"Could not get request from api")
	
	}
	body, err := ioutil.ReadAll(api.Body)
	if err != nil {
		return SendError(ctx, session,"Failed to read response body! ", err.Error())
	}
	var found UrbanDictList
	err = json.Unmarshal(body, &found)
	if err != nil {
		return SendError(ctx, session,"Error decoding json: ", err.Error())
	}
	if len(found.List) < 1 {
                return SendError(ctx, session,"The requested word is not defined in urban dictionary!")

	}
	urban := found.List[0]

	_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Title:       "Urban Dictionary",
		Description: "Word: " + urban.Word + "\n" + "Author: " + urban.Author + "\n" + "Definition: " + urban.Definition + "\n" + "Link: " + urban.Permalink,
	})
	return nil
}
