package utils

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/sandertv/go-raknet"
	"github.com/sandertv/gophertunnel/query"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type QueryResponse struct {
	Hostname    string
	Hostport    string
	Maxplayers  string
	Players     []string
	MOTD        string
	Version     string
	Plugins     string
	Software    string
	Whitelist   string
	Map         string
}

var re = regexp.MustCompile(`(?i)§[0-9A-GK-OR§]`)

const (
	ShortQuery QueryType = iota
	LongQuery
)

type QueryType int

func Query(ip, port string, ctx ctx.Ctx, session *discordgo.Session, t QueryType) error {
	args := ctx.GetArgs()
	if len(args) > 1 {
		if p, err := strconv.Atoi(args[1]); err == nil {
			if p >= 0 && p <= 65535 {
				port = args[1]
			}
		}
	}
	start := time.Now()
	if t == ShortQuery {
		b, err := raknet.Ping(ip + ":" + port)
		if err != nil {
			return err
		}
		a := strings.Split(string(b), ";")
		_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
			Title: "Ping Response for " + args[0] + "!",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Ran by " + ctx.GetAuthor().String() + " | Time: " + strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', 3, 64) + "s",
				IconURL: ctx.GetAuthor().AvatarURL(""),
			},
			Fields: []*discordgo.MessageEmbedField{
				{Name: "🖇 Software", Value: a[7]},
				{Name: "💾 Version", Value: a[3] + " (Protocol: " + a[2] + ")"},
				{Name: "🎉 MOTD", Value: StripColors(a[1])},
				{Name: "👥 Players", Value: a[4] + "/" + a[5]},
			},
			Color: Pink,
		})
	} else {
		r, err := gopherQuery(ip + ":" + port)
		if err != nil {
			return err
		}
		var rest string
		players := strings.Join(r.GetPlayers(), ", ")
		if len(players) > 800 {
			rest = players[800:]
			players = players[:800]
		}
		pluginLength, pluginList := r.GetPlugins()
		embed := &discordgo.MessageEmbed{
			Title: "Query Response for " + ip + ":" + port + "!",
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Ran by " + ctx.GetAuthor().String() + " | Time: " + strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', 3, 64) + "s",
				IconURL: ctx.GetAuthor().AvatarURL(""),
			},
			Fields: []*discordgo.MessageEmbedField{
				{Name: "🖇 Software", Value: r.GetStringResposne(r.Software)},
				{Name: "💾 Version", Value: r.GetStringResposne(r.Version)},
				{Name: "🏳 Whitelist", Value: r.GetStringResposne(r.Whitelist)},
				{Name: "💻 Plugins (" + pluginLength + ")", Value: pluginList},
				{Name: "🗺 Map", Value: r.GetStringResposne(r.Map)},
				{Name: "🎉 MOTD", Value: r.GetStringResposne(r.MOTD)},
				{Name: "👥 Players (" + r.GetPlayerCount() + "/" + r.Maxplayers + ")", Value: players},
			},
			Color: Pink,
		}
		_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)
		embed.Fields = []*discordgo.MessageEmbedField{
			{Name: "👥 Players (" + r.GetPlayerCount() + "/" + r.Maxplayers + ")", Value: rest},
		}
		_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)
	}
	return nil
}

func gopherQuery(address string) (QueryResponse, error) {
	data, err := query.Do(address)
	return QueryResponse{
		Hostname:    data["hostip"],
		Hostport:    data["hostport"],
		Maxplayers:  data["maxplayers"],
		Players:     strings.Split(data["players"], ", "),
		MOTD:        StripColors(data["hostname"]),
		Version:     data["version"],
		Plugins:     data["plugins"],
		Software:    data["server_engine"],
		Whitelist:   data["whitelist"],
		Map:         data["map"],
	}, err
}

func (rsp *QueryResponse) GetStringResposne(field string) string {
	if field == "" {
		return "N/A"
	}
	return field
}

func (rsp QueryResponse) GetPlugins() (string, string) {
	if rsp.Plugins != "" {
		plugins := rsp.Plugins[strings.Index(rsp.Plugins, ":") + 2:]
		list := strings.Split(plugins, "; ")
		return strconv.Itoa(len(list)), strings.Join(list, ", ")
	}
	return "0", "This server has plugin query disabled."
}

func (rsp QueryResponse) GetPlayers() []string {
	if len(rsp.Players) > 0 && rsp.Players[0] != "" {
		return rsp.Players
	}
	return []string{"There are no players online!"}
}

func (rsp QueryResponse) GetPlayerCount() string {
	if rsp.Players[0] == "" {
		return "0"
	}
	return strconv.Itoa(len(rsp.Players))
}

func StripColors(text string) string {
	return re.ReplaceAllString(text, "")
}