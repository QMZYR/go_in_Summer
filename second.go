package main

import (
	"fmt"
	"sync"
)

func san(zs, ls chan bool) {
	defer waitGroup.Done()
	defer close(ls)
	for i := 0; i < 100; i++ {
		<-zs
		fmt.Println("李四")
		ls <- true

	}
}
func si(ls, ww chan bool) {
	defer waitGroup.Done()
	defer close(ww)
	for i := 0; i < 100; i++ {
		<-ls
		fmt.Println("王五")
		ww <- true

	}
}
func wu(ww, zl chan bool) {
	defer waitGroup.Done()
	defer close(zl)
	for i := 0; i < 100; i++ {
		<-ww
		fmt.Println("赵六")
		zl <- true

	}
}
func liu(zl, zs chan bool) {
	defer waitGroup.Done()
	defer close(zs)
	for i := 0; i < 100; i++ {
		<-zl
		fmt.Println("张三")
		zs <- true

	}
}

var waitGroup sync.WaitGroup

func main() {
	zs := make(chan bool, 1)
	ls := make(chan bool, 1)
	ww := make(chan bool, 1)
	zl := make(chan bool, 1)
	zl <- true
	go san(zs, ls)
	go si(ls, ww)
	go wu(ww, zl)
	go liu(zl, zs)
	waitGroup.Add(4)
	waitGroup.Wait()
}
