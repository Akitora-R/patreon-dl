package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"testing"
)

func TestChromeDP(t *testing.T) {
	o := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer("http://localhost:1082"),
	)
	cx, cancel := chromedp.NewExecAllocator(context.Background(), o...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(cx)
	defer cancel()
	defer cancel()
	chromedp.Run(ctx,
		chromedp.Navigate("https://www.patreon.com/home"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}
			for _, cookie := range cookies {
				if strings.Contains(cookie.Domain, "patreon") {
					log.Println(cookie.Name, cookie.Value)
				}
			}
			return nil
		}),
	)
}
