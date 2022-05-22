package main

import (
    "fmt"
    "image"
    "image/color"
    "time"
    "math"
    g "github.com/AllenDang/giu"
)


var circle1 *CircleObject
var line1 *LineObject

type LineObject struct {
    id string
    startX float64
    startY float64
    endX float64
    endY float64
    beta float64
    constant float64
}


type CircleObject struct {
    id      string
    xAxis      float64
    yAxis      float64
    speedOnX   float64
    speedOnY   float64
    circleRadius     float64
}


func Line(id string, startX float64, startY float64, endX float64, endY float64) *LineObject {
    beta := (endY - startY)/(endX-startX)
    constant := startY - startX * beta


    return &LineObject{
        id:      id,
        startX:   startX,
        startY:   startY,
        endX: endX,
        endY: endY,
        beta: beta,
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
        id:      id,
        xAxis:   xAxis,
        yAxis:   yAxis,
        speedOnX: speedOnX,
        speedOnY: speedOnY,
        circleRadius: circleRadius,
    }
}

func (c *CircleObject) Update() {
    c.xAxis += c.speedOnX
    c.yAxis += c.speedOnY
    if c.xAxis > 1200 - 2 * c.circleRadius || c.xAxis < 0  {
        c.speedOnX = -c.speedOnX
    }
    if c.yAxis > 900 - 2 * c.circleRadius || c.yAxis < 0  {
        c.speedOnY = -c.speedOnY
    }
}

func (c *CircleObject) Build() {
    ctr := image.Pt(int(c.xAxis), int(c.yAxis))

    // Draw circle
    center := ctr.Add(image.Pt(int(c.circleRadius), int(c.circleRadius)))

    canvas := g.GetCanvas()
    canvas.AddCircle(center, float32(c.circleRadius), color.RGBA{200, 12, 12, 255}, int(c.circleRadius), 2)
}

func collisionCheck(c *CircleObject, l *LineObject) {
    d  := math.Abs(l.beta * (c.xAxis + c.circleRadius) - (c.yAxis + c.circleRadius) + l.constant)/math.Sqrt(l.beta * l.beta + 1)
    if d < c.circleRadius {
        fmt.Println("Collision!")
    } else {
        fmt.Println("We are cool!")
    }

}


func loop() {
    
    g.SingleWindow().Layout(
        circle1,
        line1,
    )
    circle1.Update()
    line1.Update()
    collisionCheck(circle1, line1)
}

func main() {
    circle1 = Circle("Circle",100,100,1,1, 100)
    line1 = Line("Line",0,800,1200,500)

    wnd := g.NewMasterWindow("Custom Widget", 1200, 900, 0)

    go func() {
        for {
            time.Sleep(time.Duration(1 * time.Millisecond))
            g.Update()
        }
    }()

    wnd.Run(loop)
}