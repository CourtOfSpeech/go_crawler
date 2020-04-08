package model

//Profile 用户的信息
type Profile struct {
	Name       string `json:"name,omitempty" xml:"name"`
	Gender     string `json:"gender,omitempty" xml:"gender"`
	Age        int    `json:"age,omitempty" xml:"age"`
	Height     int    `json:"height,omitempty" xml:"height"`
	Weight     int    `json:"weight,omitempty" xml:"weight"`
	Income     string `json:"income,omitempty" xml:"income"`
	Marriage   string `json:"marriage,omitempty" xml:"marriage"`
	Education  string `json:"education,omitempty" xml:"education"`
	Occupation string `json:"occupation,omitempty" xml:"occupation"`
	Hokou      string `json:"hokou,omitempty" xml:"hokou"`
	Xinzuo     string `json:"xinzuo,omitempty" xml:"xinzuo"`
	House      string `json:"house,omitempty" xml:"house"`
	Car        string `json:"car,omitempty" xml:"car"`
}
