package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var cookie string = "bid=F3CqafGd14c; douban-fav-remind=1; __utmv=30149280.23292; gr_user_id=9d0f971e-34e6-482a-a0a3-74bccf3f7ca1; _vwo_uuid_v2=D2F6045FA6F8752869DCA358F786C8991|6b411d5b77fcccb5152945fbb15ed40d; ll=\"118282\"; dbcl2=\"232925458:jB8hokGuA04\"; ck=_s3a; __gads=ID=27af2271b3dd81d8-22bd76c6a7c9001b:T=1624278029:RT=1624278029:S=ALNI_MZVBbkQ9YVoilD6rNyHkZ8S8iSs2A; push_doumail_num=0; push_noty_num=0; ct=y; __utma=30149280.1929384996.1619837822.1623307796.1624278030.7; __utmc=30149280; __utmz=30149280.1624278030.7.5.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmt=1; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=b3ef104d-62a8-45c9-b16f-814227bf0f46; gr_cs1_b3ef104d-62a8-45c9-b16f-814227bf0f46=user_id:1; __utmt_douban=1; __utma=81379588.882518171.1621589003.1621589003.1624278122.2; __utmc=81379588; __utmz=81379588.1624278122.2.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/group/nanshanzufang/; _pk_ref.100001.3ac3=[\"\",\"\",1624278122,\"https://www.douban.com/group/nanshanzufang/?ref=sidebar\"]; _pk_ses.100001.3ac3=*; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_b3ef104d-62a8-45c9-b16f-814227bf0f46=true; _vwo_uuid_v2=D2F6045FA6F8752869DCA358F786C8991|6b411d5b77fcccb5152945fbb15ed40d; __utmb=30149280.17.5.1624278118191; __utmb=81379588.3.10.1624278122; _pk_id.100001.3ac3=14e8b2137ec830a1.1621589003.2.1624278151.1621589003."
//var cookie string = "sid=1b0fa74b-5215-402c-8f3a-c3c82a820280; __channelId=901045%2C0; ec=kJJtLlcT-1622174234622-091091de817d51983206136; notificationPreAuthorizeSwitch=121147; token=1833930374.1622174367679.c742a731d7063f063122196ad12891ca; refreshToken=1833930374.1622260767679.3d480d912925adf3bcb5942f344a9820; recommendId=%7B%22main-flow%22%3A%22baseline%22%2C%22off-feat%22%3A%22v1%22%2C%22feat-config%22%3A%22v1%22%2C%22model-version%22%3A%22v11%22%7D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1622174461; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1622305485; _efmdata=yqt2BxUJXEPvueQTcf5XAG%2BNnw1SRDv8sZT850dAT%2BscG3LzejZuEsE%2BDlOJWDOueNsodHONDg91neuQ5XEqlKVSROpokqlj8%2FL2zoct87w%3D; _exid=5lmjNGy2SGKtaP99JJ%2Bg3XBs3FHkON4JQT8plXalEciFItLWrcBd5AbbuGha%2FTJDeWR0ggr2Vfg1QD%2FjGHgXWQ%3D%3D; _pc_myzhenai_showdialog_=1; _pc_myzhenai_memberid_=%22%2C1833930374%22"
var ratelimit = time.Tick(1000 * time.Microsecond)

// 模拟浏览器访问
func WebFetch(url string) ([]byte, error) {
	<-ratelimit
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error: get url: %s", url)
	}
	// 设置header属性
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.62")
	reqest.Header.Set("Cookie", cookie)
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, fmt.Errorf("ERROR:get url: %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	//fmt.Printf("%s", utf8Reader)
	return ioutil.ReadAll(utf8Reader)
}

// 模拟浏览器访问
func WebFetch2(url string) (string, error) {
	<-ratelimit
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Error: get url: %s", url)
	}
	// 设置header属性
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.62")
	reqest.Header.Set("Cookie", cookie)
	resp, err := client.Do(reqest)
	if err != nil {
		return "", fmt.Errorf("ERROR:get url: %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && data == nil {
		log.Fatalln(err)
	}
	//fmt.Sprintf("%s",data)
	return fmt.Sprintf("%s", data), nil
}

// http代理访问
func Fetch2(weburl string) ([]byte, error) {
	<-ratelimit
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1087") //根据定义Proxy func(*Request) (*url.URL, error) 返回url.URL
	}
	transport := &http.Transport{Proxy: proxy}
	clien := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", weburl, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: get url: %s", weburl)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.62")
	resq, err := clien.Do(req)
	bodyReader := bufio.NewReader(resq.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error : %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
