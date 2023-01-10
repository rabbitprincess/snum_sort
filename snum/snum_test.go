package snum

import (
	"fmt"
	"testing"
)

func Test_snum(t *testing.T) {
	{
		snum := New(0)
		fmt.Println(snum)
	}
	{
		var num int = -1
		snum := New(num)
		fmt.Println(snum.String())
	}
	{
		str := "1.2345"
		snum := New(str)
		fmt.Println(snum.String())
	}
	{
		snum := New(1)
		snum2 := New(snum)
		fmt.Println(snum2.String())
	}
}
