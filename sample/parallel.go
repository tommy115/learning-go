package sample

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	mainLoopNum = 2000
	subLoopNum  = 1000
)

func init() {
	// importした順に呼ばれる？
	fmt.Println("init in parallel")
}

func ParallelSample() {
	logger := log.New(os.Stdout, "[ParallelSample] ", log.Ldate|log.Ltime|log.Lshortfile)
	/*
			logger.Println("mainRoutine")
			mainRoutine()
		logger.Println("mainRoutineNoName")
		mainRoutineNoName()
	*/
	logger.Println("env")
	outputEnvironment()
}

func mainRoutine() {
	go sub()
	for {
		time.Sleep(mainLoopNum)
		fmt.Println("main routine")
	}
}

func sub() {
	for {
		time.Sleep(subLoopNum)
		fmt.Println("sub loop")
	}
}

func mainRoutineNoName() {
	go func() {
		for {
			time.Sleep(subLoopNum)
			fmt.Println("sub loop")
		}
	}()
	for {
		time.Sleep(mainLoopNum)
		fmt.Println("main routine")
	}
}

func outputEnvironment() {
	// goルーチをを3つ作る。
	for index := 0; index < 3; index++ {
		go func() {
			//出力されないのはなぜ？
			fmt.Println("go routine")
		}()
	}
	time.Sleep(mainLoopNum)
	fmt.Printf("Num CPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoRoutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
