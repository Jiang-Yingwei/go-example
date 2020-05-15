package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//hello.ForTest()
	//hello.IfTest(12)
	//hello.SwitchTest()
	//hello.ArrayTest()
	//hello.SliceTest()
	//hello.MapTest()
	//hello.RangeTest()

	//_, total := hello.ValuesAndParams(1, 2, 3)
	//fmt.Println(total)
	//hello.ClosuresTest()
	//hello.PointersTest()
	//hello.StructTest()
	//hello.InterfaceTest()

	//thread.GoroutineTest()
	//thread.ChannelTest()
	//thread.ChannelSelectTest()
	//thread.ChannelTimeoutTest()
	//thread.NonBlockingChannelTest()
	//thread.ChannelCloseTest()
	//thread.RangeOverChannelsTest()
	//thread.TimerTest()
	//thread.TickerTest()
	//thread.WorkerPoolTest()
	//thread.RateLimitingTest()
	//thread.AtomicCounterTest()
	//thread.MutexTest()
	//thread.StatefulTest()
	fmt.Println("app elapsed:", time.Since(t))
}
