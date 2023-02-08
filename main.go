package main

import "app/db"

func main() {
	db.Init()
	defer db.Client.Close()
}
