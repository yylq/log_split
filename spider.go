package main

import (
	"errors"
	"fmt"
)

type Spider struct {
	ind []int
}

func NewSpider(inds []int) *Spider {
	return &Spider{ind: inds}
}

const (
	st_start          = 0
	st_brackets_going = 1
	st_quotes_going   = 2
	st_end            = 3
)

func (this Spider) ParseInd(s string) ([]string, error) {
	its, err := this.ParseItem(s)
	if err != nil {
		return nil, err
	}
	fmt.Println(this.ind)
	if len(its) < len(this.ind) {
		return nil, errors.New("no ind item")
	}
	item := make([]string, len(this.ind))
	for i := range this.ind {
		if this.ind[i] > len(its) {
			return nil, errors.New("overload max item")
		}
		item[i] = its[this.ind[i]]
	}
	return item, nil
}
func (this Spider) ParseItem(s string) ([]string, error) {
	if len(s) == 0 {
		return make([]string, 0), nil
	}
	state := st_start
	buf := []byte(s)
	item := make([]string, 0)
	last := 0
	for i := 0; i < len(buf); {
		switch state {
		case st_start:
			switch buf[i] {
			case ' ':
				if last != i {
					item = append(item, string(buf[last:i]))
				}
				i++
				last = i
			case '"':
				if last != i {
					item = append(item, string(buf[last:i]))
				}

				state = st_quotes_going
				i++
				last = i
			case '[':
				if last != i {
					item = append(item, string(buf[last:i]))
				}

				state = st_brackets_going
				i++
				last = i
			case '\r':
			case '\n':
				if last != i {
					item = append(item, string(buf[last:i]))
				}
				state = st_end
			default:
				//fmt.Print(i, ",", len(buf), "\n")
				if i == len(buf)-1 {
					fmt.Print(i, ",", string(buf[i]), ",", len(buf), "\n")
					item = append(item, string(buf[last:i+1]))

					state = st_end
					last = i

				}
				i++

			}
		case st_brackets_going:
			switch buf[i] {
			case ']':
				item = append(item, string(buf[last:i]))
				last = i + 1
				state = st_start
				i++
			default:
				i++
			}
		case st_quotes_going:
			switch buf[i] {
			case '"':
				item = append(item, string(buf[last:i]))
				last = i + 1
				state = st_start
				i++
			default:
				i++

			}
		case st_end:

			break
		default:
			return nil, errors.New("spider invalid state error")

		}
	}

	return item, nil
}
