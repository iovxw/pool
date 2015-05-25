package pool_test

import (
	"fmt"
	"testing"
	"time"

	p "github.com/Bluek404/pool"
)

func Test_Main(t *testing.T) {
	pool := p.NewPool(10, 2, func() {
		time.Sleep(time.Second * 1)
		fmt.Print(0)
	})

	pool.Run()

	pool.Wait()
}

func Test_Stop(t *testing.T) {
	pool := p.NewPool(-1, 10, func() {
		time.Sleep(time.Second * 1)
		fmt.Print(0)
	})

	pool.Run()
	fmt.Println("\nrunning")

	time.Sleep(time.Microsecond * 10)

	pool.Stop()
	fmt.Println("\nstopped")

	pool.Wait()
	fmt.Println("\ndone")
	time.Sleep(time.Second*2)
}
