package events

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
)

// Ready event
func Ready(s *discord.Session, event *discord.Ready) {
	s.UpdateStatus(0, "Go, Discord, go!")
	fmt.Println("Bot is ready to GO!")
}
