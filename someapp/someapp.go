package main

import "fmt"

type someapp struct {
	username string
	use      int
	isopen   bool
}

func (app *someapp) open() {
	if app.isopen {
		fmt.Printf("%sさんのアプリはすでに開いています\n\n", app.username)
	} else {
		app.use++
		if app.use > 2 {
			fmt.Printf("%sさんの使用期間は満了です。ご購入をご検討ください\n\n", app.username)
			app.isopen = true
		}
	}
}

func (app *someapp) close() {
	if app.isopen {
		fmt.Printf("さようなら%sさん、アプリを終了します\n\n", app.username)
		app.isopen = false
	}
}

func newapp(username string) someapp {
	fmt.Printf("ようこそ%sさん、初めてのご使用になります\n", username)
	return someapp{username, 1, true}
}

func main() {
	fmt.Println("[apple さんがアプリを開く]")
	appleApp := newapp("apple")
	appleApp.close()

	fmt.Println("[apple さんが再びアプリを再び開く]")
	appleApp.open()
	appleApp.close()

	fmt.Println("[orange さんがアプリを開く]")
	orangeApp := newapp("orange")
	orangeApp.close()

	fmt.Println("[orange さんが再びアプリを再び開く]")
	orangeApp.open()

	fmt.Println("[apple さんが3回目のアプリを再び開く]")
	appleApp.open()

	fmt.Println("[orange さんが3回目のアプリを再び開く]")
	orangeApp.open()
}
