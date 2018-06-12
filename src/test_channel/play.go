package test_channel

import (
  "sync"
  "fmt"
  "time"
  "math/rand"
)

var wg sync.WaitGroup  //用来等待程序完成

//main()函数前执行
func init() {
  rand.Seed(time.Now().UnixNano())
}

func main() {
  //创建一个无缓存通道
  court := make(chan int)
  
  //计数加2，表示等待两个goroutine完成
  wg.Add(2)
  
  //启动两个goroutine代表两名选手
  go player("jobs", court)
  go player("cook", court)
  
  //发球，给通道赋值
  court <- 1
  
  //等待比赛结果
  wg.Wait()
}

func player(name string, court chan) {
  //在函数完成是调用done通知main程序完成
  defer wg.Done()
  
  for {
    //等待球被击打过来
    ball, ok := court
    if !ok {
      //如果通道被关闭，说明对手已经丢球
      fmt.Printf("Player %s Won\n", name)
      return
    } 
    
    //通过产生的随机数求13的余数，来决定选手是否丢球
    n := rand.Intn(100)
    if n % 13 == 0 {
      fmt.Printf("Player %s missed\n", name)
      close(court)
      return
    }
    
    //比赛还在进行，显示击球数，并将击球数加1
    fmt.Printf("Player %s hit %d\n", name, ball)
    ball++
    
    //将球击出给对手
    court <- ball
  }
}
