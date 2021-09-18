package commandMap

import (
	"github.com/Jviguy/SpeedyCmds/command"
	"github.com/Jviguy/SpeedyCmds/command/commandGroup"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/Jviguy/SpeedyCmds/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Map struct {
	commands map[string]command.Command
	disabled []string
	groups map[string]commandGroup.Group
}

const (
	DisabledCommand = "command"
	DisabledCategory = "category"
)

func (m *Map) RegisterCommandGroup(name string,group commandGroup.Group) {
	if !m.DoesGroupExist(name) && m.CanRegisterGroup(name) {
		m.groups[name] = group
	}
}

func (m *Map) GetGroups() map[string]commandGroup.Group {
	return m.groups
}

func (m *Map) GetGroup(name string) commandGroup.Group {
	if m.DoesGroupExist(name) {
		return m.groups[name]
	}
	return nil
}

func (m *Map) CanRegisterGroup(name string) bool {
	return m.commands[name].Execute == nil && m.GetGroup(name) == nil
}

func (m *Map) DoesGroupExist(name string) bool {
	_,b := m.groups[name]
	return b
}

func (m *Map) Execute(command string,c ctx.Ctx,s *discordgo.Session) error {
	switch true {
	case m.CanExecute(command):
		if m.Disabled(command){
			return sendDisabled(command, DisabledCommand, c, s)
		}
		category := m.commands[strings.ToLower(command)].Category
		if m.Disabled(category){
			return sendDisabled(category, DisabledCategory, c, s)
		}
		return m.commands[strings.ToLower(command)].Execute(c,s)
	case m.DoesGroupExist(command):
		if len(c.GetArgs()) > 0 {
			args,cmd := shift(c.GetArgs(),0)
			if m.GetGroup(command).CanExecute(cmd) {
				ct := ctx.New(args, c.GetMessage(), s)
				//custom ctx for the custom args needed
				return m.GetGroup(command).Execute(cmd, ct, s)
			}
		}
	default:
		for name, cmd := range m.commands {
			for _, alias := range cmd.Aliases {
				if alias == strings.ToLower(command) {
					if m.Disabled(name){
						return sendDisabled(command, DisabledCommand, c, s)
					}
					if m.Disabled(cmd.Category){
						return sendDisabled(cmd.Category, DisabledCategory, c, s)
					}
					return cmd.Execute(c,s)
				}
			}
		}
		_, _ = s.ChannelMessageSendEmbed(c.GetChannel().ID, &discordgo.MessageEmbed{
			Title: "Unknown Command: " + command,
			Description: "Did you mean: " + utils.FindClosest(command, utils.GetAllKeysCommands(m.GetAllCommands())),
		})
	}
	return nil
}

func (m *Map) GetAllCommands() map[string]command.Command {
	cs := m.GetCommands()
	for k,g := range m.GetGroups(){
		for name,cmd := range g.GetCommands(){
			cs[k+" "+name] = cmd
		}
	}
	return cs
}

func (m *Map) RegisterCommand(name string,command command.Command, override bool) {
	if m.CanRegisterCommand(name) || override {
		m.commands[strings.ToLower(name)] = command
	}
}

func (m *Map) CanRegisterCommand(name string) bool {
	return m.commands[name].Execute == nil && m.GetGroup(name) == nil
}

func (m *Map) Disable(name string){
	m.disabled = append(m.disabled, name)
}

func (m *Map) Enable(name string){
	var commands []string
	for _, v := range m.disabled {
		if v != name {
			commands = append(commands, v)
		}
	}
	m.disabled = commands
}

func (m *Map) Disabled(name string) bool {
	for _, v := range m.disabled {
		if v == name {
			return true
		}
	}
	return false
}

func (m *Map) GetDisabled() []string {
	return m.disabled
}

func (m *Map) GetCommands() map[string]command.Command {
	return m.commands
}

//noinspection ALL
func New() *Map {
	return &Map{commands: map[string]command.Command{},groups: map[string]commandGroup.Group{}}
}

func (m *Map) CanExecute(name string) bool {
	_, ok := m.commands[name]
	return ok
}

func shift(a []string,i int) ([]string,string) {
	b := a[i]
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a,b
}

func sendDisabled(name string, t string, ctx ctx.Ctx, session *discordgo.Session) error {
	_, err := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Description: "The " + t + " `" + name + "` is currently disabled.",
		Color:       16711680,
	})
	return err
}