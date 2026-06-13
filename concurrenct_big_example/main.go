package main

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

type Data struct {
	Value int
	X     float64
	Y     float64
	Type  string
}

func round(x float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return math.Round(x*pow) / pow
}

func TrackAirHumidity(ctx context.Context, ch chan Data) {
	for {
		n := 1 + rand.IntN(10)

		select {
		case <-ctx.Done():
			return

		case <-time.After(time.Second * time.Duration(n)):
		}

		x := round(rand.Float64()*10, 2)
		y := round(rand.Float64()*10, 2)

		humidity := 30 + rand.IntN(40)
		data := Data{humidity, x, y, "Humidity"}

		select {
		case <-ctx.Done():
			return

		case ch <- data:
		}
	}
}
func TrackAirPressure(ctx context.Context, ch chan Data) {
	for {
		n := 1 + rand.IntN(10)

		select {
		case <-ctx.Done():
			return

		case <-time.After(time.Second * time.Duration(n)):
		}

		x := round(rand.Float64()*10, 2)
		y := round(rand.Float64()*10, 2)

		pressure := 30 + rand.IntN(40)
		data := Data{pressure, x, y, "Pressure"}

		select {
		case <-ctx.Done():
			return

		case ch <- data:
		}
	}
}

func TrackActivity(ctx context.Context, ch chan Data) {
	for {
		n := 1 + rand.IntN(10)

		select {
		case <-ctx.Done():
			return

		case <-time.After(time.Second * time.Duration(n)):
		}

		x := round(rand.Float64()*10, 2)
		y := round(rand.Float64()*10, 2)

		activity := 30 + rand.IntN(40)
		data := Data{activity, x, y, "Activity"}

		select {
		case <-ctx.Done():
			return

		case ch <- data:
		}
	}
}

func start(
	task func(context.Context, chan Data),
) (chan Data, context.CancelFunc) {

	n := 1 + rand.IntN(5)

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan Data)

	var wg sync.WaitGroup

	for range n {
		wg.Go(func() {
			task(ctx, ch)
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch, cancel
}

func main() {
	pressureCh, pressureCancel := start(TrackAirPressure)
	humidityCh, humidityCancel := start(TrackAirHumidity)
	activityCh, activityCancel := start(TrackActivity)
	go func() {
		n := 5 + rand.IntN(10)
		fmt.Println("After", n, "seconds pressure channel will be closed")
		time.Sleep(time.Second * time.Duration(n))
		pressureCancel()
		fmt.Println("Pressure channel is closed")
	}()
	go func() {
		n := 5 + rand.IntN(10)
		fmt.Println("After", n, "seconds humidity channel will be closed")
		time.Sleep(time.Second * time.Duration(n))
		humidityCancel()
		fmt.Println("Humidity channel is closed")
	}()
	go func() {
		n := 5 + rand.IntN(10)
		fmt.Println("After", n, "seconds activity channel will be closed")
		time.Sleep(time.Second * time.Duration(n))
		activityCancel()
		fmt.Println("Activity channel is closed")
	}()
	alive := 3
	for alive > 0 {
	select {
	case data, ok := <-pressureCh:
		if !ok {
			pressureCh = nil
			alive--
			continue
		}
		pp.Println(data)

	case data, ok := <-humidityCh:
		if !ok {
			humidityCh = nil
			alive--
			continue
		}
		pp.Println(data)

	case data, ok := <-activityCh:
		if !ok {
			activityCh = nil
			alive--
			continue
		}
		pp.Println(data)
	}
}
}
