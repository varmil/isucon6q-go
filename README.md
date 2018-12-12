
## watch .go files, and execute make command (in GUEST)
```sh
inotifywait -e create,delete,modify,move -mr /home/isucon/webapp/go | while read;do while read -t 0.5;do :;done;make -C /home/isucon/webapp/go ;done
```

## rsync *.go files (in HOST)
```sh
  config.vm.synced_folder "./", "/home/isucon/webapp/go/", type: "rsync",
    rsync__args: ["-a", "--include=*.go", "--exclude=*"]
```