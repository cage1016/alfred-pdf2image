/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/cage/alfred-pdf2image/cmd"
)

func main() {
	cmd.Execute()
}
