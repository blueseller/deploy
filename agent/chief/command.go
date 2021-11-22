package chief

import (
	"fmt"
	"sync"
	"time"

	commandPb "github.com/blueseller/deploy.git/api/agent/command/v1"
)

type CommandServices struct{}

func (c *CommandServices) Command(allStr commandPb.StreamCommandSerivce_CommandServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, err := allStr.Recv()
			if err != nil {
				break
			}
			fmt.Println(data)
		}
		wg.Done()
	}()

	go func() {
		for {
			//allStr.Send()
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
