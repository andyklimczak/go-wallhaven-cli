# Go wallhaven CLI

---

Simple and unofficial CLI for interacting with the [wallhaven](https://wallhaven.cc) API. Used primarily to download wallpapers from collections.

Stop manually saving wallpapers and just favorite wallpapers on wallhaven instead.

## Install

---

Install at least Golang 1.23.

Then build the `go-wallhaven` executable:

```shell
git clone
go build
sudo cp go-wallhaven /usr/local/bin
```

## Usage

---

### Download

```shell
Usage:
  go-wallhaven download [flags]

Flags:
  -h, --help          help for download
  -t, --threads int   number of threads (default 4)

Global Flags:
  -a, --apikey string             destination directory
  -c, --collection-label string   collection label
  -d, --destination string        destination directory (default "~/.wallpapers")
  -u, --username string           username of the wallhaven user who owns the collection
  -v, --verbose                   verbose output
```

### Example

Download from `testuser`'s collection named `Desktop` into `~/.my_wallpapers`:
```shell
go-wallhaven download -d "~/.my_wallpapers" -a "APIKEY" -c "Desktop" -u "testuser"
```

Due to the structure of the wallhaven API, both the `apikey` and `username` are required parameters.

### Usage as a Service

#### Systemd

Systemd can be used to automatically download wallpapers as you add them to your collection on Wallhaven.

Create a new systemd service at `/etc/systemd/system/go-wallhaven.service` and paste this template and replace `USER` with your user and add parameters to the `go-wallhaven` command:

Be sure to create the directory for `WorkingDirectory` before starting the service.

```shell
[Unit]
Description=Go-wallhaven
Documentation=
After=network.target network-online.target
Requires=network-online.target

[Service]
Type=simple
User=USER
Group=USER
WorkingDirectory=/home/USER/.wallpaper
ExecStart=go-wallhaven download -d "/home/USER/.my_wallpapers" -a "APIKEY" -c "Desktop" -u "USERNAME" -v
ExecReload=go-wallhaven download -d "/home/USER/.my_wallpapers" -a "APIKEY" -c "Desktop" -u "USERNAME" -v
TimeoutStopSec=5s
LimitNOFILE=1048576
PrivateTmp=true
ProtectSystem=full
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target

```

This will run `go-wallhaven` on boot and download the wallpapers in `testuser`'s `Desktop` collection into your `~/.my_wallpapers` folder.

Start and enable the service with:
```shell
sudo systemctl start go-wallhaven
sudo systemctl enable go-wallhaven
```

Then point the program that loads your wallpapers to the folder `~/.my_wallpapers`:
```shell
feh --randomize --bg-fill ~/.my_wallpapers/*
```
