package main

import "fmt"
import settings "github.com/andrewrothstein/dcbv2/setting"

func main() {
	x := settings.CreateLiteralSetting("name", "value")
	xi, xj := x.Get()
	fmt.Println(x, xi, xj)
	y := settings.CreateEnvSetting([]string{"TARGET", "USER"}, "_", "DCB")
	yi, yj := y.Get()
	fmt.Println(y, yi, yj)
	z := settings.CwdSetting{}
	zi, zj := z.Get()
	fmt.Println(z, zi, zj)
}
