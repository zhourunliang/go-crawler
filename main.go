package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "time"
    "os"
    // "strconv"
    "github.com/PuerkitoBio/goquery"
    "log"
)
//定义新的数据类型
type Spider struct {
    url    string
    header map[string]string
}


//定义 Spider get的方法
func (task Spider) get_html_header() string {
    client := &http.Client{}
    req, err := http.NewRequest("GET", task.url, nil)
    if err != nil {
    }
    for key, value := range task.header {
        req.Header.Add(key, value)
    }
    resp, err := client.Do(req)
    if err != nil {
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    }
    return string(body)

}
func parse()  {
    header := map[string]string{
        "Host": "movie.douban.com",
        "Connection": "keep-alive",
        "Cache-Control": "max-age=0",
        "Upgrade-Insecure-Requests": "1",
        "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
        "Referer": "http://zhourunliang.github.io/",
    }

    f, err := os.Create(".\\spider.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    //写入标题
    f.WriteString("博客"+"\r\n")

    url:="http://zhourunliang.github.io/"
    fmt.Println("url= ", url)
    spider := &Spider{url, header}
    html := spider.get_html_header()
    // fmt.Println("html= ", html)

    //标题
    pattern1:=`<a class="article-title" href="(.*?)">(.*?)</a>`
    rp1 := regexp.MustCompile(pattern1)
    find_txt1 := rp1.FindAllStringSubmatch(html,-1)

    //正文
    pattern2:=`<a class="article-title" href="(.*?)">(.*?)</a>`
    rp2 := regexp.MustCompile(pattern2)
    find_txt2 := rp2.FindAllStringSubmatch(html,-1)


    for i:=0;i<len(find_txt1);i++{
        fmt.Printf("%s\n",find_txt1[i][2] )
        f.WriteString(find_txt1[i][2]+"\r\n")
        f.WriteString(find_txt2[i][2]+"\r\n")
    }
}

func goqueryParse()  {
    doc, err := goquery.NewDocument("http://zhourunliang.github.io")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(doc.Html())            //.Html()得到html内容
    title := doc.Find(".article-title").Text()
}


func main() {
    t1 := time.Now() // get current time
    // parse()
    goqueryParse()
    elapsed := time.Since(t1)
    fmt.Println("爬虫结束,总共耗时: ", elapsed)

}
