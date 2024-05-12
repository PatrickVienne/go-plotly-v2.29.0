package grob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

// FigOptions enables users to pass their custom cdn or local plotly file reference to allow building offline solutions
type FigOptions struct {
	HeadContent string
}

type Options struct {
	FigOptions
	Addr string
}

// ToWriter saves the figure as standalone HTML. It still requires internet to load plotly.js from CDN.
func ToWriter(fig *Fig, w io.Writer, options ...FigOptions) {
	_, err := w.Write(figToBuffer(fig, options...).Bytes())
	if err != nil {
		panic(errors.New(fmt.Sprintf("failed to write figure to passed writer: %v", err)))
	}
}

// ToHtml saves the figure as standalone HTML. It still requires internet to load plotly.js from CDN.
func ToHtml(fig *Fig, path string, options ...FigOptions) {
	buf := figToBuffer(fig, options...)
	os.WriteFile(path, buf.Bytes(), os.ModePerm)
}

// Show displays the figure in your browser.
// Use serve if you want a persistent view
func Show(fig *Fig, options ...FigOptions) {
	buf := figToBuffer(fig, options...)
	browser.OpenReader(buf)
}

const cdnUrl = `<script src="https://cdn.plot.ly/plotly-2.19.0.min.js"></script>`

// computeFigOptions allows to default to the common cdn reference
func computeFigOptions(options ...FigOptions) string {
	if len(options) > 1 {
		panic(errors.New(fmt.Sprintf("only use 1 option at a time. Passed: %d. options: %+v", len(options), options)))
	}
	if len(options) == 0 {
		return cdnUrl
	}
	return options[0].HeadContent
}

// figToBuffer with optional parameter options, only the first argument is evaluated
func figToBuffer(fig *Fig, options ...FigOptions) *bytes.Buffer {
	var headContent string = computeFigOptions(options...)

	figBytes, err := json.Marshal(fig)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("plotly").Parse(getBaseHtml(headContent))
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, string(figBytes))
	return buf
}

// Serve creates a local web server that displays the image using plotly.js
// Is a good alternative to Show to avoid creating tmp files.
func Serve(fig *Fig, opt ...Options) {
	opts := computeOptions(Options{
		FigOptions: FigOptions{HeadContent: cdnUrl},
		Addr:       "localhost:8080",
	}, opt...)

	mux := &http.ServeMux{}
	srv := &http.Server{
		Handler: mux,
		Addr:    opts.Addr,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf := figToBuffer(fig, opts.FigOptions)
		buf.WriteTo(w)
	})

	log.Print("Starting server")
	if err := srv.ListenAndServe(); err != nil {
		log.Print(err)
	}
	log.Print("Stop server")
}

func computeOptions(def Options, opt ...Options) Options {
	if len(opt) == 1 {
		opts := opt[0]
		if opts.Addr != "" {
			def.Addr = opts.Addr
			def.HeadContent = opts.HeadContent
		}
	}
	return def
}

func getBaseHtml(head string) string {
	return fmt.Sprintf(`
<head>
	%s
</head>
<body>
	<div id="plot"></div>
	<script>
		data = JSON.parse('{{ . }}')
		Plotly.newPlot('plot', data);
	</script>
</body>
`, head)
}
