package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func main() {
	w := &webPage{url: "http://www.google.com/"}
	fmt.Printf("%p\n", w)
	w.get()
	fmt.Printf("URL: %s Error: %s Length: %d\n", w.url, w.err, len(w.body))

}
