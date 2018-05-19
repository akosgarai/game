package area

import(
    "testing"
)

func TestNew(t *testing.T) {
    area := New()
    if area.Type != "grassland" {
        t.Error("It should be grassland!")
    }
}
