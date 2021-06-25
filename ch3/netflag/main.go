package main

import (
	"fmt"
	tmp "net"
)

func IsUp(v tmp.Flags) bool     { return v&tmp.FlagUp == tmp.FlagUp }
func TurnDown(v *tmp.Flags)     { *v &^= tmp.FlagUp }
func SetBroadcast(v *tmp.Flags) { *v |= tmp.FlagBroadcast }
func IsCast(v tmp.Flags) bool   { return v&(tmp.FlagBroadcast|tmp.FlagMulticast) != 0 }

func main() {

	var v tmp.Flags = tmp.FlagMulticast | tmp.FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
