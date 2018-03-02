package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var ifile string

func init() {

	flag.StringVar(&ifile, "f", "", "-f=app.log ")
}
func main() {

	flag.Parse()

	ind := 0
	inds := make([]int, 0)
	var reader io.Reader

	if ifile == "" {
		reader = os.Stdin
	} else {
		f, err := os.Open(ifile)
		if err != nil {
			return
		}
		reader = f
	}
	fmt.Println(ifile)
	fargs := flag.Args()
	if len(fargs) == 0 {
		inds = append(inds, 0)
	} else {
		for i := range fargs {
			k, err := strconv.Atoi(fargs[i])
			if err != nil {
				return
			}
			inds = append(inds, k)
		}
	}
	fmt.Println(flag.Args())
	spider := NewSpider(inds)
	inputReader := bufio.NewReader(reader)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			break
		}
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue
		}
		os.Stdout.WriteString(input)

		item, err := spider.ParseInd(input)
		if err != nil {

			break
		}
		if len(item) == 0 {
			continue
		}
		i := ind
		if len(item) < i {
			i = len(item)
		}
		fmt.Println(strings.Join(item, " "))

	}

}
