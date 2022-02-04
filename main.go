package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

import "github.com/kelseyhightower/envconfig"

func writeRow(sb *strings.Builder, key string, values ...string) {
	sb.WriteString("<tr>")
	sb.WriteString("<td>")
	sb.WriteString(key)
	sb.WriteString("</td>")

	sb.WriteString("<td>")
	for _, v := range values {
		sb.WriteString(v)
		sb.WriteString("<br />")
	}
	sb.WriteString("<td>")
	sb.WriteString("</tr>")
}

func main() {
	config := &Config{}
	envconfig.MustProcess("", config)

	a := &Action{config: config}

	m := http.NewServeMux()
	m.HandleFunc("/", a.handleRequest)

	s := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: m,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
	}
}

type Action struct {
	config *Config
}

func (a *Action) handleRequest(w http.ResponseWriter, r *http.Request) {
	sb := &strings.Builder{}

	p := r.URL.Path
	sb.WriteString("<h1>")
	sb.WriteString(p)
	sb.WriteString("</h1>")

	sb.WriteString("<table>")
	writeRow(sb, "Host", r.Host)
	writeRow(sb, "Method", r.Method)
	writeRow(sb, "Proto", r.Proto)
	writeRow(sb, "RemoteAddr", r.RemoteAddr)
	writeRow(sb, "URL", r.URL.String())
	for k, vlist := range r.Header {
		writeRow(sb, k, vlist...)
	}
	sb.WriteString("</table>")

	if a.config.OidcJwtService != "" {
		oidcdata := r.Header.Get("X-Amzn-Oidc-Data")
		if oidcdata != "" {
			resp, err := http.Post(a.config.OidcJwtService, "text/plain", strings.NewReader(oidcdata))
			sb.WriteString("<pre>")
			if err != nil {
				sb.WriteString(err.Error())
			} else {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					sb.WriteString(err.Error())
				} else {
					defer resp.Body.Close()
					sb.Write(body)
				}
			}
			sb.WriteString("</pre>")
		}
	}
	_, err := w.Write([]byte(sb.String()))
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
	}
}
