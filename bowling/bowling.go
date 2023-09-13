package bowling

type Gamer interface{
  Role()
}

type Game struct {
  Score int
}

func (g Game) Roll(pins int) {

}

