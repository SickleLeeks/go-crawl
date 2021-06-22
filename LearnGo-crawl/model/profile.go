package model

import "strconv"

type Profile struct {
	Name          string
	Marry         string
	Age           int
	Constellation string
	Height        int
	Weight        int
	WorkArea      string
	Salary        string
}

func (p Profile) String() string {
	return "姓名：" + p.Name + " 年龄 :" + string(p.Age) + " 身高:" + strconv.Itoa(p.Height) + " 婚姻状态：" + p.Marry + " 星座：" + p.Constellation + " 收入:" + p.Salary
}
