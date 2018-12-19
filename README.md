
## watch .go files, and execute make command (in GUEST)
```sh
sudo apt install -y inotify-tools
inotifywait -e modify -mr /home/isucon/webapp/go | while read;do while read -t 0.5;do :;done;make -C /home/isucon/webapp/go ;done
```


## rsync *.go files and Makefile (in HOST)
```sh
  config.vm.synced_folder "./", "/home/isucon/webapp/go/", type: "rsync",
    owner: "isucon",
    group: "isucon",
    rsync__args: ["-a", "--include=*.go", "--include=Makefile", "--exclude=*"]
```


## pprof (in GUEST)
```sh
go get -u github.com/google/pprof
sudo apt install -y graphviz
pprof -http="0.0.0.0:8888" localhost:6060/debug/pprof/profile
```
