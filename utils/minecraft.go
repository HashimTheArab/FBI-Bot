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

type Response struct {
	ProfileUsers []struct {
		ID       string `json:"id"`
		Settings []struct {
			ID    string `json:"id"`
			Value string `json:"value"`
		} `json:"settings"`
	} `json:"profileUsers"`
}

type QueryResponse struct {
	Hostname   string
	Hostport   string
	Maxplayers string
	Players    []string
	MOTD       string
	Version    string
	Plugins    string
	Software   string
	Whitelist  string
	Map        string
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
				Text:    "Ran by " + ctx.GetAuthor().String() + " | Time: " + strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', 3, 64) + "s",
				IconURL: ctx.GetAuthor().AvatarURL(""),
			},
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Server Software", Value: a[7]},
				{Name: "Server Version", Value: a[3] + " (Protocol: " + a[2] + ")"},
				{Name: "MOTD", Value: StripColors(a[1])},
				{Name: "Players", Value: a[4] + "/" + a[5]},
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
				{Name: "Server Software", Value: r.GetStringResponse(r.Software)},
				{Name: "Server Version", Value: r.GetStringResponse(r.Version)},
				{Name: "Whitelist", Value: r.GetStringResponse(r.Whitelist)},
				{Name: "Server Plugins (" + pluginLength + ")", Value: pluginList},
				{Name: "Map", Value: r.GetStringResponse(r.Map)},
				{Name: "MOTD", Value: r.GetStringResponse(r.MOTD)},
				{Name: "Players (" + r.GetPlayerCount() + "/" + r.Maxplayers + ")", Value: players},
			},
			Color: Pink,
		}
		_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)
		embed.Fields = []*discordgo.MessageEmbedField{
			{Name: "Players (" + r.GetPlayerCount() + "/" + r.Maxplayers + ")", Value: rest},
		}
		_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)
	}
	return nil
}

func gopherQuery(address string) (QueryResponse, error) {
	data, err := query.Do(address)
	return QueryResponse{
		Hostname:   data["hostip"],
		Hostport:   data["hostport"],
		Maxplayers: data["maxplayers"],
		Players:    strings.Split(data["players"], ", "),
		MOTD:       StripColors(data["hostname"]),
		Version:    data["version"],
		Plugins:    data["plugins"],
		Software:   data["server_engine"],
		Whitelist:  data["whitelist"],
		Map:        data["map"],
	}, err
}

func (rsp *QueryResponse) GetStringResponse(field string) string {
	if field == "" {
		return "N/A"
	}
	return field
}

func (rsp QueryResponse) GetPlugins() (string, string) {
	if rsp.Plugins != "" {
		plugins := rsp.Plugins[strings.Index(rsp.Plugins, ":")+2:]
		list := strings.Split(plugins, "; ")
		return strconv.Itoa(len(list)), strings.Join(list, ", ")
	}
	return "0", "Server has plugin querying disabled."
}

func (rsp QueryResponse) GetPlayers() []string {
	if len(rsp.Players) > 0 && rsp.Players[0] != "" {
		return rsp.Players
	}
	return []string{"No players are online."}
}

func (rsp QueryResponse) GetPlayerCount() string {
	if rsp.Players[0] == "" {
		return "0"
	}
	return strconv.Itoa(len(rsp.Players))
}

func (x Response) Name() string {
	if name := x.ProfileUsers[0].Settings[6].Value; name != "" {
		return name
	}
	return "Unavailable"
}

func (x Response) Location() string {
	if location := x.ProfileUsers[0].Settings[8].Value; location != "" {
		return location
	}
	return "Unavailable"
}

func (x Response) Bio() string {
	if bio := x.ProfileUsers[0].Settings[7].Value; bio != "" {
		return bio
	}
	return "Unavailable"
}

func StripColors(text string) string {
	return re.ReplaceAllString(text, "")
}
