package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"github.com/bwmarrin/discordgo"
)

var Token = os.Getenv("TOKEN")
var saved_scriptures = make(map[string]string)

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	message_array := strings.Split(m.Content, " ")

	if message_array[0] == "!heading" {

		_, found := saved_scriptures[message_array[1]]
		if !found {
			saved_scriptures[message_array[1]] = string(m.ReferencedMessage.Content)
			s.ChannelMessageSendReply(m.ChannelID, "Heading saved...", m.Reference())
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "Error: This heading is already saved!", m.Reference())
		}
	}
	if message_array[0] == "!show" {
		value, found := saved_scriptures[message_array[1]]
		if found {
			message, _:= s.ChannelMessageSendReply(m.ChannelID, value, m.Reference())
			s.MessageThreadStart(m.ChannelID, message.ID, message_array[1], 2)
			fmt.Println(message.ID)
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "Error: Heading not found!", m.Reference())
		}
	}
}