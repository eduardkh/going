package gethttpstat

// https://github.com/davecheney/httpstat - alternative

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/eduardkh/go-httpstat"
)

func Gethttpstat(args string) httpstat.Result {
	req, err := http.NewRequest("GET", args, nil)
	if err != nil {
		log.Fatal(err)
	}

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	result.End(time.Now())

	return result
}
