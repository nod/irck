package main

/*
 * SERIOUSLY UGLY HACKY MESSING AROUND
 * USER BEWARE
 * THERE ARE NO TESTS
 * this is not TDD, this is CDD.  (Coffee|Curiosity) Driven Development
 */

import (
    "crypto/tls"
    "flag"
    "fmt"
    "time"

    "github.com/thoj/go-ircevent"
    "github.com/nod/irck/slackbridge"
    "github.com/nod/irck/event"
)

var ircChan = make(chan event.Event, 32)

type irckEvent struct {
    origin string
    body string
    author string
}

type ircConfig struct {
    channel string // if doesnt start with # then direct nick
    server string // addr:port
    nick string // joe
    ssl bool
    cmdPrefix string
    ircPrefix string
}

var ircCfg = ircConfig{}

func runIrcLoop(irccon *irc.Connection, c chan event.Event) {
    var ev event.Event
    for {
        select {
        case ev = <-c:
            irccon.Privmsg(ircCfg.channel, ev.Body)
            time.Sleep(time.Second)
        }
    }
}

func runIrc(ircOpts ircConfig) {
    irccon := irc.IRC(ircOpts.nick, "irck")
    irccon.VerboseCallbackHandler = true
    irccon.Debug = true
    irccon.UseTLS = ircOpts.ssl
    irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    irccon.AddCallback("001",
        func(e *irc.Event) { irccon.Join(ircOpts.channel) })
    irccon.AddCallback("366", func(e *irc.Event) {  })
    // irccon.AddCallback("PRIVMSG", routeIRC)
    err := irccon.Connect(ircOpts.server)
    if err != nil {
        fmt.Printf("Err %s", err )
        return
    }
    go runIrcLoop(irccon, ircChan)
    irccon.Loop()
}

type genericConfig struct {
    slacktoken string
    slackchan string
}

func setupCfg() (genericConfig) {
    genCfg := genericConfig{}
    // irc related config
    flag.StringVar(&ircCfg.server, "server", "", "irc host:port to connect to")
    flag.BoolVar(&ircCfg.ssl, "ssl", true, "irc ssl to server")
    flag.StringVar(&ircCfg.channel, "channel",
        "", "irc channel to connect to" )
    flag.StringVar(&ircCfg.nick, "nick",
        "", "irc nick for bot to use" )
    flag.StringVar(&genCfg.slacktoken, "slacktok",
        "", "slack api token")
    flag.StringVar(&genCfg.slackchan, "slackchan",
        "", "slack channel to join")
    flag.Parse()
    return genCfg
}

func main() {
    gencfg := setupCfg()
    fmt.Println("gencfg.slacktok", gencfg.slacktoken)
    fmt.Println("gencfg.slackchan", gencfg.slackchan)

    slackCfg := slackbridge.Config(gencfg.slackchan, gencfg.slacktoken)
    // go runIrc(ircCfg)
    slackbridge.RunSlackLoop(slackCfg, ircChan)
}

