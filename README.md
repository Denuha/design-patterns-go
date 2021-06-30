# design-patterns-go

Реализация некоторых паттернов проектирования в среде golang.
Шаблоны взяты от сюда: https://refactoring.guru/ru/design-patterns

## Adapter

`
go run adapter/main.go 
`

Дано: круглое отверстие и круглые колышки (в пакете `round`) и квадратные колышки в пакете (`square`)
Задача: необходимо просчитать возможность вместимости квадратного колышка в круглое отверстие

Используется переопределение функции `GetRadius()` интерфейса круглых колышек в интерфейсе адаптера для квадратных колышков.
## Command

`
go run command/main.go
`

В данной реализации происходит ведение истории выполненных операций.
Операции выполняются просто в виде логирования, сообщая в CLI что они произошли:).

Команды (копировать, вырезать, вставить, отменить) приводятся к общему типу `ICommand`, что позволяет хранить их в одном массиве, а так же вызывать их с помощью метода `ExecuteCommand()`
