package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"math"
	"strconv"
	"time"
)

func PingCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) > 0 {
		return utils.Query(ctx.GetArgs()[0], "19132", ctx, session, utils.ShortQuery)
	}
	msg, _ := session.ChannelMessageSend(ctx.GetChannel().ID, "blow me")
	m1, _ := ctx.GetMessage().Timestamp.Parse()
	m2, _ := msg.Timestamp.Parse()
	_, _ = session.ChannelMessageEdit(ctx.GetChannel().ID, msg.ID, "Latency: " + m2.Sub(m1).String() + "\nAPI Latency: " + session.HeartbeatLatency().String())
	return nil
}

func StatsCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	uptime := time.Now().Unix() - UpTime.Unix()

	day := math.Floor(float64(uptime / 86400))
	hourSeconds := uptime % 86400
	hour := math.Floor(float64(hourSeconds / 3600))
	minuteSec := hourSeconds % 3600
	minute := math.Floor(float64(minuteSec / 60))
	remainingSec := minuteSec % 60
	second := math.Ceil(float64(remainingSec))

	days := strconv.FormatFloat(day, 'f', 0, 64)
	hours := strconv.FormatFloat(hour, 'f', 0, 64)
	minutes := strconv.FormatFloat(minute, 'f', 0, 64)
	seconds := strconv.FormatFloat(second, 'f', 0, 64)
	var ts string

	if int(day) == 1 {
		ts += days + " day, "
	} else {
		ts += days + " days, "
	}

	if int(hour) == 1 {
		ts += hours + " hour, "
	} else {
		ts += hours + " hours, "
	}

	if int(minute) == 1 {
		ts += minutes + " minute, "
	} else {
		ts += minutes + " minutes, "
	}

	if int(second) == 1 {
		ts += "and " + seconds + " second"
	} else {
		ts += "and " + seconds + " seconds"
	}

	_, _ = SendEmbed(ctx, session, &discordgo.MessageEmbed{
		Color:       0xF0BBCE,
		Author:      &discordgo.MessageEmbedAuthor{
			URL:          utils.DiscordLink,
			Name:         "Bot Information",
			IconURL:      "https://media.discordapp.net/attachments/814542881594671155/858957822918524978/unknown.png?width=458&height=473",
		},
		Fields:      []*discordgo.MessageEmbedField{
			{
				Name: "Uptime",
				Value: session.State.User.Username + " has been online for " + ts,
			},
			{
				Name: "Servers",
				Value: session.State.User.Username + " is in " + strconv.Itoa(len(session.State.Guilds)) + " servers!",
			},
			{
				Name: "About",
				Value: session.State.User.Username + " is coded by **" + utils.Author + "** in the Go programming language.\n" + "This bot is open source, the source code can be found at " + utils.GithubLink,
			},
		},
	})
	return nil
}