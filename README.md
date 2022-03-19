# Economy
The "pretty" straight forward api to create an economy for your server

To start of just use the New function
```go
package main

import (
	"github.com/provsalt/economy"
	"github.com/provsalt/economy/provider"
)
func main() {
	e := economy.New(provider.NewSQLite("database/sqlite3.db"))
}
```

## Balance

```go
func main() {
    bal, ohno := e.Balance(player.UUID())
    if ohno != nil {
        panic(ohno) 
    }
    fmt.Println(bal)
}
  ```

## Increase
```go
ohno = e.Increase(player.UUID(), 500)
if ohno != nil {
	panic(ohno)
}
```

## Decrease
```go
ohno = e.Decrease(player.UUID(), 500)
if ohno != nil {
	panic(ohno)
}
```

## Set
```go
ohno = e.Set(player.UUID(), 500)
if ohno != nil {
	panic(ohno)
}
```

## Close
```go
ohno = e.Close()
if ohno != nil {
	panic(ohno)
}
```
