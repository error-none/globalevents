# Global Events Package

The `globalevents` package provides a simple and efficient mechanism for working with global events in an application. It is based on the [EventBus](https://github.com/asaskevich/EventBus) library and implements the Publisher-Subscriber pattern.

## Key Features

- Global access to the event bus through a singleton
- Thread-safe event bus initialization
- Simple API for event subscription and publishing

## Usage

### Subscribing to an Event

```go
err := globalevents.Subscribe("eventName", funcHandler)
if err != nil {
    // handle error
}
```

### Unsubscribing from an Event

```go
err := globalevents.Unsubscribe("eventName", funcHandler)
if err != nil {
    // handle error
}
```

### Publishing an Event

```go
globalevents.Publish("eventName", arg1, arg2, ...)
```

## Implementation Details

- Uses `sync.Once` for thread-safe event bus initialization
- Provides a single access point to the event bus through the `getBus()` function
- All event operations are performed through a single bus instance

## Usage Example

```go
package main

import (
    "fmt"
    "github.com/error-none/globalevents"
)

func main() {
    // Subscribe to an event from anywhere in the project
    globalevents.Subscribe("userCreated", func(userID string) {
        fmt.Printf("User created: %s\n", userID)
    })

    // Publish an event from anywhere in the project
    globalevents.Publish("userCreated", "123")
}
```

---

# Пакет globalevents

Пакет `globalevents` предоставляет простой и эффективный механизм для работы с глобальными событиями в приложении. Он основан на библиотеке [EventBus](https://github.com/asaskevich/EventBus) и реализует паттерн "Издатель-Подписчик" (Publisher-Subscriber).

## Основные возможности

- Глобальный доступ к шине событий через синглтон
- Потокобезопасная инициализация шины событий
- Простой API для подписки на события и их публикации

## Использование

### Подписка на событие

```go
err := globalevents.Subscribe("eventName", funcHandler)
if err != nil {
    // обработка ошибки
}
```

### Отписка от события

```go
err := globalevents.Unsubscribe("eventName", funcHandler)
if err != nil {
    // обработка ошибки
}
```

### Публикация события

```go
globalevents.Publish("eventName", arg1, arg2, ...)
```

## Особенности реализации

- Использует `sync.Once` для потокобезопасной инициализации шины событий
- Предоставляет единую точку доступа к шине событий через функцию `getBus()`
- Все операции с событиями выполняются через единый экземпляр шины

## Пример использования

```go
package main

import (
    "fmt"
    "github.com/error-none/globalevents"
)

func main() {
    // Подписка на событие в любом месте проекта
    globalevents.Subscribe("userCreated", func(userID string) {
        fmt.Printf("Пользователь создан: %s\n", userID)
    })

    // Публикация события в любом месте проекта
    globalevents.Publish("userCreated", "123")
}
```
