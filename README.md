# mastodon-webview

Tiny little app that presents the Masodon web app in its own window
using [webview](https://github.com/zserge/webview).

It is essentially just the simple example from the webview repo with the
window size adjusted, my Mastodon instance of choice set as the URL, and
some additional cruft needed for use as a Mac app.

## Preview

![Screenshot](screenshot.png)

## Build

You'll need a working Go development environment setup as well
gtk-webkit2 (OS package names will vary) , on a Mac an install of Xcode
should have you covered.

Install [webview](https://github.com/zserge/webview) with:

    $ go get github.com/zserge/webview

Clone this repo:

    $ mkdir -p $GOPATH/src/github.com/rosstimson

    $ git clone https://github.com/rosstimson/mastodon-webview.git $GOPATH/src/github.com/rosstimson/mastodon-webview

Tweak [this line](https://github.com/rosstimson/mastodon-webview/blob/master/main.go#L7)
of code for your Mastodon instance of choice.

Build with

    $ make

Or if on a mac use:

    $ make mac

Copy the built binary into `/usr/local/bin` or if on a Mac copy the entire
Mastodon.app directory into `/Applications`.

## Credit

The icon I've used for the Mac icon is from the [Masodon Press Kit](https://github.com/tootsuite/joinmastodon/blob/master/public/press-kit.zip),
there are a couple alternatives there you might prefer.

## License

Licensed under the [3-Clause BSD License](https://opensource.org/licenses/BSD-3-Clause).
