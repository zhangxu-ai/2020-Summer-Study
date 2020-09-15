package main

import (
	"./timewheel"
	"log"
	"net"
)


var tw =timewheel.NewTimeWheel(10,1, func(c interface{}) { //时间轮一共11s，每次1s
	log.Println("超时处理")
	c.(net.Conn).Close()
})

func main() {
	//创建监听的套接字
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Println("建立监听套接字失败，err = ", err)
		return
	}
	defer listen.Close()
	tw.Start()
	//创建用于通信的套接字
	for {
		conn, err := listen.Accept()
		log.Println("接收到新连接，", conn.RemoteAddr().String())
		if err != nil {
			log.Println("与客户端建立连接失败，err = ", err)
		}
		defer conn.Close()


		/*buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err != nil {
			log.Println("read failed,err = ",err)
			return
		} else {
			chan1 <- n
		}*/
		go handler(conn)
		//go func() {
		//	for true {
		//		timeout := time.Duration(10) * time.Second
		//		select {
		//		case <-time.After(timeout):
		//			log.Println("读取客户端心跳包超时，客户端可能已掉线！")
		//			conn.Close()
		//		case _ = <-chan11:
		//			log.Println("收到回复，取消超时检测")
		//			break
		//		}
		//	}
		//
		//}()
		//go handler(conn)
	}
}

func handler(conn net.Conn) {
	//获取客户端发来的消息
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf) //从套接字中读取的数据存储到buf切片中
		if err != nil {
			log.Println("从客户端读取数据失败，err = ", err)
			return
		}
		log.Println("客户端发来的数据为：", string(buf[:n]))
		tw.Active(conn)

		_, err1 := conn.Write([]byte("OK"))
		if err1 != nil {
			log.Println("读取失败！", err1)
			return
		}
	}
}
