package task

import (
	log "github.com/sirupsen/logrus"
	"mdu/explorer/data-syncer/config"
	"time"
)

type Task struct {
	name   string
	workFn func() error
	ticker *time.Ticker
	stop   chan int
}

func NewTask(name string, wFn func() error) *Task {
	t := &Task{}
	t.stop = make(chan int)
	t.ticker = time.NewTicker(time.Second * time.Duration(config.DefaultConfig.IntervalTime))
	t.workFn = wFn
	t.name = name

	return t
}

func (t *Task) Start() {
	go func() {
		log.Infof("task %s started", t.name)
		for {
			select {
			case <-t.ticker.C:
				log.Debugln(time.Now())
				err := t.workFn()
				if nil != err {
					panic(err)
				}
			case <-t.stop:
				t.ticker.Stop()
				break
			}
		}
	}()
}

func (t *Task) Stop() {
	close(t.stop)
}
