package irck

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/nlopes/slack"
)

type slackConfig struct {
    channel string // if doesnt start with # then direct nick
    token string
}

func SlackBridgeConfig(channel string, token string) (slackConfig) {
    return slackConfig{channel,token}
}

func RunSlackLoop(slackCfg slackConfig, evtChan chan Event) {
    sapi := slack.New(
        slackCfg.token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(
            os.Stdout,
            "slack-bot: ",
            log.Lshortfile|log.LstdFlags )), )
    rtm := sapi.NewRTM()
    go rtm.ManageConnection()

    for msg := range rtm.IncomingEvents {
		//fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello
		case *slack.ConnectedEvent:
			//fmt.Println("Infos:", ev.Info)
			//fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			fmt.Println("joining chan:", slackCfg.channel )
			rtm.SendMessage(
                rtm.NewOutgoingMessage("slackbridge online", slackCfg.channel) )
        case *slack.MessageEvent:
            fmt.Printf("Message: %v\n", ev)
            ee := MakeEvent("sl", "ev", "slacky", time.Now())
            fmt.Printf("ee: %s\n", ee.Body)
            // chans,_ := sapi.GetChannels(true)
        case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)
		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())
		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
        default:
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
    }

}

