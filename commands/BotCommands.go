package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
)

func PingCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	msg, _ := session.ChannelMessageSend(ctx.GetChannel().ID, "blow me")
	m1, _ := ctx.GetMessage().Timestamp.Parse()
	m2, _ := msg.Timestamp.Parse()
	_, _ = session.ChannelMessageEdit(ctx.GetChannel().ID, msg.ID, "Latency: " + m2.Sub(m1).String() + "\nAPI Latency: " + session.HeartbeatLatency().String())
	return nil
}