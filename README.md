[![GoDoc](https://godoc.org/github.com/mephux/komanda-cli?status.svg)](https://godoc.org/github.com/mephux/komanda-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/mephux/komanda-cli)](https://goreportcard.com/report/github.com/mephux/komanda-cli)

# Komanda CLI

This is the sister app of https://github.com/mephux/komanda.
I thought it would be fun so I did it. Komanda-cli is built using the awesome [gocui](https://github.com/jroimartin/gocui) package.

Would love some help to get it 1:1 with irssi.
Maybe embed lua,mruby or something else for the script lang.

# You Look Purdy

![komanda](http://i.imgur.com/UbBYVRq.png)
![Komanda-Channel](http://i.imgur.com/4vjrNxg.png)

## Download

  [Komanda Downloads](https://github.com/mephux/komanda-cli/releases)

## Usage

  ```bash
usage: komanda [<flags>]

The Komanda Command-line IRC Client

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -s, --ssl            enable ssl
  -i, --insecure       insecure ssl - skip verify. (self-signed certs)
  -h, --host=HOST      hostname
  -p, --port=PORT      port
  -n, --nick=NICK      nick
  -u, --user=USER      server user
  -P, --server-password=SERVER-PASSWORD
                       server password
      --nick-password=NICK-PASSWORD
                       nick password
  -a, --auto           auto-connect on startup.
  -c, --config=CONFIG  configuration file location
      --version        Show application version.
  ```

## Keyboard

  * `esc, right-arrow-key` change to next channel
  * `esc, left-arrow-key`  change to previous channel
  * `ctrl+n`               change to next window
  * `ctrl+p`               change to previous window
  * `ctrl+alt+p`           scroll up
  * `ctrl+alt+n`           scroll down
  * `page-up`              scroll up
  * `page-down`            scroll down
  * `tab`                  move to next active window
  * `enter`                scroll to bottom of window (if input is empty)
  * `/help`                for everything else

## /help output

```bash
* ==================== HELP COMMANDS ====================
* /exit  - exit komanda-cli
* /connect  - connect to irc using passed arguments
* /status  - status command
* /help  - help command
* /join <channel> - join irc channel
* /part [channel] - part irc channel or current if no channel given
* /clear  - clear current view
* /logo  - logo command
* /version  - version command
* /nick <nick> - nick irc channel
* /pass <password> - pass irc channel
* /raw <command> [data] - raw command
* /topic [channel] [topic] - set topic for given channel or current channel if empty
* /window <id,#channel> - change window example: /window #komanda
* /names  - list channel names
* /query <user> [message] - send private message to user
* /who <nick> - send who command to server
* /whois <nick> - send whois command to server
* /me [message] - send action message to channel
* /notice <channel/nick> <message> - send notice message to channel or nick
* /shrug  - Shrugging Emoji
* /tableflip  - TableFlip Emoji
* /kick <channel> <nick> [message] - kick user from channel. /kick #komanda mephux
* /away [message] - set status to away with a message or none to toggle away atatus
* ==================== HELP COMMANDS ====================
```

## Features

  * config file support (change colors, time formats etc.)
  * irc server auth
  * auto nickserv identify
  * auto-join channels
  * activity monitoring (new messages/highlights)
  * color nick
  * znc support
  * 256 colors
  * tab complete
  * new window per channel
  * history
  * cross-platform desktop notifications

## Config File Example

```toml
[Komanda]
  debug = false
  log_file = "/home/dweb/.komanda/komanda.log"

[Server]
  host = "irc.freenode.net"
  port = "6697" # Common NON-SSL Port 6667
  ssl = true
  insecure = false
  nick = "Komanda"
  user = "Komanda"
  nick_password = ""

  # if you server password and nick password are the same all you need
  # to set is this option.
  server_password = ""
  auto_connect = false

  # if you use a IRC bouncer this wont be needed
  channels = ["#komanda"]
  
  # filter all IRC join/part/quit messages
  filter_joinquit = false

# http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html
[Color]
  black = 0
  white = 15
  red = 160
  purple = 92
  logo = 75
  yellow = 11
  green = 119
  menu = 209
  my_nick = 119
  other_nick_default = 14
  timestamp = 247
  my_text = 129
  header = 57
  query_header = 11
  current_input_view = 215
  notice = 219
  action = 118

# https://golang.org/pkg/time/#pkg-constants
[Time]
  message_format = "15:04"
  notice_format = "02 Jan 06 15:04 MST"
  menu_format = "03:04:05 PM"
```

## TODO

  * Support for ban/op releated commands
  * IRC colors (for notice etc...)
    - https://en.wikipedia.org/wiki/Caret_notation
    - https://github.com/myano/jenni/wiki/IRC-String-Formatting

## Self-Promotion

Like komanda-cli? Follow the repository on
[GitHub](https://github.com/mephux/komanda-cli) and if
you would like to stalk me, follow [mephux](http://dweb.io/) on
[Twitter](http://twitter.com/mephux) and
[GitHub](https://github.com/mephux).
