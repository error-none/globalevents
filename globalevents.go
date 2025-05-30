package globalevents

import (
	"sync"

	"github.com/asaskevich/EventBus"
)

var (
	once sync.Once
	bus  EventBus.Bus
)

func getBus() EventBus.Bus {
	once.Do(func() {
		bus = EventBus.New()
	})
	return bus
}

// Subscribe подписывает на событие
func Subscribe(event string, fn interface{}) error {
	return getBus().Subscribe(event, fn)
}

// Unsubscribe отписывает от события
func Unsubscribe(event string, fn interface{}) error {
	return getBus().Unsubscribe(event, fn)
}

// Publish отправляет событие
func Publish(event string, args ...interface{}) {
	getBus().Publish(event, args...)
}
