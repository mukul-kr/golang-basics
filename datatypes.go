package main

import "fmt"

func main() {

	// datatypes
	// lets explore all datatypes present in golangss
	var a int8 = 10
	var b int16 = 10
	var c int32 = 10
	var d int64 = 10
	var e uint8 = 10
	var f uint16 = 10
	var g uint32 = 10
	var h uint64 = 10
	var i int = 8999999999999999999                    // max supported by int
	var j uint = 9999999999999999999                   // max supported by uint
	var k rune = 214444444                             // max supported by it
	var l byte = 255                                   // max limit after which it will throw error
	var m uintptr = 1999999999999999999                // max length supported  i.e after thiss it throws error
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m) //idk man sometimes highest value it supports..
	//

	// lets try float
	var n float32 = 154541546546545641521054548545848445451.3333333333355454845515452136546521654165465165416352654684514684644545454151152154154
	var o float64 = 151021521513665632631454364665651165546546546546548333554548455154521365465216541654651654163526546845146846445446354184654654844452415154154151524541521541541515636156443521465456325154684651321351354654654152135454568458463545846541521354635465635163545846545646546485468451654654546545646546546511111111111.0215333554548455154521365465216541654651654163526546845146846445441215215
	fmt.Println(n, o)
	//

	// go also supports complex type
	// lets try it 🙂
	var p complex128 = complex(1, 5) // so first argument is real part whereas second argument is imaginary
	var q complex64 = complex(1, 4)  // same as above
	fmt.Println(p, q)
	//

	// now lets try boolean values ☠️
	var r bool = true
	var s bool = b == (b + 2)
	t := 5 == 6
	fmt.Println(r, s, t)
	//

	// lets try string
	var u string = "hello world"
	v := "I love ...... but she doesn't love me back but so what let it be the way it is"
	fmt.Println(u)
	fmt.Println(v)
	//

	// Almost covered all datatypes present in golang
	// The end
}
