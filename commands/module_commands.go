package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"strings"
)

func ModuleCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) < 1 {
		return GetCommand("module").SendUsage(ctx, session)
	}
	if !isPrim(ctx, session) {
		return nil
	}
	isCategory := func(name string) bool {
		for _, c := range Categories {
			if name == c {
				return true
			}
		}
		return false
	}
	args := ctx.GetArgs()
	switch args[0] {
	case "enable":
		if len(args) < 2 {
			_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
				Title:       "Invalid Usage!",
				Description: "Usage: " + handler.Prefix + "module enable <name>",
				Color:       utils.Red,
			})
			return err
		}
		name := ctx.GetArgs()[1]
		t := "command"
		if isCategory(name) {
			t = "category"
		}
		if handler.GetCommandMap().Disabled(name) {
			handler.GetCommandMap().Enable(name)
			_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
				Description: "The " + t + " `" + name + "` has been enabled!",
				Color:       utils.Green,
			})
			return err
		}
		return SendError(ctx, session, "There is no disabled command or category with the name `"+name+"`.")
	case "disable":
		if len(args) < 2 {
			_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
				Title:       "Invalid Usage!",
				Description: "Usage: " + handler.Prefix + "module disable <name>",
				Color:       utils.Red,
			})
			return err
		}
		name := ctx.GetArgs()[1]
		t := "command"
		if isCategory(name) {
			t = "category"
		}
		if handler.GetCommandMap().Disabled(name) {
			_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
				Description: "The " + t + " `" + name + "` is already disabled.",
				Color:       utils.Red,
			})
			return err
		}
		if _, ok := commands[name]; ok {
			handler.GetCommandMap().Disable(name)
			_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
				Description: "The command `" + name + "` has been disabled!",
				Color:       utils.Red,
			})
			return err
		}
		if isCategory(name) {
			handler.GetCommandMap().Disable(name)
			_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
				Description: "The category `" + name + "` has been disabled!",
				Color:       utils.Red,
			})
			return err
		}
		return SendError(ctx, session, "There is no command or category with the name `"+name+"`.")
	case "list":
		d := strings.Join(handler.GetCommandMap().GetDisabled(), ", ")
		if d == "" {
			d = "There are currently no commands/categories disabled."
		}
		_, err := SendEmbed(ctx, session, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: "Disabled commands/categories",
			},
			Description: d,
			Color:       utils.Pink,
		})
		return err
	}
	return GetCommand("module").SendUsage(ctx, session)
}
