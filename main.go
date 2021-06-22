package main

import (
	"fmt"
	"regexp"
)

const str = `<div class="content"><table><tbody><tr><th><a href="http://album.zhenai.com/u/1050662423" target="_blank">简单生活</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>吉林白山</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>53</td> <td><span class="grayL">学   历：</span>中专</td> </tr> <tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身   高：</span>163</td></tr></tbody></table> <div class="introduce">有钱没钱无所谓，诚实守信，成熟稳重，热爱生活，热爱运动！</div></div> <div class="item-btn">打招呼</div></div><div class="list-item"><div class="photo"><a href="http://album.zhenai.com/u/1150089585" target="_blank"><img src="https://photo.zastatic.com/images/photo/287523/1150089585/7079241692022118.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="雨落南洋"></a><div class="content"><table><tbody><tr><th><a href="http://album.zhenai.com/u/1150089585" target="_blank">雨落南洋</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>男士</td> <td><span class="grayL">居住地：</span>吉林白山</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>61</td>  <td><span class="grayL">月   薪：</span>3000元以下</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身   高：</span>178</td></tr></tbody></table> <div class="introduce">人生靠自己，自己的路自己走，哪怕是千辛万苦，总会苦尽甘来</div></div> <div class="item-btn">打招呼</div></div><div class="list-item"><div class="photo"><a href="http://album.zhenai.com/u/1366749015" target="_blank"><img src="https://photo.zastatic.com/images/photo/341688/1366749015/47191374507209421.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="哈哈"></a></div> <div class="content"><table><tbody><tr><th><a href="http://album.zhenai.com/u/1366749015" target="_blank">哈哈</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>吉林白山</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>51</td> <td><span class="grayL">学   历：</span>高中及以下</td> </tr> <tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身   高：</span>169</td></tr></tbody></table> <div class="introduce">我希望用我的真诚能找到我的另一半，能一生呵护爱护我，我也会用我的一生来陪伴爱你</div></div> <div class="item-btn">打招呼</div></div><div class="list-item"><div class="photo"><a href="http://album.zhenai.com/u/1480617459" target="_blank"><img src="https://photo.zastatic.com/images/photo/370155/1480617459/7098550971512238.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="下一站的幸福"></a></div> <div class="content"><table><tbody><tr><th><a href="http://album.zhenai.com/u/1480617459" target="_blank">下一站的幸福</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>男士</td> <td><span class="grayL">居住地：</span>吉林白山</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>38</td>  <td><span class="grayL">月   薪：</span>5001-8000元</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身   高：</span>173</td></tr></tbody></table> <div class="introduce">我是一个不完美的人！许多事想不到！但是我有一颗包容的心！只要你不嫌弃！我就不放弃</div></div> <div class="item-btn">打招呼</div></div><div class="list-item"><div class="photo"><a href="http://album.zhenai.com/u/1389891283" target="_blank"><img src="https://photo.zastatic.com/images/photo/347473/1389891283/55377882768684878.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="新月"></a></div>`
const ageRe = `<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`
const marry = `<div data-v-8b1eac0c="" class="m-btn purple">(未婚|已婚|离异|丧偶)</div>`
const priceRe = `<span class="pl">定价:</span>([^<]+)<br>`
const scoreRe = `<strong class="ll rating_num " property="v:average">([^<]+)</strong>`
const infoRe = `<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`
const cityUrlRe = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`

//const cityUrlRe = `<a data-v-1573aa7c="" href="(http://www.zhenai.com/zhenghun/[^"]+)">([^\s]+)</a>`

const cityListRe = `<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`
func main() {
	re := regexp.MustCompile(cityRe)
	match := re.FindAllStringSubmatch(str,-1)
	fmt.Println(match[0][1])
	//if len(match) >= 2 {
	//	fmt.Println(match[1])
	//}
	//dom,err := goquery.NewDocumentFromReader(strings.NewReader(str))
	//if err!=nil{
	//	log.Fatalln(err)
	//}
	//dom.Find(".purple-btns div:nth-child(1)").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//})
}
