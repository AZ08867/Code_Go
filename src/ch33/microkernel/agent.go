package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const (
	Waiting = iota
	Running
	Stopped
)

type EventReceiver interface {
	OnEvent(evt Event)
}

type Event struct {
	Src  string
	Data string
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type CollectorError struct {
	Name string
	Err  error
}

type CollectorsError struct {
	CollectorErrors []CollectorError
}

func (ce CollectorsError) Error() string {
	var errStr string
	for _, err := range ce.CollectorErrors {
		errStr += fmt.Sprintf("%s: %v\n", err.Name, err.Err)
	}
	return errStr
}

var ErrWrongStateError = errors.New("agent in wrong state")

type Agent struct {
	collectors    map[string]Collector
	evtBuf        chan Event
	cancel        context.CancelFunc
	ctx           context.Context
	state         int
	eventHandlers map[string]func(event Event)
}

func (agt *Agent) RegisterEventHandler(eventType string, handler func(event Event)) {
	if agt.eventHandlers == nil {
		agt.eventHandlers = make(map[string]func(event Event))
	}
	agt.eventHandlers[eventType] = handler
}

func (agt *Agent) EventProcessGoroutine() {
	var evtSeg [10]Event
	for {
		evtIndex := 0
		for evtIndex < 10 {
			select {
			case evt, ok := <-agt.evtBuf:
				if !ok {
					return // 通道已关闭，退出goroutine
				}
				evtSeg[evtIndex] = evt
				evtIndex++
			case <-agt.ctx.Done():
				return
			}
		}

		// 处理收集到的事件
		agt.processEvents(evtSeg[:evtIndex]) // 只处理实际接收到的事件
	}
}

func (agt *Agent) processEvents(events []Event) {
	for _, evt := range events {
		if handler, ok := agt.eventHandlers[evt.Src]; ok {
			handler(evt)
		} else {
			// 如果没有找到对应的处理器，可以记录日志或执行默认操作
			fmt.Printf("未找到事件 %s 的处理器\n", evt.Src)
		}
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: make(map[string]Collector),
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
	return &agt
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return ErrWrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
	var errs CollectorsError
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for name, collector := range agt.collectors {
		wg.Add(1)
		go func(name string, collector Collector, ctx context.Context) {
			defer wg.Done()
			err := collector.Start(ctx)
			if err != nil {
				mutex.Lock()
				errs.CollectorErrors = append(errs.CollectorErrors,
					CollectorError{Name: name, Err: err})
				mutex.Unlock()
			}
		}(name, collector, agt.ctx)
	}
	wg.Wait()
	if len(errs.CollectorErrors) > 0 {
		return errs
	}
	return nil
}

func (agt *Agent) stopCollectors() error {
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err := collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				CollectorError{Name: name, Err: err})
		}
	}
	if len(errs.CollectorErrors) > 0 {
		return errs
	}
	return nil
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return ErrWrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcessGoroutine()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return ErrWrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) destroyCollectors() error {
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err := collector.Destroy(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				CollectorError{Name: name, Err: err})
		}
	}
	if len(errs.CollectorErrors) > 0 {
		return errs
	}
	return nil
}

func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return ErrWrongStateError
	}
	return agt.destroyCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}
