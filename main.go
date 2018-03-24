package main

import "github.com/zserge/webview"

func main() {
	webview.Open("Mastodon",
		// Change this to your Mastodon instance's URL.
		"https://bsd.network", 380, 750, true)
}
