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

func ParallelSample() {
	logger := log.New(os.Stdout, "ParallelSample", log.Ldate|log.Ltime|log.Lshortfile)
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
	go func() {
		//		time.Sleep(10000)
		fmt.Println("go routine")
	}()
	fmt.Printf("Num CPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoRoutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
