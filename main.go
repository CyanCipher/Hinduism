package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

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

	if message_array[0] == "!heading" && len(message_array) > 1 {

		_, found := saved_scriptures[message_array[1]]
		if !found {
			if m.ReferencedMessage != nil {
				saved_scriptures[message_array[1]] = string(m.ReferencedMessage.Content)
				s.ChannelMessageSendReply(m.ChannelID, "Heading saved...", m.Reference())
			}
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "Error: This heading is already saved!", m.Reference())
		}
	}
	if message_array[0] == "!show" {
		value, found := saved_scriptures[message_array[1]]
		if found {
			if ch, err := s.State.Channel(m.ChannelID); err != nil || !ch.IsThread() {
				thread, err := s.MessageThreadStartComplex(m.ChannelID, m.ID, &discordgo.ThreadStart{
					Name: 					message_array[1],
					AutoArchiveDuration: 	60,
					Invitable: 				false,
					RateLimitPerUser: 		10,
				})
				if err != nil {
					fmt.Println(err)
					return
				}
				_, _ = s.ChannelMessageSend(thread.ID, value)
				m.ChannelID = thread.ID
			} else {
				_, _ = s.ChannelMessageSendReply(m.ChannelID, value, m.Reference())
			}
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "Error: Heading not found!", m.Reference())
		}
	}
	if message_array[0] == "!ask" && m.ChannelID == "1174952393503428701" {
		buffer, _ := s.ChannelMessageSendReply(m.ChannelID, "Scanning Database...    ", m.Reference())
		query := strings.Join(message_array[1:], " ")
		response := ask_krishna(query, "n")
		s.ChannelMessageDelete(m.ChannelID, buffer.ID)
		s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
			Type: discordgo.EmbedTypeRich,
			Title: "Hinduism Bot",
			Description: response,
			Color: 973812,
		}, m.Reference())
	}

	if message_array[0] == "!ref" && m.ChannelID == "1174952393503428701" {
		buffer, _ := s.ChannelMessageSendReply(m.ChannelID, "Scanning Database...    ", m.Reference())
		query := strings.Join(message_array[1:], " ")
		response := ask_krishna(query, "r")
		s.ChannelMessageDelete(m.ChannelID, buffer.ID)
		s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
			Type: discordgo.EmbedTypeRich,
			Title: "Hinduism Bot",
			Description: response,
			Color: 973812,
		}, m.Reference())
	}

	if message_array[0] == ".." && m.ChannelID == "1098091395446755358" {
		buffer, _ := s.ChannelMessageSendReply(m.ChannelID, "Generating Response.    ..", m.Reference())
		query := strings.Join(message_array[1:], " ")
		response := ask_krishna(query, "g")
		s.ChannelMessageDelete(m.ChannelID, buffer.ID)
		s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
			Type: discordgo.EmbedTypeRich,
			Title: "Bhagavad Gita",
			Description: response,
			Color: 973812,
		}, m.Reference())
	}

	if message_array[0] == "!translate" {
		if m.ReferencedMessage != nil {
			buffer, _ := s.ChannelMessageSendReply(m.ChannelID, "Translating...", m.Reference())
			response := ask_krishna(m.ReferencedMessage.Content, "t")
			s.ChannelMessageDelete(m.ChannelID, buffer.ID)
			s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
				Type: discordgo.EmbedTypeRich,
				Title: "Translation",
				Description: response,
				Color: 973812,
			}, m.Reference())
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "Nothing to translate here...", m.Reference())
		}
	}

}

func ask_krishna(input string, option string) string { f, err := os.Create("query.txt")
	if err != nil {
		panic(err)
	}
	l, err := f.WriteString(input)
	if err != nil {
		f.Close()
		panic(err)
	}

	fmt.Println(l)
	err = f.Close()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("python", "chat.py", option)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(out)
}
