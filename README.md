# go-twitch-helix

[![travis-ci status](https://api.travis-ci.org/Julusian/go-twitch-helix.png)](https://travis-ci.org/Julusian/go-twitch-helix)
[![Coverage Status](https://coveralls.io/repos/github/Julusian/go-twitch-helix/badge.svg?branch=master)](https://coveralls.io/github/Julusian/go-twitch-helix?branch=master)
[![GoDoc](https://godoc.org/github.com/Julusian/go-twitch-helix?status.svg)](https://godoc.org/github.com/Julusian/go-twitch-helix)

Go library for accessing the new [Twitch-API](https://dev.twitch.tv/docs/).

Additional api calls are available for some Twitch v5 Api, to fill in the gaps until the new Api is complete.

## Usage

To install `go-twitch-helix` run the command:

```bash
$ go get github.com/Julusian/go-twitch-helix
```

Full docs at [GoDocs](https://godoc.org/github.com/Julusian/go-twitch-helix)

Here's an example program that gets the top 10 twitch games:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Julusian/go-twitch-helix/helix"
	"github.com/Julusian/go-twitch-helix/twitch"
)

func main() {
	client := twitch.NewApiClient(&http.Client{}, os.Getenv("CLIENT_ID"))

	opts := &helix.StreamsParams{
		Limit: 10,
	}

	res, rate, err := helix.GetStreams(client, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rate limiter %d/%d remaining", rate.Remaining, rate.Limit)

	for i, s := range res.Data {
		fmt.Printf("%d - %s (%d)\n", i+1, s.Title, s.ID)
	}
}
```

### Authentication

**TODO**

## License

All files under this repository fall under the MIT License (see the file LICENSE).
