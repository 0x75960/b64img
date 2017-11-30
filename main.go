package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var template = `<image src="data:%s;base64,%s"/>` + "\n"
var srcType = flag.String("t", "image/png", "type to embed")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			"b64img:\n\tgenerate base64-embeded image tag for markdown/html documents\n\nUsage:\n\tb64img [options] /path/to/target/image\n\nOptions:\n\n",
		)

		flag.PrintDefaults()
	}
}

func base64encode(p string) (encoded string, err error) {

	f, err := os.Open(p)
	if err != nil {
		return encoded, err
	}

	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return encoded, err
	}

	encoded = base64.StdEncoding.EncodeToString(d)

	return
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalln("something wrong. please see --help")
	}

	e, err := base64encode(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(template, *srcType, e)
}
