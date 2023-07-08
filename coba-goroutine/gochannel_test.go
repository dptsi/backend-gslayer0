package cobagoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelGo(t *testing.T) {
	myChannel := make(chan string)
	defer close(myChannel)

	go func() {
		time.Sleep(2 * time.Second)
		myChannel <- "Bintang Nur"
		fmt.Println("Selesai mengirim data")
	}()

	rec := <-myChannel
	time.Sleep(5 * time.Second)
	fmt.Println(rec)
}
