package main

import "supermarine1377/yebis/src"

func main() {
	app, err := src.NewApp()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
