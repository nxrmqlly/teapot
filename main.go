package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const port = "4180"

const msg = `hii,
I want to run this ysws called Teapot (#teapot)
[ name comes from HTTP/1.1 418 I'm a teapot; see the headers! ]

Basically:

- YS: ANYTHING that speaks HTTP: an api, a graphql server, api clients, reverse proxy, load balancer, tunnels
- WS: Stickers, Domain Grants, AI Grants, Raspberry Pis, .... TEA! .... and more...
- projects must implement HTTP meaningfully and contain 418 I'm a teapot somewhere.

If you like it, consider joining #teapot <3 thanks!

    -- @riiitam
`

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) Inc() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	return c.count
}

func main() {
	c := counter{count: 0}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("teapot'd: %d times\n", c.Inc())

		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(w, msg)
	})

	fmt.Printf("Listening on port %s!\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
