mastodon :
	go build -o mastodon

mac :
	go build -o Mastodon.app/Contents/MacOS/Mastodon

clean :
	rm -f mastodon
	rm -f Mastodon.app/Contents/MacOS/Mastodon

.PHONY: mac clean
