package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"

	g "github.com/AllenDang/giu"
)

var circle1 *CircleObject
var line1 *LineObject
var gravity float64 = 0.98
var lineAlpha float64
var circleSpeed float64

type LineObject struct {
	id       string
	startX   float64
	startY   float64
	endX     float64
	endY     float64
	beta     float64
	constant float64
}

type CircleObject struct {
	id           string
	xAxis        float64
	yAxis        float64
	speedOnX     float64
	speedOnY     float64
	circleRadius float64
}

func Line(id string, startX float64, startY float64, endX float64, endY float64) *LineObject {
	beta := (endY - startY) / (endX - startX)
	constant := startY - startX*beta

	return &LineObject{
		id:       id,
		startX:   startX,
		startY:   startY,
		endX:     endX,
		endY:     endY,
		beta:     beta,
		constant: constant,
	}
}

func (l *LineObject) Update() {
	// fmt.Println("I am static")
}

func (l *LineObject) Build() {

	startPoint := image.Pt(int(l.startX), int(l.startY))
	endPoint := image.Pt(int(l.endX), int(l.endY))

	// Draw line

	canvas := g.GetCanvas()
	canvas.AddLine(startPoint, endPoint, color.RGBA{200, 12, 12, 255}, 2)

}

func Circle(id string, xAxis float64, yAxis float64, speedOnX float64, speedOnY float64, circleRadius float64) *CircleObject {
	return &CircleObject{
		id:           id,
		xAxis:        xAxis,
		yAxis:        yAxis,
		speedOnX:     speedOnX,
		speedOnY:     speedOnY,
		circleRadius: circleRadius,
	}
}

func (c *CircleObject) Update() {
	c.xAxis += c.speedOnX
	c.yAxis += c.speedOnY
	if c.xAxis > 1200-2*c.circleRadius || c.xAxis < 0 {
		c.speedOnX = -c.speedOnX
	}
	if c.yAxis > 900-2*c.circleRadius || c.yAxis < 0 {
		c.speedOnY = -c.speedOnY
	} else {
		c.speedOnY += gravity
	}
	circleSpeed = math.Sqrt(c.speedOnX*c.speedOnX + c.speedOnY*c.speedOnY)
}

func (c *CircleObject) Build() {
	ctr := image.Pt(int(c.xAxis), int(c.yAxis))

	// Draw circle
	center := ctr.Add(image.Pt(int(c.circleRadius), int(c.circleRadius)))

	canvas := g.GetCanvas()
	canvas.AddCircle(center, float32(c.circleRadius), color.RGBA{200, 12, 12, 255}, int(c.circleRadius), 2)
}

func collisionCheck(c *CircleObject, l *LineObject) bool {
	d := math.Abs(l.beta*(c.xAxis+c.circleRadius)-(c.yAxis+c.circleRadius)+l.constant) / math.Sqrt(l.beta*l.beta+1)
	if d < c.circleRadius {
		fmt.Println("Collision!")
		collisionResolve(c, l)
		return true
	} else {
		fmt.Println("We are cool!")
		return false
	}
}

func collisionResolve(c *CircleObject, l *LineObject) {
	speedTheta := math.Atan2(c.speedOnY, c.speedOnX)

	c.speedOnX = circleSpeed * math.Cos(2*lineAlpha-speedTheta)
	c.speedOnY = circleSpeed * math.Sin(2*lineAlpha-speedTheta)

}

// func groundCheck(c *CircleObject, l *LineObject) bool {
// 	if c.xAxis > 1200-2*c.circleRadius || c.xAxis < 0 {
// 		groundResolveX(c)
// 		return true
// 	}
// 	if c.yAxis > 900-2*c.circleRadius || c.yAxis < 0 {
// 		groundResolveY(c)
//         c.speedOnY += gravity
// 		return true
// 	}
// 	return false
// }

// func groundResolveX(c *CircleObject) {
// 	c.speedOnX = -c.speedOnX
// 	circleSpeed = math.Sqrt(c.speedOnX*c.speedOnX + c.speedOnY*c.speedOnY)
// }

// func groundResolveY(c *CircleObject) {
// 	c.speedOnY = -c.speedOnY
// 	circleSpeed = math.Sqrt(c.speedOnX*c.speedOnX + c.speedOnY*c.speedOnY)
// }

// func addGravity(c *CircleObject) {
// 	c.speedOnY += gravity
// }

func loop() {

	g.SingleWindow().Layout(
		circle1,
		line1,
	)
	circle1.Update()
	line1.Update()
	// groundCheck(circle1, line1)
	collisionCheck(circle1, line1)
	// groundCollided := groundCheck(circle1, line1)
	// if !groundCollided {
	//     fmt.Println("Gound not colided")
	// 	lopeCollided := collisionCheck(circle1, line1)
	// 	if !lopeCollided {
	//         fmt.Println("lope not colided")
	// 		addGravity(circle1)
	// 	}
	// }
}

func main() {
	circle1 = Circle("Circle", 100, 100, 1, 2, 100)
	line1 = Line("Line", 0, 800, 1200, 500)

	lineAlpha = math.Atan2((line1.endY - line1.startY), (line1.endX - line1.startX))

	wnd := g.NewMasterWindow("Custom Widget", 1200, 900, 0)

	go func() {
		for {
			time.Sleep(time.Duration(1 * time.Millisecond))
			g.Update()
		}
	}()

	wnd.Run(loop)
}
