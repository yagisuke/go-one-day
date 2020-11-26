package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

func main() {
	pome := dog{"ポメ", "ポメラニアン", 25}
	maru := dog{"マル", "マルチーズ", 22}
	shiba := dog{"シバ", "柴犬", 40}

	dogs := []dog{pome, maru, shiba}

	for i := 0; i < len(dogs); i++ {
		fmt.Printf("%s の %s さんです\n", dogs[i].group, dogs[i].name)
	}
}
