package main

import (
	"fmt"
)

func main() {
	test1()
	// test2()
	// test3()
	// testReadConfig("192.168.111.10")
	testUpdateConfig("192.168.111.10")
	// testStop(917468916832346395)
	// testStart()
}

func test1() {
	fmt.Println(testDo(1)) // testclient -> A -> B
	fmt.Println(testDo(10))
	fmt.Println(testDo(100))
}

func test2() {
	testLB()
}

func test3() {
	testmsg()
	testmsgp()

}
