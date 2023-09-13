package bowling

import (
  "testing"
  // "fmt"
)

var game = Game{
  0,
}

func TestCanRoll(t *testing.T){
  game.Roll(0);
}

func TestGutterGame(t *testing.T){
  for i := 0; i < 20; i++ {
    game.Roll(0);
  }
  if game.Score != 0 {
    t.Errorf("Got %q, wanted %q", game.Score, 0)
  }
}

