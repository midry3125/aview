package main

import (
    "fmt"
    "io"
    "os"
    "time"
    "net/http"
    "encoding/json"
    "aview/tui"
)

func GetSeason(month time.Month) int {
    var season int
    switch int(month) {
        case 1, 2, 3:
            season = 1 
        case 4, 5, 6:
            season = 2
        case 7, 8, 9:
            season = 3
        case 10, 11, 12:
            season = 4
    }
    return season
}

func GetInfo() []tui.Information {
    var info []tui.Information
    datetime := time.Now()
    var api_url string = fmt.Sprintf("https://api.moemoe.tokyo/anime/v1/master/%d/%d", datetime.Year(), GetSeason(datetime.Month()))
    fmt.Printf("Getting information from '%s'", api_url)
    res, err := http.Get(api_url)
    if err != nil {
        fmt.Printf("couldn't access '%s'\n", api_url)
        os.Exit(1)
    }
    defer res.Body.Close()
    byte_data, _ := io.ReadAll(res.Body)
    err = json.Unmarshal(byte_data, &info)
    if err != nil {
        fmt.Println(string(byte_data))
        fmt.Println(err)
        os.Exit(1)
    }
    return info
}

func main() {
    response := GetInfo()
    tui.Run(response)
}