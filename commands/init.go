package commands

import (
	"github.com/bwmarrin/discordgo"
	discord "github.com/bwmarrin/discordgo"
)

// Context command context
type Context struct {
	Session *discord.Session
	Message *discord.MessageCreate
	Guild   *discord.Guild
	Channel *discord.Channel
	Author  *discord.User
	Args    []string
	Prefix  string
}

// CommandMap is a map that gets the user's command input and retrieves its respective function
var CommandMap = make(map[string]*Command)

// AliasMap finds the commands of each alias
var AliasMap = make(map[string]string)

// Command is a command object
type Command struct {
	Name        string
	Description string
	Aliases     []string
	Dev         bool
	Run         func(*Context)
}

// Send a message to the channel
func (ctx *Context) Send(content string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Channel.ID, content)
}

// SendComplex an embed/complex message to the channel
func (ctx *Context) SendComplex(content string, embed *discord.MessageEmbed) (*discord.Message, error) {
	data := &discord.MessageSend{Content: content, Embed: embed}
	return ctx.Session.ChannelMessageSendComplex(ctx.Channel.ID, data)
}

// NewCommand creates a new command
func NewCommand(name, description string) (cmd *Command, existing bool) {
	_, existing = CommandMap[name]
	if existing {
		return nil, existing
	}
	cmd = &Command{Name: name, Description: description}
	return cmd, existing
}

// RegisterCommand adds the command to the CommandMap
func RegisterCommand(cmd *Command) {
	CommandMap[cmd.Name] = cmd
	for _, alias := range cmd.Aliases {
		AliasMap[alias] = cmd.Name
	}
}

// UnregisterCommand removes the command from the CommandMap
func UnregisterCommand(cmd string) {
	delete(CommandMap, cmd)
}
