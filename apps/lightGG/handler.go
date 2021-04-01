package lightGG

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Spider struct {
	url    string
	header map[string]string
}

func (r *Spider) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		log.Infof(fmt.Sprintf("检查网页错误%+v", err.Error()))
	}
	for key, value := range r.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Infof(fmt.Sprintf("检查网页错误%+v", err.Error()))
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Infof(fmt.Sprintf("检查网页错误%+v", err.Error()))
	}
	return string(body)
}

/*
https://www.light.gg/ 网站处理
*/
// LightGGChecker 检查lightgg 网站链接是否正确
func LightGGChecker(url string) bool {
	log.Infof(fmt.Sprintf("正在检查light gg网页[%s]", url))
	header := map[string]string{
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	}
	spider := Spider{url, header}
	html := spider.get_html_header()
	// 404标题
	pattern2 := `<h2>(.*?)</h2>`
	rp2 := regexp.MustCompile(pattern2)
	find_txt2 := rp2.FindAllStringSubmatch(html, -1)
	for i := 0; i < len(find_txt2); i++ {
		if strings.Contains(find_txt2[i][1], "404") {
			log.Infof(fmt.Sprintf("[%s] 错误", url))
			return false
		}
	}
	log.Infof(fmt.Sprintf("[%s] 正确", url))
	return true
}

// UrlShotCutHandler 网页快照截图
func UrlShotCutHandler(url, filename string) {
	// Start Chrome
	// Remove the 2nd param if you don't need debug information logged
	// ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run Tasks
	// List of actions to run in sequence (which also fills our image buffer)
	var imageBuf []byte
	if err := chromedp.Run(ctx, fullScreenshot(url, 100, &imageBuf)); err != nil {
		log.Infof(err.Error())
	}

	// Write our image to file
	if err := ioutil.WriteFile(filename, imageBuf, 0644); err != nil {
		log.Infof(err.Error())
	}
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// 全屏截图
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	// log.Infof("截图等待2秒，防止先")
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, cssContentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(cssContentSize.Width)), int64(math.Ceil(cssContentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      cssContentSize.X,
					Y:      cssContentSize.Y,
					Width:  cssContentSize.Width,
					Height: cssContentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
