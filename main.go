package main

import (
	"github.com/zserge/webview"
)

func main() {
	url := webServer()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Internet Checker")
	w.SetSize(400, 600, webview.HintFixed)
	w.Navigate(url)
	println("Launching window...")
	w.Run()

}
