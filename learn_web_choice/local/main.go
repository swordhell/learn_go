package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	width  = 40
	height = 20
)

var (
	segments = []string{"奖品 A", "奖品 B", "奖品 C", "奖品 D", "谢谢参与"}
	angle    = 0.0
	speed    = 5.0
)

func drawRoulette() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawWheel()
	drawPointer()
	termbox.Flush()
}

func drawWheel() {
	for i, segment := range segments {
		drawSegment(i, segment)
	}
}

func drawSegment(index int, text string) {
	anglePerSegment := 360.0 / float64(len(segments))
	startAngle := angle - anglePerSegment/2
	endAngle := startAngle + anglePerSegment

	centerX := width / 2
	centerY := height / 2
	radius := 10.0

	for theta := startAngle; theta <= endAngle; theta += 0.1 {
		x := int(float64(centerX) + radius*cos(theta))
		y := int(float64(centerY) + radius*sin(theta))
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}

	// Draw the text in the center of the segment
	textX := int(float64(centerX) + radius*cos(startAngle+anglePerSegment/2))
	textY := int(float64(centerY) + radius*sin(startAngle+anglePerSegment/2))
	printText(textX, textY, text)
}

func drawPointer() {
	centerX := width / 2
	centerY := height / 2
	pointerLength := 5.0

	pointerX := int(float64(centerX) + pointerLength*cos(angle))
	pointerY := int(float64(centerY) + pointerLength*sin(angle))

	termbox.SetCell(pointerX, pointerY, '↑', termbox.ColorRed, termbox.ColorDefault)
}

func printText(x, y int, text string) {
	for i, char := range text {
		termbox.SetCell(x+i, y, char, termbox.ColorWhite, termbox.ColorDefault)
	}
}

func cos(theta float64) float64 {
	return float64(width) * 0.5 * (1 - math.Cos(theta*math.Pi/180))
}

func sin(theta float64) float64 {
	return float64(height) * 0.5 * math.Sin(theta*math.Pi/180)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	quit := false
	for !quit {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					quit = true
				case termbox.KeyEnter:
					spinRoulette()
					drawRoulette()
					awardPrize()
				}
			}
		default:
			angle += speed
			if angle >= 360 {
				angle -= 360
			}
			drawRoulette()
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func spinRoulette() {
	rand.Seed(time.Now().UnixNano())
	angle = rand.Float64() * 360
}

func awardPrize() {
	time.Sleep(2 * time.Second) // Simulate the spinning
	selectedIndex := int(angle) / (360 / len(segments))
	fmt.Println("恭喜你获得:", segments[selectedIndex])
}
