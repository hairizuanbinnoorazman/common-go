package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/hairizuanbinnoorazman/common-go/chat"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Chat cli tool",
	Long:  `A quick chat cli to show an example of how some of the common modules can be used`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the chat cli",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
	},
}

var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Using slack command to send message to slack",
	Long:  `All soft has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		slackToken := "xoxo-auth-value"
		channelID := "channelID"
		chatService := chat.Slack{SlackToken: slackToken, ChannelID: channelID}
		var buffer bytes.Buffer
		for _, val := range args {
			buffer.WriteString(val)
		}
		client := http.Client{}
		err := chatService.SendMessage(context.TODO(), &client, buffer.String())
		if err != nil {
			fmt.Println("Error when sending message over")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(slackCmd)
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	execute()
}
