package main

import "github.com/todoList_ver2/environment/router"

func main() {

	e := router.NewRouter()

	e.Start(":8888")
}
