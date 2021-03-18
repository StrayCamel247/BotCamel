package baseapis

// /*
//    __author__ : stray_camel
//   __description__ : select channel 结合使用
//   __REFERENCES__:
//   __date__: 2021-03-18
// */
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"reflect"
// 	"time"
// )

// type (
// 	product struct {
// 		id  int // 生产者序号
// 		val int // 产品
// 	}
// 	producer struct {
// 		id   int // 序号
// 		chnl chan *product
// 	}
// )

// var (
// 	producerList []*producer
// 	notifynew    chan int
// 	updatedone   chan int
// )

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	notifynew = make(chan int, 1)
// 	updatedone = make(chan int, 1)
// 	ticker := time.NewTicker(time.Second)
// 	cases := update(ticker)
// 	for {
// 		chose, value, _ := reflect.Select(cases)
// 		switch chose {
// 		case 0: // 有新的生产者
// 			cases = update(ticker)
// 			updatedone <- 1
// 		case 1:
// 			// 创建新的生产者
// 			if len(producerList) < 5 {
// 				go newproducer()
// 			}
// 		default:
// 			item := value.Interface().(*product)
// 			fmt.Printf("消费: 值=%d 生产者=%d\n", item.val, item.id)
// 		}
// 	}
// }
// func update(ticker *time.Ticker) (cases []reflect.SelectCase) {
// 	// 新生产者通知
// 	selectcase := reflect.SelectCase{
// 		Dir:  reflect.SelectRecv,
// 		Chan: reflect.ValueOf(notifynew),
// 	}
// 	cases = append(cases, selectcase)
// 	// 定时器
// 	selectcase = reflect.SelectCase{
// 		Dir:  reflect.SelectRecv,
// 		Chan: reflect.ValueOf(ticker.C),
// 	}
// 	cases = append(cases, selectcase)
// 	// 每个生产者
// 	for _, item := range producerList {
// 		selectcase = reflect.SelectCase{
// 			Dir:  reflect.SelectRecv,
// 			Chan: reflect.ValueOf(item.chnl),
// 		}
// 		cases = append(cases, selectcase)
// 	}
// 	return
// }
// func newproducer() {
// 	newitem := &producer{
// 		id:   len(producerList) + 1,
// 		chnl: make(chan *product, 100),
// 	}
// 	producerList = append(producerList, newitem)
// 	notifynew <- 1
// 	<-updatedone
// 	go newitem.run()
// }
// func (this *producer) run() {
// 	for {
// 		time.Sleep(time.Duration(int(time.Millisecond) * (rand.Intn(1000) + 1)))
// 		item := &product{
// 			id:  this.id,
// 			val: rand.Intn(1000),
// 		}
// 		fmt.Printf("生产: 值=%d 生产者=%d\n", item.val, item.id)
// 		this.chnl <- item
// 	}
// }
