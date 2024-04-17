# snum-sortable
 - Sortable decimal number utility
 - No zero padding for alignment
 - Support bytes alignment including negative numbers and decimals
 - Supports up to 96 integers and 32 decimal places
 - implement ericlagergren/decimal ( github.com/ericlagergren/decimal )


## Quick Start

```go

package main

import (
	"fmt"

	"github.com/rabbitprincess/snum_sort/snum"
	"github.com/rabbitprincess/snum_sort/sort"
)

func main() {
	sn1 := snum.New("123456789.987654321")
	fmt.Println("snum :", sn1)

	st1 := sort.New(sn1)
	fmt.Println("sort :", st1)

	enc, _ := st1.MarshalJSON()
	fmt.Println("json :", enc)

	st1New := sort.New(0)
	st1New.UnmarshalJSON(enc)
	fmt.Println("new :", st1New)
}

```

# max length
 - positive : 65 byte (header 1 bt + body 64 bt)
 - negative : 66 byte (header 1 bt + body 64 bt + 0xFF 1 bt)

# header
 - 255 = positive 96 digit ( 1e95 <= x < 1e96 )
 - 254 = positive 95 digit ( 1e94 <= x < 1e95 )
 .....
 - 160 = positive 1 digit ( 1 <= x < 10 )
 - 159 = positive -1 digit ( 0.1 <= x < 1 )
 .....
 - 128 = positive -32 digit ( 0 <= x < 1e-31 ) - ! include zero !

 - 127 = negative -32 digit ( -1e-31 < x <= -1e-32 )
 .....
 - 96 = negative -1 digit ( -1 < x <= -0.1 )
 - 95 = negative 1 digit ( -10 < x <= -1 )
 .....
 - 1 = negative 95 digit ( -1e95 < x <= -1e94 )
 - 0 = negative 96 digit ( -1e96 < x <= -1e95 )

buf[1:]
 - Decimal string compressed to 1 byte per 2 digits
 - if negative, stored as complement
