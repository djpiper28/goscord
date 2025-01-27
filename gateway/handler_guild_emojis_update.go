package gateway

import "github.com/Goscord/goscord/gateway/event"

type GuildEmojisUpdateHandler struct{}

func (_ *GuildEmojisUpdateHandler) Handle(s *Session, data []byte) {
	_, err := event.NewGuildEmojisUpdate(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : updated emojis are not added to the state

	//s.Bus().Publish("guildEmojisUpdate", guild)
}
