package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//RandomWalk ... global type
type RandomWalk struct {
	Direction           []string
	DirectionPercentage []float64
	MoveX               []int
	MoveY               []int
	CurrentStateX       int
	CurrentStateY       int
	Route               plotter.XYs
	Moves               []string
	Finish              int
}

//NewRandomWalk .. what to do if there is new rand walk
func NewRandomWalk(percentage []float64, finish int) RandomWalk {
	r := RandomWalk{
		Direction:           []string{"North", "Northeast", "East", "Southeast", "South", "Southwest", "West", "Northwest"},
		DirectionPercentage: percentage,
		MoveX:               []int{0, 1, 1, 1, 0, -1, -1, -1},
		MoveY:               []int{1, 1, 0, -1, -1, -1, 0, 1},
		CurrentStateX:       0,
		CurrentStateY:       0,
		Finish:              finish,
	}

	for !(r.CurrentStateX == r.Finish && r.CurrentStateY == r.Finish) {
		fmt.Println(r.CurrentStateX, r.CurrentStateY)
		r.NextStep()
	}
	fmt.Println("here")
	return r
}

//NextStep .. if doing nextStep
func (r *RandomWalk) NextStep() {
	nextX := r.CurrentStateX
	nextY := r.CurrentStateY
	moves := false
	move := -1
	source := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source)

	for !moves {
		nextX = r.CurrentStateX
		nextY = r.CurrentStateY
		randomNumber := r1.Float64()
		for i, perc := range r.DirectionPercentage {
			if perc > randomNumber {
				move = i
				nextX += r.MoveX[i]
				nextY += r.MoveY[i]
				break
			}
		}
		fmt.Println(nextX, nextY)
		if nextX >= 0 && nextX <= r.Finish && nextY >= 0 && nextY <= r.Finish {
			moves = true
		}
	}
	r.Moves = append(r.Moves, r.Direction[move])
	r.CurrentStateX = nextX
	r.CurrentStateY = nextY
	r.Route = append(r.Route, plotter.XY{X: float64(nextX), Y: float64(nextY)})
}

func main() {
	randomwalk := NewRandomWalk([]float64{0.19, 0.43, 0.60, 0.70, 0.72, 0.75, 0.85, 1.00}, 20)
	fmt.Println(randomwalk.Moves)
	fmt.Println(randomwalk.Route)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	if err := plotutil.AddLinePoints(p, randomwalk.Route); err != nil {
		log.Fatal(err.Error())
	}

	p.X.Max = 20
	p.Y.Max = 20

	//make a plot to png file
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "randomwalk_result.png"); err != nil {
		panic(err)
	}

	//duration delay for 10 seconds when open the .exe file
	fmt.Println("It will close in 60 seconds automatically.")
	fmt.Println("After you run it. Open the plot in randomwalk_result.png (same directory)")
	fmt.Println("- Yusron Hanan Z.V.I (1301184476) - IF-42-INT")
	duration := time.Duration(60) * time.Second
	time.Sleep(duration)
}
