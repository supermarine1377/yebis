package main

import "supermarine1377/yebis/pkg"

func main() {
	app, err := pkg.NewApp()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
