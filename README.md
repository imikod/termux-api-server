Termux API Server
=================
Server for [termux-api](https://github.com/termux/termux-api)

Usage
=====
```
$ termux-api-server
```
Server starts with default port `8000`
You can change the port to `8991` like so:
```
$ termux-api-server -p 8991
```
and open [http://localhost:8991](http://localhost:8991)

Build from source
=================
- Install go
```
$ apt install golang
```
- Set GOPATH
```
$ cd ~
$ mkdir go go/bin go/pkg go/src
$ export GOPATH=~/go
```
- Install termux-api-server
```
$ go get github.com/imikod/termux-api-server
```

Security Warning
================

Make sure your device has no data connection and it is only connected to your own secure wlan.
Otherwise sensitive information could be exposed to the internet.
