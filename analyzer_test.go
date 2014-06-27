package color_analyzer

import (
    "testing"

    "os"
    "bufio"
    "image"
    "image/color"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

    "fmt"
)
var (
    imageFile = "images/blue.jpg"
    blue  color.Color = color.RGBA{0, 0, 255, 255}
    tolerance = 2000
    colors = map[string]color.Color{
        "images/blue.jpg": color.RGBA{0, 0, 255, 255},
        "images/kite_runner.jpeg": color.RGBA{143, 165, 116, 255},
        "images/ruby.jpeg": color.RGBA{243, 243, 219, 255},
    }
)

func TestColors(t *testing.T){
    for file, color := range colors {
        fmt.Printf("\nTesting file: %v \n", file)
        if !checkImage(file, color){
            t.Error("Color of file not within tolerance: " + file)
        }
    }
}

func checkImage(file string, expectedColor color.Color) bool{
    f, err := os.Open(file)
    if err != nil {
        panic(err.Error())
    }

    defer f.Close()
    img, typeString , err := image.Decode(bufio.NewReader(f))
    fmt.Println(typeString)
    if err != nil {
        panic(err.Error())
    }
    guessedColor := Analyze(img)
    return withinTolerance(expectedColor, guessedColor, tolerance)
}

func withinTolerance(c0, c1 color.Color, tolerance int) bool {
    r0, g0, b0, a0 := c0.RGBA()
    r1, g1, b1, a1 := c1.RGBA()
    r := delta(r0, r1)
    g := delta(g0, g1)
    b := delta(b0, b1)
    a := delta(a0, a1)

    fmt.Printf("C0: r:%v, g:%v, b:%v, a:%v\n", r0,g0,b0,a0)
    fmt.Printf("C1: r:%v, g:%v, b:%v, a:%v\n", r1,g1,b1,a1)
    fmt.Printf("Deltas: r:%v, g:%v, b:%v, a:%v\n", r,g,b,a)

    return r <= tolerance && g <= tolerance && b <= tolerance && a <= tolerance
}

func delta(u0, u1 uint32) int {
    d := int(u0) - int(u1)
    if d < 0 {
        return -d
    }
    return d
}
