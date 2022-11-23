package cache

import (
    "fmt"
    "os"
    "strings"
    "io/ioutil"
    "encoding/json"
    "path/filepath"
    "aview/tui"
)

var rootdir string

func GetCache(date string) ([]tui.Information, error) {
    var info []tui.Information
    fmt.Printf("Searching cache:  %s...\n", date)
    path := filepath.Join(rootdir, strings.Replace(date, "/", "-", 1))
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return info, err
    }
    err = json.Unmarshal(data, &info)
    return info, nil
}

func SaveCache(contents []byte, date string) error {
    path := filepath.Join(rootdir, strings.Replace(date, "/", "-", 1))
    return ioutil.WriteFile(path, contents, 0666)
}

func init() {
    rootdir = filepath.Join(os.Getenv("APPDATA"), "aview")
    if rootdir == "" {
        rootdir = filepath.Join(os.Getenv("HOME"), "aview")
    }
    _, err := os.Stat(rootdir)
    if os.IsNotExist(err) {
        os.Mkdir(rootdir, 0666)
    }
}