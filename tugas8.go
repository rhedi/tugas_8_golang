package main

import "fmt"
import "runtime"
import "math/rand"
import "time"

func main()  {
  rand.Seed(time.Now().Unix())
  runtime.GOMAXPROCS(2)

  var pesan = [] string {"Apa Kabar Teman-Teman"}
  fmt.Println("Pesan :", pesan)
  var message = make (chan int)

  go kirimdata(pesan, message)
  terimadata(pesan, message)
}

func kirimdata(pesan []string, ch chan <- int)  {
  for i:=0;true;i++{
    ch <- i
    time.Sleep(time.Duration(rand.Int()%10 +1) * time.Second)
  }
}
func terimadata(pesan []string, ch <- chan int)  {
  loop:
  for{
    select{
    case data := <-ch:
      fmt.Print("Apa kabar teman-teman ", data, "\n")
    case <-time.After(time.Second * 5):
      fmt.Println("Timeout, Tidak ada aktivitas dalam 5 detik")
      break loop
    }
  }
}
