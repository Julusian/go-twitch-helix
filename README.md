# go-twitch-helix

[![travis-ci status](https://api.travis-ci.org/Julusian/go-twitch-helix.png)](https://travis-ci.org/Julusian/go-twitch-helix)
[![Coverage Status](https://coveralls.io/repos/github/Julusian/go-twitch-helix/badge.svg?branch=master)](https://coveralls.io/github/Julusian/go-twitch-helix?branch=master)
[![GoDoc](https://godoc.org/github.com/Julusian/go-twitch-helix?status.svg)](https://godoc.org/github.com/Julusian/go-twitch-helix)

Go library for accessing the new [Twitch-API](https://dev.twitch.tv/docs/).

Additional api calls are available for some Twitch v5 Api, to fill in the gaps until the new Api is complete.

## Usage

To install `go-twitch-helix` run the command:

```bash
$ go get github.com/julusian/go-twitch-helix
```

Full docs at [GoDocs](https://godoc.org/github.com/julusian/go-twitch-helix)

Here's an example program that gets the top 10 twitch games:

```go
package main

import (
	"fmt"
	"github.com/julusian/go-twitch-helix/twitch"
	"github.com/julusian/go-twitch-helix/twitchapi"
	"log"
	"net/http"
)

func main() {
	// client := twitchapi.NewClient(&http.Client{}, os.Getenv("CLIENT_ID"))
	// opt := &twitch.ListOptions{
	// 	Limit:  10,
	// 	Offset: 0,
	// }

	// games, err := client.Games.Top(opt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i, s := range games.Top {
	// 	fmt.Printf("%d - %s (%d)\n", i+1, s.Game.Name, s.Viewers)
	// }
}
```

### Authentication

**TODO**

## License

All files under this repository fall under the MIT License (see the file LICENSE).
