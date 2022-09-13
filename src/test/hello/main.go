package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/dom"
	_ "github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start")

	fmt.Println("Start Init")
	opts := []chromedp.ExecAllocatorOption{
		chromedp.ExecPath("/usr/bin/chromium-browser"),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("mute-audio", true),
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	//var resa *runtime.RemoteObject
	fmt.Println("Chromium Start 3")

	var res string
	//var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPhone11Pro),
		chromedp.Navigate(`https://www.ricardo.ch/de/a/schreibtisch-1164808140/`),
		//chromedp.EvaluateAsDevTools(`window.ricardo`, &resa, chromedp.EvalAsValue),
		chromedp.ActionFunc(func(ctx context.Context) error {
			//var resa *runtime.RemoteObject
		//	chromedp.Evaluate(`window.ricardo`,resa)
		//	json_byte ,_:= resa.MarshalJSON()
		//
		//	var out bytes.Buffer
		//
		//	_ = json.Indent(&out, json_byte,"","\t")
		//
		//	color.Red(out.String())
			node, err := dom.GetDocument().Do(ctx)
		//	if err != nil {
		//		return err
		//	}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),

	)
	fmt.Println("Chromium End")

	if err != nil {
		fmt.Println(err)
	}

	doc, err := html.Parse(strings.NewReader(res))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	removeScript(doc)
	//fmt.Println("Make Json")

	//jsonByte, _ := json.Marshal(resa)
	//var out bytes.Buffer
	//_ = json.Indent(&out, jsonByte, "", "\t")
	//fmt.Println("Create File Json")
	//f, err := os.Create("data.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//fmt.Println("Write in File")
	//
	//_, err2 := f.WriteString(out.String())
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	fmt.Println("done")

}

func removeScript(n *html.Node) {
	// if note is script tag<
	if n.Type == html.ElementNode && n.Data == "script" {
		var buf bytes.Buffer
		w := io.Writer(&buf)
		html.Render(w, n)
		fmt.Println(buf.String())
	}
	// traverse DOM
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScript(c)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4444", nil))
}
