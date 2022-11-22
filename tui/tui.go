package tui

import (
    "fmt"
    "os"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

type Information struct {
    ProductCompanies string `json:"product_companies"`
    PublicUrl string `json:"public_url"`
    Title string `json"title"`
    ShortTitle string `json:"title_short1"`
    TwitterAccount string `json:"twitter_account"`
    TwitterHashTag string `json:"twitter_hash_tag"`
}

func Run(info []Information) {
    app := createApplication(info)
    err := app.Run()
    if err != nil {
        os.Exit(1)
    }
}

func make_desc(info Information) string {
    return fmt.Sprintf("Title:  %s\nShort Title:  %s\nProduct Company:  %s\nPublic Site:  %s\nTwitter ID:  @%s\nTwitter Hashtag:  #%s", info.Title, info.ShortTitle, info.ProductCompanies, info.PublicUrl, info.TwitterAccount, info.TwitterHashTag)
}

func createApplication(info []Information) *tview.Application {
    app := tview.NewApplication()
    pages := tview.NewPages()
    description := tview.NewTextView()
    description.SetTitle("Information")
    description.SetBorder(true)
    description.SetText(make_desc(info[0]))
    anime_titles := tview.NewList()
    anime_titles.SetMainTextColor(tcell.GetColor("#00FFFF"))
    anime_titles.SetSelectedBackgroundColor(tcell.GetColor("#008000"))
    for _, i := range info {
        anime_titles.AddItem(i.Title, "", 0, func() {})
        anime_titles.SetChangedFunc(func(n int, _ string, _ string, _ rune) {
            description.Clear().SetText(make_desc(info[n]))
        })
    }
    flex := tview.NewFlex()
    flex.AddItem(anime_titles, 0, 1, true)
    flex.AddItem(description, 0, 1, false)
    pages.AddPage("main", flex, true, true)
    app.SetRoot(pages, true)
    return app
}