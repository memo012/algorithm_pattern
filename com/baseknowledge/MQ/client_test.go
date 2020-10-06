package MQ


import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestClient(t *testing.T)  {
	b := NewClient()
	b.SetConditions(100)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		topic := fmt.Sprintf("Golang梦工厂%d", i)
		payload := fmt.Sprintf("asong%d", i)

		ch, err := b.Subscribe(topic)
		if err != nil {
			t.Fatal(err)
		}

		wg.Add(1)
		go func() {
			e := b.GetPayLoad(ch)
			if e != payload {
				t.Fatalf("%s expected %s but get %s", topic, payload, e)
			}
			if err := b.Unsubscribe(topic, ch); err != nil {
				t.Fatal(err)
			}
			wg.Done()
		}()

		if err := b.Publish(topic, payload); err != nil {
			t.Fatal(err)
		}
	}

	wg.Wait()
}

// 使用一个定时器，向一个主题定时推送消息
func TestClient_Publish(t *testing.T) {
	m := NewClient()
	m.SetConditions(10)
	topic := "memolei"
	ch, err := m.Subscribe(topic)
	if err != nil {
		fmt.Println("subscribe failed")
		return
	}

	go OncePub(topic, m)
	OnceSub(ch, m)
	defer m.Close()
}

// 定时推送
func OncePub(topic string, c *Client)  {
	t := time.NewTimer(10 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			err := c.Publish(topic, "memolei加油")
			if err != nil {
				fmt.Println("pub message failed")
			}
		default:
		}
	}
}

// 接受订阅消息
func OnceSub(m <-chan interface{}, c *Client)  {
	for {
		val := c.GetPayLoad(m)
		fmt.Printf("get message is %s\n", val)
	}
}

// 使用一个定时器，定时向多个主题发送消息
func TestClient_Subscribe(t *testing.T) {
	// 多个topic测试
	m := NewClient()
	defer m.Close()
	m.SetConditions(10)
	topic := ""
	for i := 0; i < 10; i++ {
		topic = fmt.Sprintf("Golang_%02d", i)
		go Sub(m, topic)
	}
	ManyPub(m)
}

func ManyPub(c *Client) {
	t := time.NewTimer(10 * time.Second)
	defer t.Stop()
	fmt.Println(time.Now())
	for {
		select {
		case <-t.C:
			for i := 0; i < 10; i++ {
				// 多个topic推送不同的消息
				top := fmt.Sprintf("Golang_%02d", i)
				payload := fmt.Sprintf("memoGo_%02d", i)
				err := c.Publish(top, payload)
				if err != nil {
					fmt.Println("pub message failed")
				}
			}
		default:
		}
	}
}

func Sub(c *Client, top string)  {
	ch, err := c.Subscribe(top)
	if err != nil {
		fmt.Printf("sub top:%s faled\n", top)
	}
	for {
		val := c.GetPayLoad(ch)
		if val != nil {
			fmt.Printf("%s get message is %s\n", top, val)
		}
	}
}

