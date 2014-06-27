package color_analyzer

import "image"
import "image/color"

func Analyze(img image.Image) color.Color {
    /* return img.At(2,2) */
    /* return color.RGBA{0,0,0,0} */
    return mostUsedColor(img)
}

func mostUsedColor(img image.Image) color.Color {
    counts := make(map[color.Color]int)
    var bestColor color.Color
    bestCount := 0

    b := img.Bounds()
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {

            currColor := img.At(x,y)
            counts[currColor]++
            currCount := counts[currColor]
            if currCount > bestCount {
                bestCount = currCount
                bestColor = currColor
            }
        }
    }
    return bestColor
}
