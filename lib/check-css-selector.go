package checkcssselector

import (
	"fmt"
	"net/http"
	"os"

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
	res, err := http.Get(opts.Url)
	if err != nil {
		return checkers.Unknown(fmt.Sprint(err))
	}
	// TODO: check status code
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
