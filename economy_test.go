package economy

import (
	"github.com/google/uuid"
	"github.com/provsalt/economy/handler"
	"github.com/provsalt/economy/provider"
	"os"
	"testing"
)

// TestEconomy tests the economy package
func TestEconomy(t *testing.T) {
	id := uuid.New()
	sql, err := provider.NewSQLite("test.sqlite")
	if err != nil {
		panic(err)
	}
	e := New(sql, handler.NopEconomyHandler{})
	if e == nil {
		t.Error("New: Economy is nil")
	}

	err = e.Set(id, 1000)

	if err != nil {
		t.Error(err)
	}

	bal, err := e.Balance(id)

	if err != nil {
		t.Error(err)
	}

	if bal != 1000 {
		t.Errorf("Init: Balance is not 1000, But %d", bal)
	}

	err = e.Increase(id, 100)

	if err != nil {
		t.Error(err)
	}

	bal, err = e.Balance(id)

	if bal != 1100 {
		t.Errorf("Increase: Balance is not 1100, but %d", bal)
	}

	err = e.Decrease(id, 200)

	if err != nil {
		t.Error(err)
	}

	bal, err = e.Balance(id)

	if err != nil {
		t.Error(err)
	}

	if bal != 900 {
		t.Errorf("Decrease: Balance is not 900, but %d", bal)
	}

	err = e.Close()

	if err != nil {
		t.Error(err)
	}

	_ = os.Remove("test.sqlite")
}
