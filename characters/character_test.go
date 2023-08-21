package characters

import (
	"go-rpg/proto"
	"testing"
)

func TestNewChar(t *testing.T) {
	want := Character{
		Character: &proto.Character{
			Name:   "teste",
			Race:   0,
			Class:  0,
			Health: 100,
		},
	}
	got := new(&proto.Character{
		Name:   "teste1",
		Race:   0,
		Class:  0,
		Health: 100,
	})
	if got != want {
		t.Logf("working right: different data")
	}
}

func BenchmarkNewChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		new(&proto.Character{
			Name:   "teste",
			Race:   0,
			Class:  0,
			Health: 100,
		})
	}
}
