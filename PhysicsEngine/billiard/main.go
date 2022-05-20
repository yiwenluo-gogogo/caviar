package main

import (
    "fmt"
    "image"
    "image/color"
    // "reflect"

    g "github.com/AllenDang/giu"
)

type CircleButtonWidget struct {
    id      string
    clicked func()
}

func CircleButton(id string, clicked func()) *CircleButtonWidget {
    return &CircleButtonWidget{
        id:      id,
        clicked: clicked,
    }
}

func (c *CircleButtonWidget) Build() {
    width, _ := g.CalcTextSize(c.id)
    var padding float32 = 12.0

    // pos := g.GetCursorPos()
    // fmt.Println("pos is: ", pos)
    // fmt.Println(reflect.TypeOf(pos.X))
    ctr := image.Pt(300, 20)


    // Calcuate the center point
    radius := int(width/2 + padding*2)

    // Place a invisible button to be a placeholder for events
    buttonWidth := float32(radius) * 2
    g.InvisibleButton().Size(buttonWidth, buttonWidth).OnClick(c.clicked).Build()

    // If button is hovered
    drawActive := g.IsItemHovered()

    // Draw circle
    center := ctr.Add(image.Pt(radius, radius))

    canvas := g.GetCanvas()
    if drawActive {
        canvas.AddCircleFilled(center, float32(radius), color.RGBA{12, 12, 200, 255})
    }
    canvas.AddCircle(center, float32(radius), color.RGBA{200, 12, 12, 255}, radius, 2)

    // Draw text
    // canvas.AddText(center.Sub(image.Pt(int((width)/2), int(height/2))), color.RGBA{255, 255, 255, 255}, c.id)
}


func onCircleButton() {
    fmt.Println("Circle Button")
}

func loop() {
    g.SingleWindow().Layout(
        CircleButton("Circle Button", onCircleButton),
    )
}

func main() {
    wnd := g.NewMasterWindow("Custom Widget", 1200, 900, g.MasterWindowFlagsNotResizable)
    wnd.Run(loop)
}