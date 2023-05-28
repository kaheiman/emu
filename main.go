package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Set your Discord bot token
	token := "<bot_token>"

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Open a websocket connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection:", err)
		return
	}

	// Set the server ID where you want to create the channel
	serverID := "<server_id>"

	// Delete the channel
	// _, err = dg.ChannelDelete("1112241859071324202")
	// if err != nil {
	// 	fmt.Println("Error deleting channel:", err)
	// 	return
	// }

	// Channel deleted successfully
	fmt.Println("Channel deleted successfully.")

	// Set the channel name and type
	channelName := "<channel_name>"
	channelType := discordgo.ChannelTypeGuildText // For a text channel, use discordgo.ChannelTypeGuildText

	// Set channel permissions to allow everyone to read and send messages
	permissionOverwrite := discordgo.PermissionOverwrite{
		ID:    serverID,
		Type:  discordgo.PermissionOverwriteTypeRole,
		Allow: discordgo.PermissionReadMessages | discordgo.PermissionSendMessages,
	}
	var permissionOverwrites []*discordgo.PermissionOverwrite
	permissionOverwrites = append(permissionOverwrites, &permissionOverwrite)

	// Create the channel
	channel, err := dg.GuildChannelCreateComplex(serverID, discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 channelType,
		UserLimit:            20,
		ParentID:             "",                   // Set the parent ID if you want the channel to be in a specific category, otherwise leave it as an empty string
		PermissionOverwrites: permissionOverwrites, // Set permission overwrites if needed, otherwise leave it as nil
		NSFW:                 false,
	})

	if err != nil {
		fmt.Println("Error creating channel:", err)
		return
	}

	// Channel created successfully
	fmt.Printf("Channel complex created successfully: \n %+v", channel)
	// Close the Discord session when you're done
	defer dg.Close()
}
