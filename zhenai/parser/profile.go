package parser

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\p{Han}]+座)[^<]+</div>`)

//var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
//var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
//var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([\p{Han}]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([[\p{Han}]+车)</div>`)

//ParseProfile 用户信息的解析器
func ParseProfile(contents []byte) engine.ParserResult {
	//声明用户信息的结构
	proflie := model.Profile{}

	//年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {

	} else {
		proflie.Age = age
	}

	//身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		log.Println("proflie.Height: ", err)
	} else {
		proflie.Height = height
	}
	//收入
	proflie.Income = extractString(contents, incomeRe)
	//星座
	proflie.Xinzuo = extractString(contents, xinzuoRe)
	//籍贯
	proflie.Hokou = extractString(contents, hokouRe)
	//房字
	proflie.House = extractString(contents, houseRe)
	//车
	proflie.Car = extractString(contents, carRe)
	//婚姻状况
	//proflie.Marriage = extractString(contents, marriageRe)

	result := engine.ParserResult{Items: []interface{}{proflie}}
	return result
}

//根据内容和对应的正则表达式，返回对应的字符串
func extractString(contents []byte, re *regexp.Regexp) string {
	//这里查找年龄只取一个值，所以就用FindSubmatch()
	// A return value of nil indicates no match.
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
