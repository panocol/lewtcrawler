package main

import "fmt"
import "net/http"

import "github.com/PuerkitoBio/fetchbot"

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(handler))

	queue := f.Start()
	queue.SendStringHead("https://ece.osu.edu/news/2015/10/texas-inst.-scholars-program")
	queue.Close()
}

func handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())
}