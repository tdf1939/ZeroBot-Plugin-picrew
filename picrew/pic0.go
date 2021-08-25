package picrew

import (
	"math/rand"
	"time"
)

var Pic0 = SetCf(`185483`)

func ToPic0() string {
	rand.Seed(time.Now().UnixNano())
	Main := make(ImsMap)
	var ims0 = []string{
		"素体",
		"眉",
		"目_左",
		"目_右",
		"口",
		"服",
	}
	var ims1 = []string{
		"背景",
		"耳",
		"頬",
		"ほくろ",
		"目のハイライト_左",
		"目のハイライト_右",
	}
	var hand = []string{
		"手",
		"手２",
	}
	var hair0 = []string{
		"前髪",
		"後ろ髪",
	}
	var hair1 = []string{
		"横髪左",
		"横髪右",
	}
	var hair2 = []string{
		"アホ毛",
		"後ろ髪2",
		"ケモミミ",
	}
	var ps0 = []string{
		"眼鏡とか",
		"ヘアアクセ",
		"ヘアアクセ２",
		"帽子",
		"首元",
		"服装飾",
		"袖",
		"上着",
	}
	var ps1 = []string{
		"袖",
	}
	//必须组件
	Main.SetIm(Pic0, ims0, "", 1)
	//头发
	hcolall := Pic1[hair0[0]].Cols
	hcol := hcolall[rand.Intn(len(hcolall))]
	Main.SetIm(Pic0, hair0, hcol, 1)
	Main.SetIm(Pic0, hair1, hcol, 3)
	Main.SetIm(Pic0, hair2, hcol, 4)
	//选择手
	if hcol := rand.Intn(2); hcol == 1 {
		Main.SetIm(Pic0, []string{hand[0]}, "", 1)
		Main.SetIm(Pic0, ps1, "", 2)
	} else {
		Main.SetIm(Pic0, []string{hand[1]}, "", 1)
	}
	Main.SetIm(Pic0, ims1, "", 3)
	Main.SetIm(Pic0, ps0, "", 3)
	// 制图
	return Main.Save(50)
}
