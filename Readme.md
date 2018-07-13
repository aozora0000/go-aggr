# go-aggr
Download Flv Stream From [AGGR](http://agqr.jp/)

## How to use
```main.go
package main

import (
    "github.com/aozora0000/go-agqr"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

func main() {
    path, _ := filepath.Abs("temp.flv")
    client, err := agqr.New(path)
    if err != nil {
        log.Fatalln(err.Error())
        os.Exit(1)
    }
    defer client.Close()
    
    callback := func(path string) error {
        // Downloaded Callback From FilePath
        return nil
    }
    
    if err = client.Start(time.Minute*30, callback); err != nil {
        log.Fatalln(err.Error())
        os.Exit(1)
    }
}
```