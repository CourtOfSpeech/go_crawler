package persist

import (
	"crawler/model"
	"testing"
)

func Test_save(t *testing.T) {
	type args struct {
		item interface{}
	}
	testData := model.Profile{
		Name:   "在水伊人",
		Age:    44,
		Height: 155,
		Income: "8千-1.2万",
		Xinzuo: "魔羯座",
		Hokou:  "四川成都",
		House:  "已购房",
		Car:    "未买车"}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"在水伊人", args{item: testData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			save(tt.args.item)
		})
	}
}
