# emersyx_apis [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url] [![GoDoc][godoc-img]][godoc-url]

This repository has been archived and is not under development anymore. The relevant code has been merged into the main
[emersyx repository][emersyx-repo]. This repository now serves for historical purposes only, as it contains the code for
emersyx version 0.1.

APIs and interface specifications for emersyx modules. The following go modules are present:

* `emcomapi` - common APIs for platform components (e.g. messages, receptors, routers, processors)
* `emircapi` - APIs defining the common interface for an IRC gateway and the IRC event type
* `emtgapi` - APIs defining the common interface for a Telegram gateway and the Telegram event types

## TODO

The following websites/services/APIs are on my TODO list to add support for:

* [VirusTotal Public API v2.0][1] as described in the [official documentation][2]
* [Reddit API][3]
* [GMail API][4], though probably based on the [go library][5]
* [Slack Bot Users][6]
* [Discord API][7]
* [Github API][8]

[build-img]: https://travis-ci.org/emersyx/emersyx_apis.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_apis
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_apis
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_apis
[godoc-img]: https://godoc.org/emersyx.net/emersyx_apis?status.svg
[godoc-url]: https://godoc.org/emersyx.net/emersyx_apis
[emersyx-repo]: https://github.com/emersyx/emersyx
[1]: https://www.virustotal.com/en/documentation/public-api/
[2]: https://developers.virustotal.com/v2.0/reference
[3]: https://www.reddit.com/dev/api/
[4]: https://developers.google.com/gmail/api/
[5]: https://developers.google.com/gmail/api/quickstart/go
[6]: https://api.slack.com/bot-users
[7]: https://discordapp.com/developers
[8]: https://developer.github.com/v3/
