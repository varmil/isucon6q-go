## rsync *.go files and Makefile (in HOST)
```sh
  config.vm.synced_folder "./", "/home/isucon/webapp/go/", type: "rsync",
    owner: "isucon",
    group: "isucon",
    rsync__args: ["-a", "--include=*.go", "--include=Makefile", "--exclude=*"]
```


## watch .go files, and execute make command (in GUEST)
```sh
sudo apt install -y inotify-tools
inotifywait -e modify -mr /home/isucon/webapp/go | while read;do while read -t 0.5;do :;done;make -C /home/isucon/webapp/go ;done
```


## pprof (in GUEST)
```sh
# app.goのimportに追加
    "net/http"
    _ "net/http/pprof"
    
# app.goのmain()の先頭に記述
    go func() {
        log.Println(http.ListenAndServe(":6060", nil))
    }()
```

```sh
go get -u github.com/google/pprof
sudo apt install -y graphviz
pprof -http="0.0.0.0:8888" localhost:6060/debug/pprof/profile

# stop firewalld if needed
sudo systemctl stop firewalld
sudo systemctl disable firewalld
```


## logging (GUEST)
```sh
sudo journalctl -u isuda.go.service -ef
```


## dstat (GUEST)
[勝手のいいdstatコマンドオプション](https://blog.masu-mi.me/post/2015/02/28/dstat_options/)
