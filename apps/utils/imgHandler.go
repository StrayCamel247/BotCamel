package utils

// Copyright 2010 The Freetype-Go Authors. All rights reserved.
// Use of this source code is governed by your choice of either the
// FreeType License or the GNU General Public License version 2 (or
// any later version), both of which can be found in the LICENSE file.

// +build example
//
// This build tag means that "go install github.com/golang/freetype/..."
// doesn't install this example program. Use "go run main.go" to run it or "go
// install -tags=example" to install it.

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

var (
	dpi      = flag.Float64("dpi", 100, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./apps/utils/chinese.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "full", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white textArray on a black background")
)

func init() {

}

// var textArray = []string{
// 	`https://github.com/StrayCamel247/BotCamel 领bug修修我吧(๑•̀ㅂ•́)و✧`,
// 	`======BotCamel基础功能=====`,
// 	`++ [menu 菜单 功能] 功能列表 0x0`,
// 	`======命运2基础功能========`,
// 	`++ [week 周报] 周报信息查询 0x1`,
// 	`++ [day 日报] 日报信息查看 0x2`,
// 	`++ [xiu nine 老九] 老九信息查询 0x3`,
// 	`++ [trial 试炼 train] 试炼最新动态 0x4`,
// 	`++ [dust 光尘] 赛季光尘商店 0x5`,
// 	`++ [random 骰子] 骰子功能 0x6`,
// 	`++ [perk 词条] 查询词条/模组信息 0x7`,
// 	`++ [item 物品] 查询物品信息-提供light-gg信息 0x8`,
// 	`++ [npc] 查询npc信息 0x9`,
// 	`++ [skill] 查询技能等信息 0xa`,
// 	`++ [pve] 查询pve信息 0xb`,
// 	`++ [pvp] 查询pvp信息 0xc`,
// }

// 按照回车切分字符串
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 文本处理成数组
// 返回数组-数组最长元素的长度-数组长度
func textArrayHandler(text string) (arr []string, maxL, height int) {
	arr = strings.Split(text, "\n")
	for _, v := range arr {
		maxL = max(maxL, utf8.RuneCountInString(v))
	}
	height = len(arr)
	return
}

// Text2ImgHandler
// 文本转图片
func Text2ImgHandler(text string) (filePath string, err error) {
	textArray, textWight, textHeight := textArrayHandler(text)
	flag.Parse()
	println(textWight, textHeight)
	dir, _ := os.Getwd()
	filePath = path.Join(dir, "test.menu")
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {

		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {

		return
	}

	// Initialize the contextArray.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, textWight*8, textHeight*27))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the textArray.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(*size)>>6))
	for _, s := range textArray {
		_, err = c.DrawString(s, pt)
		if err != nil {
			return
		}
		pt.Y += c.PointToFixed(*size * *spacing)
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create(filePath)
	if err != nil {
		return

	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		return

	}
	err = b.Flush()
	if err != nil {
		return

	}
	log.Infof(fmt.Sprintf("[图片生成]%s", filePath))
	return
}

// func main() {
// 	var textArray = []string{
// 		`https://github.com/StrayCamel247/BotCamel 领bug修修我吧(๑•̀ㅂ•́)و✧`,
// 		`======BotCamel基础功能=====`,
// 		`++ [menu 菜单 功能] 功能列表 0x0`,
// 		`======命运2基础功能========`,
// 		`++ [week 周报] 周报信息查询 0x1`,
// 		`++ [day 日报] 日报信息查看 0x2`,
// 		`++ [xiu nine 老九] 老九信息查询 0x3`,
// 		`++ [trial 试炼 train] 试炼最新动态 0x4`,
// 		`++ [dust 光尘] 赛季光尘商店 0x5`,
// 		`++ [random 骰子] 骰子功能 0x6`,
// 		`++ [perk 词条] 查询词条/模组信息 0x7`,
// 		`++ [item 物品] 查询物品信息-提供light-gg信息 0x8`,
// 		`++ [npc] 查询npc信息 0x9`,
// 		`++ [skill] 查询技能等信息 0xa`,
// 		`++ [pve] 查询pve信息 0xb`,
// 		`++ [pvp] 查询pvp信息 0xc`,
// 	}
// 	// textArray := textArrayHandler(text)
// 	flag.Parse()

// 	// Read the font data.
// 	fontBytes, err := ioutil.ReadFile(*fontfile)
// 	if err != nil {
//
// 		return
// 	}
// 	f, err := freetype.ParseFont(fontBytes)
// 	if err != nil {
//
// 		return
// 	}

// 	// Initialize the contextArray.
// 	fg, bg := image.Black, image.White
// 	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
// 	if *wonb {
// 		fg, bg = image.White, image.Black
// 		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
// 	}
// 	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
// 	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
// 	c := freetype.NewContext()
// 	c.SetDPI(*dpi)
// 	c.SetFont(f)
// 	c.SetFontSize(*size)
// 	c.SetClip(rgba.Bounds())
// 	c.SetDst(rgba)
// 	c.SetSrc(fg)
// 	switch *hinting {
// 	default:
// 		c.SetHinting(font.HintingNone)
// 	case "full":
// 		c.SetHinting(font.HintingFull)
// 	}

// 	// Draw the guidelines.
// 	for i := 0; i < 200; i++ {
// 		rgba.Set(10, 10+i, ruler)
// 		rgba.Set(10+i, 10, ruler)
// 	}

// 	// Draw the textArray.
// 	pt := freetype.Pt(10, 10+int(c.PointToFixed(*size)>>6))
// 	for _, s := range textArray {
// 		_, err = c.DrawString(s, pt)
// 		if err != nil {
//
// 			return
// 		}
// 		pt.Y += c.PointToFixed(*size * *spacing)
// 	}

// 	// Save that RGBA image to disk.
// 	outFile, err := os.Create("out.png")
// 	if err != nil {
//
// 		os.Exit(1)
// 	}
// 	defer outFile.Close()
// 	b := bufio.NewWriter(outFile)
// 	err = png.Encode(b, rgba)
// 	if err != nil {
//
// 		os.Exit(1)
// 	}
// 	err = b.Flush()
// 	if err != nil {
//
// 		os.Exit(1)
// 	}
// }
