package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
	"log"
)

func main() {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var b1, b2 []byte
	if err := chromedp.Run(ctx,
		// emulate iPhone 7 landscape
		chromedp.Emulate(device2.IPhone7landscape),
		chromedp.Navigate(`https://www.ricardo.ch/de/a/2-tolle-reisepiele-1160121108/`),
		chromedp.CaptureScreenshot(&b1),

		// reset
		chromedp.Emulate(device.Reset),

		// set really large viewport
		chromedp.Emulate(device.Pixel2XL),
		chromedp.Navigate(`https://www.ricardo.ch/de/a/2-tolle-reisepiele-1160121108/`),
		chromedp.CaptureScreenshot(&b2),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot1.png", b1, 0o644); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("screenshot2.png", b2, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote screenshot1.png and screenshot2.png")
}
