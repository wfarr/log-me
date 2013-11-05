# log-me

Logs are cool. Quick access to structured, key-value logs is just a little bit cooler.

## Usage

```
log-me app=github status=3*

log-me app=gist current_user=wfarr --tail
```

## Installing

Right now I've not cross-compiled any of the binaries, so your best bet is to clone this repo, `go get`, and `go build` for your desired platforms.
