# Economy
The "pretty" straight forward api to create an economy for your server

To start of just use the New function
```go
package main

import (
	"github.com/provsalt/economy"
	"github.com/provsalt/economy/provider"
)

e := economy.New(provider.NewSQLite3("database/sqlite3.db"))
```

## Balance

```go
package main

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/provsalt/economy"
)

bal, ohno := e.Balance(player.XUID())
if ohno != nil {
panic(ohno)
}
fmt.Println(bal)
  ```

## Increase
```go
ohno = e.Increase(player.XUID(), 500)
if ohno != nil {
	panic(ohno)
}
```

## Decrease
```go
ohno = e.Decrease(player.XUID(), 500)
if ohno != nil {
	panic(ohno)
}
```

## Set
```go
ohno = e.Set(player.XUID(), 500)
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
