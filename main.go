package pool

import (
	"fmt"
	"net"
)

func Main(){

	// 工厂模式，提供创建连接的工厂方法
	factory    := func() (net.Conn, error) { return net.Dial("tcp", "127.0.0.1:4000") }

	// 创建一个tcp池，提供初始容量和最大容量以及工厂方法
	p, err := NewChannelPool(5, 30, factory)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	// 获取一个连接
	conn, err := p.Get()

	// Close并不会真正关闭这个连接，而是把它放回池子，所以你不必显式地Put这个对象到池子中
	conn.Close()

	// 通过调用MarkUnusable, Close的时候就会真正关闭底层的tcp的连接了
	if pc, ok := conn.(*PoolConn); ok {
		pc.MarkUnusable()
		pc.Close()
	}

	// 关闭池子就会关闭=池子中的所有的tcp连接
	p.Close()

	// 当前池子中的连接的数量
	current := p.Len()
	fmt.Printf("current length is %d", current)
}
