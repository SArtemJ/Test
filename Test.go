package main

//
//
//import (
//	//"log"
//	"math/rand"
//	"time"
//	"log"
//)
//
//type Phone struct {
//	ID int
//}
//
//type Message struct {
//	ID int
//	PhoneID int
//	Message string
//}
//
//func createChanPhones() chan Phone {
//	cp := make(chan Phone)
//	go func() {
//		for i := 0; i < 100; i++ {
//			var nP= Phone{}
//			nP.ID = i
//			cp <- nP
//		}
//		close(cp)
//	}()
//
//	return cp
//}
//
//func main() {
//
//	phones := make(chan Phone)
//	phones = createChanPhones()
//	//for v := range phones {
//	//	log.Println(v)
//	//}
//
//	message := make(chan Message)
//	message = createM(phones)
//	//time.Sleep(2 * time.Second)
//
//
//	for v := range message {
//		log.Println(v)
//	}
//}
//
//func createM(cp chan Phone) chan Message {
//	m := make(chan Message)
//
//	go func() {
//		for v := range cp {
//			go func() {
//				var newM= Message{}
//				newM.ID = rand.Intn(1000)
//				newM.PhoneID = v.ID
//				newM.Message = "some m"
//				m <- newM
//				time.Sleep(1*time.Second)
//				close(m)
//			}()
//		}
//	}()
//	return m
//}

//func(chan T) chan Done {
//	c := make(chan Done)
//	var index int
//
//	go func(){
//		for v := range T {
//			go func() {
//				go func() {
//					index = createNewIndex
//				}()
//				createNewObj.index = index
//				createNewObj.id = v.id
//				c <- createNewObj
//			}()
//		}
//	}()
//}
