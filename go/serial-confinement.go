package bakery

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"

		// baker never touches `*cake` again.
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"

		// icer never touches `*cake` again.
		iced <- caked
	}
}
