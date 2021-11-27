//A package for command routing with discordgo
package SpeedyCmds

import (
	"github.com/Jviguy/SpeedyCmds/command/commandMap"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

var CommandRegex = regexp.MustCompile(`(?m)("[^"]+"|[^\s"]+)`)

type Handler interface {
	Handle(s *discordgo.Session, msg *discordgo.MessageCreate)
	//Returns a *commandmap.Map in use for this Handler
	GetCommandMap() *commandMap.Map
}

//this is the basic premade handler with support for one prefix u can change it up how u want because of the whole interfaces
type PremadeHandler struct {
	dg   *discordgo.Session
	cmds *commandMap.Map
	//states wether to use a regex match or a simple split
	Regex  bool
	Prefix string
}

//initializes a basic pre made handler for you.
//noinspection ALL
func New(dg *discordgo.Session, cmds *commandMap.Map, Regex bool, prefix string) *PremadeHandler {
	h := &PremadeHandler{dg, cmds, Regex, prefix}
	dg.AddHandler(h.handle)
	return h
}

//called when a message is sent checks if its a bot and if it has the right prefix
func (h *PremadeHandler) handle(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}
	if strings.HasPrefix(msg.Content, h.Prefix) {
		var args []string
		args = strings.Split(strings.TrimPrefix(msg.Content, h.Prefix), " ")
		if h.Regex {
			args = CommandRegex.FindAllString(strings.TrimPrefix(msg.Content, h.Prefix), -1)
		}
		args, cmd := Shift(args, 0)
		err := h.cmds.Execute(cmd, ctx.New(args, msg, s), s)
		if err != nil {
			_, _ = s.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
				Title:       "An error occurred while executing that command",
				Description: err.Error(),
				Color:       16711680,
			})
		}
	}
}

func (h *PremadeHandler) GetCommandMap() *commandMap.Map {
	return h.cmds
}

func Shift(a []string, i int) ([]string, string) {
	b := a[i]
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a, b
}
