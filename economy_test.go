package economy

import (
	"github.com/provsalt/economy/provider"
	"os"
	"testing"
)

var xuid = "830188016623423366"

// TestEconomy tests the economy package
func TestEconomy(t *testing.T) {
	sql, err := provider.NewSQLite("test.sqlite")
	if err != nil {
		panic(err)
	}
	e := New(sql)
	if e == nil {
		t.Error("New: Economy is nil")
	}

	err = e.Set(xuid, 1000)

	if err != nil {
		t.Error(err)
	}

	bal, err := e.Balance(xuid)

	if err != nil {
		t.Error(err)
	}

	if bal != 1000 {
		t.Errorf("Init: Balance is not 1000, But %d", bal)
	}

	err = e.Increase(xuid, 100)

	if err != nil {
		t.Error(err)
	}

	err = e.Reduce(xuid, 200)

	if err != nil {
		t.Error(err)
	}

	bal, err = e.Balance(xuid)

	if err != nil {
		t.Error(err)
	}

	if bal != 900 {
		t.Errorf("Increase/Decrease: Balance is not 900, but %d", bal)
	}

	err = e.Close()

	if err != nil {
		t.Error(err)
	}

	_ = os.Remove("test.sqlite")
}
