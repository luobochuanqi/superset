package main

import (
	"github.com/luobochuanqi/superset/cmd"
	"log"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v *URLValue) String() string {
	if v.URL == nil {
		return ""
	}
	return v.URL.String()
}

func (v *URLValue) Set(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	v.URL = u
	return nil
}

func main() {
	err := cmd.Excute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
