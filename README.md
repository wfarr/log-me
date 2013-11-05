# log-me

Logs are cool. Quick access to structured, key-value logs is just a little bit cooler.

## Usage

```
log-me app=github status=3*

log-me app=gist current_user=wfarr --tail
```

## Installing

Cross-compiled builds for Darwin and Linux amd64 should be available on the [Releases](https://github.com/wfarr/log-me/releases) page.

If you're on another platform, your best bet is to clone this repo, `go get`, and `go build` for your desired platforms.
