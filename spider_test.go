package main

import (
	"bufio"

	"io"
	"os"
	"strings"
	"testing"
)

/*
func TestSimple(t *testing.T) {

	strs := `120.52.37.80 "21/Nov/2017:00:01:32 +0800" "aaaa" bbbB`

	item, err := Spider(strs)
	if err != nil {
		t.Fatal(err)
	}
	if len(item) != 4 {
		t.Fatal("Spider error")
	}
	if item[3] != "bbbB" {
		t.Fatal("Spider error")
	}

}

func TestBrucket(t *testing.T) {
	strs := `120.52.37.80 [21/Nov/2017:00:01:32 +0800] "aaaa" bbbB`
	item, err := Spider(strs)
	if err != nil {
		t.Fatal(err)
	}
	if len(item) != 4 {
		t.Fatal("Spider error")
	}
	if item[3] != "bbbB" {
		t.Fatal("Spider error")
	}

}

func TestCloudlog(t *testing.T) {
	strs :=
		`120.52.37.80 [23/Nov/2017:11:17:49 +0800] "GET http://y.play.360kan.com/vod/819/350/9a7c53fc0736749d686e83d089ede2e6/050ed6695d03e3b61dd94f024ca11d0c_960_480_600_128_9a7c53fc0736749d686e83d089ede2e6_mp4_3g8x30800_199.ts HTTP/1.1" 200 679111 CLOUD_CACHE_COMBINE_MISS "http://cdn.inf.360kan.com/swf/ShortVideo_v3.08.swf" "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36 QIHU 360SE" [ 0.000 0.000 0.000 16.499 0.000 189.811 23.247 206.310 23.247 229.557] [123.125.52.108,120.52.32.97_-] [Local] [0_0|1_0|diff] [/memdisk_cache1]`
	t.Log(strs)

	item, err := Spider(strs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(item))
	t.Log(item)

	if len(item) != 13 || item[12] != "/memdisk_cache1" {
		t.Fatal("Spider error")
	}

}
func TestWebcdnlog(t *testing.T) {
	strs :=
		`m.88dushu.com 106.39.191.113 - - [23/Nov/2017:11:31:38 +0800] "GET http://m.88dushu.com/mobile/css/jquery.min.js HTTP/1.1" "304" 343 w-f04.bjdt 0.000 UID_21523 "http://m.88dushu.com/mulu/85876/" "Mozilla/5.0 (Linux; U; Android 6.0.1; zh-CN; BIHEE A8 Build/MMB29M) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 UCBrowser/11.1.6.899 U3/0.8.0 Mobile Safari/534.30" "-" "Error from -" "cdn_edge" HIT`
	t.Log(strs)

	item, err := Spider(strs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(item))
	t.Log(item)

	if len(item) != 17 {

		t.Fatal("Spider error")
	}
	if item[16] != "HIT" {
		t.Log(len(item[16]))
		t.Fatal("Spider error")
	}
}
*/
func ReadLine(fileName string, handler func(string, *testing.T), t *testing.T) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line, t)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func Print(line string, t *testing.T) {
	t.Log(line)
}
func VaildSpider(line string, t *testing.T) {
	item, err := Spider(line)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(item))
	t.Log(item)

	if len(item) != 13 {

		t.Fatal("Spider error")
	}

	if item[12] != "/memdisk_cache1" {
		t.Log(len(item[12]))
		t.Fatal("Spider error")
	}
	t.Log(item[2])

}
func TestLogfile(t *testing.T) {
	ReadLine("one.log", VaildSpider, t)

}
