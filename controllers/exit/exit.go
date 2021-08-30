package exit

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logrus.Error("Exit reason: ", sig)
		close(terminate)
	}()

	<-terminate
	cb()
}
