package goscrape

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

func initWithJavascript(response *colly.Response) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.NoFirstRun,
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()
	}

	var res string
	if err := chromedp.Run(ctx, chromedp.Navigate(response.Request.URL.String()),
		chromedp.InnerHTML("html", &res), // Scrape whole rendered page
	); err != nil {
		return err
	}

	response.Body = []byte(res)

	return nil
}
