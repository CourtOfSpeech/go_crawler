package parser

import (
	"crawler/engine"
	"crawler/model"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseProfile(t *testing.T) {
	type args struct {
		contents []byte
		name     string
	}
	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/1402882293")
	//fmt.Println(strconv.QuoteToASCII("房"))
	contents, err := ioutil.ReadFile("profile_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)
	tests := []struct {
		name string
		args args
		want engine.ParserResult
	}{
		// TODO: Add test cases.
		{"在水伊人", args{contents, "在水伊人"},
			engine.ParserResult{
				Items: []engine.Items{
					engine.Items{
						URL:  "http://album.zhenai.com/u/1402882293",
						ID:   "http://album.zhenai.com/u/1402882293",
						Type: "zhenhun",
						Payload: model.Profile{
							Name:   "在水伊人",
							Age:    44,
							Height: 155,
							Income: "8千-1.2万",
							Xinzuo: "魔羯座",
							Hokou:  "四川成都",
							House:  "已购房",
							Car:    "未买车",
						}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseProfile(tt.args.contents, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
