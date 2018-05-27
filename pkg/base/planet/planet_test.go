package planet

import(
    "testing"
)

func TestEmpty(t *testing.T) {
    t.Skip("Empty")
}

func TestNew(t *testing.T) {
    p := New("TestName", 3)
    if p.Name != "TestName" {
        t.Error("Invalid name for plaet")
    }
    if len(p.Area) != 3 {
        t.Error("Area length invalid")
    }
}
