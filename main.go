package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

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
	m := http.NewServeMux()
	m.HandleFunc("/", handleRequest)

	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
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

	_, err := w.Write([]byte(sb.String()))
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
	}
}
