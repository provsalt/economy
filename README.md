# Economy
The "pretty" straight forward api to create an economy for your server

To start of just use the New function
```go
e := economy.New(economy.Connection{
		Username: "provsalt",
		Password: "Wowthatwasreallycool",
		IP:       "127.0.0.1:3306",
		Schema:   "economy",
	}, 3, 10)
```
Just enter the sql details and you're good to go
3 is the minimum connections and 10 is the maximum connections

## Balance
```go
ohno, bal := e.Balance(player)
	if ohno != nil {
		panic(ohno)
	}
	fmt.Println(bal)
  ```
