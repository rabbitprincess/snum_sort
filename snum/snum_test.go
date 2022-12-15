package snum

import (
	"fmt"
	"testing"
)

func Test_snum(t *testing.T) {
	{
		snum := NewSnum(0)
		fmt.Println(snum)
	}
	{
		var num int = -1
		snum := NewSnum(num)
		fmt.Println(snum.String())
	}
	{
		str := "1.2345"
		snum := NewSnum(str)
		fmt.Println(snum.String())
	}
	{
		snum := NewSnum(1)
		snum2 := NewSnum(snum)
		fmt.Println(snum2.String())
	}
}
