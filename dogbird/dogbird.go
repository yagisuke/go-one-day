package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

type bird struct {
	name   string
	group  string
	length int
}

func (d dog) describe() string {
	description := "わたしは " + d.group
	description += "、名は" + d.name
	return description
}

func (d dog) bigger(d2 dog) string {
	description := d.describe()
	description += "、わたしは" + d2.name

	diff := d.height - d2.height
	switch {
	case diff > 5:
		description += "より大きいです"
	case diff < 5:
		description += "より小さいです"
	default:
		description += "と同じくらいです"
	}
	return description
}

func (b bird) describe() string {
	description := "わたしは " + b.group
	description += "、名は" + b.name
	return description
}

func main() {
	shibaken := dog{"シバ", "柴犬", 40}
	maru := dog{"マル", "マルチーズ", 22}
	mejiro := bird{"メジロ", "スズメ", 12}

	fmt.Println(shibaken.describe())
	fmt.Println(maru.describe())
	fmt.Println(mejiro.describe())
	fmt.Println(shibaken.bigger(maru))
	fmt.Println(maru.bigger(shibaken))
}
