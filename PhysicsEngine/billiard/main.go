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
var circle2 *CircleObject
var circle3 *CircleObject
var line1 *LineObject
var gravity float64 = 0.98
var lineAlpha float64

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
    circleSpeed  float64
    mass         float64
    ctrXAxis     float64
    ctrYAxis     float64
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
        mass:         circleRadius*circleRadius,
        ctrXAxis:     xAxis + circleRadius,
        ctrYAxis:     yAxis + circleRadius,
	}
}

func (c *CircleObject) Update() {
	c.xAxis += c.speedOnX
	c.yAxis += c.speedOnY
	c.circleSpeed = math.Sqrt(c.speedOnX*c.speedOnX + c.speedOnY*c.speedOnY)
    c.ctrXAxis = c.xAxis + c.circleRadius
    c.ctrYAxis = c.yAxis + c.circleRadius
}

func (c *CircleObject) Build() {
	ctr := image.Pt(int(c.xAxis), int(c.yAxis))

	// Draw circle
	center := ctr.Add(image.Pt(int(c.circleRadius), int(c.circleRadius)))

	canvas := g.GetCanvas()
	canvas.AddCircle(center, float32(c.circleRadius), color.RGBA{200, 12, 12, 255}, int(c.circleRadius), 2)
}

func cicleCollisionCheck(c1 *CircleObject, c2 *CircleObject) bool {
    distance := math.Sqrt(math.Pow(c1.ctrYAxis - c2.ctrYAxis,2) + math.Pow(c1.ctrXAxis - c2.ctrXAxis,2))
    fmt.Println(distance)
    if  distance <= (c1.circleRadius + c2.circleRadius) {
        circleCollisionResolve(c1, c2)
        fmt.Println("Circle Collision!")
        return true
    } else {
        fmt.Println("We are cool!")
        return false
    }
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


func circleCollisionResolve(c1 *CircleObject, c2 *CircleObject) {
    c1SpeedOnXTemp := ((c1.mass - c2.mass)*c1.speedOnX + 2 * c2.mass * c2.speedOnX)/(c1.mass+c2.mass)
    c2SpeedOnXTemp := ((c2.mass - c1.mass)*c2.speedOnX + 2 * c1.mass * c1.speedOnX)/(c1.mass+c2.mass)
    c1SpeedOnYTemp := ((c1.mass - c2.mass)*c1.speedOnY + 2 * c2.mass * c2.speedOnY)/(c1.mass+c2.mass)
    c2SpeedOnYTemp := ((c2.mass - c1.mass)*c2.speedOnY + 2 * c1.mass * c1.speedOnY)/(c1.mass+c2.mass)
    c1.speedOnX = c1SpeedOnXTemp
    c2.speedOnX = c2SpeedOnXTemp
    c1.speedOnY = c1SpeedOnYTemp
    c2.speedOnY = c2SpeedOnYTemp
}


func collisionResolve(c *CircleObject, l *LineObject) {
	speedTheta := math.Atan2(c.speedOnY, c.speedOnX)

	c.speedOnX = c.circleSpeed * math.Cos(2*lineAlpha-speedTheta)
	c.speedOnY = c.circleSpeed * math.Sin(2*lineAlpha-speedTheta)

}

func groundCheck(c *CircleObject) bool {
	if c.ctrXAxis > (1200-c.circleRadius) || c.ctrXAxis < (0 + c.circleRadius) {
		groundResolveX(c)
		return true
	}
	if c.ctrYAxis > (900-c.circleRadius) || c.ctrYAxis < (0 + c.circleRadius) {
		groundResolveY(c)
		return true
	} else {
        return false
    }
}

func groundResolveX(c *CircleObject) {
	c.speedOnX = -c.speedOnX
}

func groundResolveY(c *CircleObject) {
	c.speedOnY = -c.speedOnY
}

func addGravity(c *CircleObject) {
	c.speedOnY += gravity
}

func loop() {

	g.SingleWindow().Layout(
		circle1,
        circle2,
        circle3,
		line1,
	)
	circle1.Update()
    circle2.Update()
    circle3.Update()
	line1.Update()
    c1col := groundCheck(circle1)
    c2col := groundCheck(circle2)
    c3col := groundCheck(circle3)
    c1col = c1col || collisionCheck(circle1, line1)
    c2col = c2col || collisionCheck(circle2, line1)
    c3col = c3col || collisionCheck(circle3, line1)
    c1c2col := cicleCollisionCheck(circle1, circle2)
    c1c3col := cicleCollisionCheck(circle1, circle3)
    c2c3col := cicleCollisionCheck(circle2, circle3)
    c1col = c1col || c1c2col
    c1col = c1col || c1c3col
    c2col = c2col || c1c2col
    c2col = c2col || c2c3col
    c3col = c3col || c1c3col
    c3col = c3col || c2c3col
    // if !c1col {
    //     addGravity(circle1)
    //     fmt.Println("circle 1 has gravity")
    // } else {
    //     fmt.Println("circle 1 no gravity")
    // }
    // if !c2col {
    //     addGravity(circle2)
    // }
    // if !c3col {
    //     addGravity(circle3)
    // }
    fmt.Println(c1col)
    fmt.Println(c2col)
    fmt.Println(c3col)
}

func main() {
	circle1 = Circle("Circle", 100, 100, 3, 3, 100)
    circle2 = Circle("Circle", 400, 400, -5, -8, 50)
    circle3 = Circle("Circle", 700, 100, 2, 3, 70)
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
