/**
 ******************************************************************************
 * @file    main.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package main

import (
	"cnc/framework"
	"embed"
	"fmt"
	"github.com/gookit/color"
)

//go:embed all:template/dist
var Template embed.FS

func main() {
	fmt.Println("[cnc][main]：" + color.Gray.Text("starting..."))
	Framework.Init(Template)
}
