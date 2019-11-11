package main

// @Title 数据类型

import "fmt"

func main() {

	// 布尔型
	var boo1 bool = true
	var boo2 bool = false

	fmt.Println("boo1", boo1)
	fmt.Println("boo2", boo2)

	// 数字类型

	// 整型
	var int1 int = 123

	fmt.Println("int1", int1)
	fmt.Println("int1 , int » int8", int8(int1))

	var int2 int8 = 127

	fmt.Println("int2", int2)

	var int3 int16 = 32767

	fmt.Println("int3", int3)

	var int4 int32 = 2147483647

	fmt.Println("int4", int4)

	var int5 int64 = 9223372036854775807

	fmt.Println("int5", int5)
	fmt.Println("int5 , int64 » int", int(int5))

	var uint1 uint = 254

	fmt.Println("uint1", uint1)

	var uint2 uint8 = 255

	fmt.Println("uint2", uint2)

	var uint3 uint16 = 65535

	fmt.Println("uint3", uint3)

	var uint4 uint32 = 4294967295

	fmt.Println("uint4", uint4)

	var uint5 uint64 = 18446744073709551615

	fmt.Println("uint5", uint5)

	// 浮点型
	var float1 float32 = 123456789.0123456789

	fmt.Println("float1", float1)

	var float2 float64 = 123456789.0123456789

	fmt.Println("float2", float2)

	// 实数和虚数
	var complex1 complex64 = 1 + 2i

	fmt.Println("complex1", complex1)

	var complex2 complex128 = 2 + 3i

	fmt.Println("complex2", complex2)

	// 字符类型

}
