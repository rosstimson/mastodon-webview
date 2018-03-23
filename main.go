package main

import "github.com/zserge/webview"

func main() {
	webview.Open("Mastodon",
		"https://bsd.network", 380, 750, true)
}
