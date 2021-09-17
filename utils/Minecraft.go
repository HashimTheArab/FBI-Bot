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
	Plugins     []string
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
		_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
			Title: "Query Response for " + ip + ":" + port + "!",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Ran by " + ctx.GetAuthor().String() + " | Time: " + strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', 3, 64) + "s",
				IconURL: ctx.GetAuthor().AvatarURL(""),
			},
			Fields: []*discordgo.MessageEmbedField{
				{Name: "🖇 Software", Value: r.GetSoftware()},
				{Name: "💾 Version", Value: r.GetVersion()},
				{Name: "🏳 Whitelist", Value: r.GetWhitelist()},
				{Name: "💻 Plugins (" + strconv.Itoa(r.GetPluginLength()) + ")", Value: strings.Join(r.GetPlugins(), ", ")},
				{Name: "🗺 Map", Value: r.GetMap()},
				{Name: "🎉 MOTD", Value: r.GetMOTD()},
				{Name: "👥 Players (" + r.GetPlayerCount() + "/" + r.Maxplayers + ")", Value: strings.Join(r.GetPlayers(), ", ")},
			},
			Color: Pink,
		})
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
		Plugins:     strings.Split(data["plugins"], ", "),
		Software:    data["server_engine"],
		Whitelist:   data["whitelist"],
		Map:         data["map"],
	}, err
}

func (rsp QueryResponse) GetSoftware() string {
	if sf := rsp.Software; sf != "" {
		return rsp.Software
	}
	return "Invalid Response"
}

func (rsp QueryResponse) GetVersion() string {
	if vr := rsp.Version; vr != "" {
		return rsp.Version
	}
	return "Invalid Response"
}

func (rsp QueryResponse) GetPluginLength() int {
	pl := strings.Split(strings.Join(rsp.Plugins, ""), ";")
	return len(pl)
}

func (rsp QueryResponse) GetWhitelist() string {
	if sf := rsp.Whitelist; sf != "" {
		return rsp.Whitelist
	}
	return "Invalid Response"
}

func (rsp QueryResponse) GetMap() string {
	if mp := rsp.Map; mp != "" {
		return rsp.Map
	}
	return "Invalid Response"
}

func (rsp QueryResponse) GetMOTD() string {
	if mt := rsp.MOTD; mt != "" {
		return rsp.MOTD
	}
	return "Invalid Response"
}

func (rsp QueryResponse) GetPlugins() []string {
	sf := strings.Split(strings.Join(rsp.Plugins[:], ""), ":")
	if len(sf) > 1 {
		return sf[1:]
	}
	return []string{"This server has plugin query disabled."}
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