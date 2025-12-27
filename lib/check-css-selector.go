package checkcssselector

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type selectorOpts struct {
	Url      string `short:"U" long:"url" description:"URL to check" required:"true"`
	Selector string `short:"S" long:"selector" description:"CSS selector to find in the DOM" required:"true"`
}

func Do() {
	opts := &selectorOpts{}
	_, err := flags.ParseArgs(opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	ckr := opts.run()
	ckr.Name = "CSS Selector"
	ckr.Exit()
}

func (opts *selectorOpts) run() *checkers.Checker {
	client := &http.Client{
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return fmt.Errorf("stopped after %d redirects", len(via))
			}
			return nil
		},
	}

	req, err := http.NewRequest("GET", opts.Url, nil)
	if err != nil {
		return checkers.Unknown(fmt.Sprint(err))
	}

	req.Header.Set("User-Agent", "go-check-css-selector/1.0")

	res, err := client.Do(req)
	if err != nil {
		return checkers.Unknown(fmt.Sprint(err))
	}
	if res.StatusCode >= 400 {
		return checkers.Critical(fmt.Sprintf("HTTP %d: %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return checkers.Unknown(fmt.Sprint(err))
	}

	checkSt := checkers.OK
	var msg string
	sel := opts.Selector
	if doc.Find(sel).Length() == 0 {
		msg = fmt.Sprintf("`%s` is not found in `%s`", sel, opts.Url)
		checkSt = checkers.CRITICAL
	} else {
		msg = fmt.Sprintf("`%s` is found `%s`", sel, opts.Url)
	}

	return checkers.NewChecker(checkSt, msg)
}
