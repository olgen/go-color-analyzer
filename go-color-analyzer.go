package main

import (
    "net/http"
    "os"
    "log"
    "fmt"

    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "image/color"
    "bufio"
)

const (
    DarkThreshold =    5 / 255.0 * 65535.0
    LightThreshold = 250 / 255.0 * 65535.0
)
func main(){
    http.HandleFunc("/color", httpHandler )

    port := portSetting()
    log.Printf("Handler listening on port:%v", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func httpHandler(writer http.ResponseWriter, req *http.Request){
    url := req.URL.Query().Get("url")
    if url != "" {
        handleUrl(writer, url)
    } else {
        http.Error(writer, "No url param given!", 422)
    }
}

func handleUrl(writer http.ResponseWriter, url string) {
    img, err := RetrieveImageFromUrl(url)
    if err != nil {
        log.Printf(err.Error())
        http.Error(writer, err.Error(), 500)
        return
    }

    if img == nil {
        msg := "Image is nil!"
        log.Printf(msg)
        http.Error(writer, msg, 500)
        return
    }

    color := Analyze(img)
    hexColor := HexColor(color)
    fmt.Println("Got color: %v", hexColor)
    fmt.Fprintf(writer, "%v", hexColor)
}

func portSetting() string {
    port := os.Getenv("PORT")
    if port == "" {
        panic("No PORT env-var given!")
    }
    return ":" + port
}

func RetrieveImageFromUrl(url string) (image.Image, error) {
    log.Printf("Getting image from url=%v", url)
    response, err := http.Get(url)
    log.Printf("Status: %v", response.Status)
    log.Printf("Encoding: %v", response.TransferEncoding)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    img, imgType, err := image.Decode(bufio.NewReader(response.Body))
    log.Printf("ImageType: %v", imgType)

    return img, nil
}

func HexColor(col color.Color) string {
    r,g,b,_ := col.RGBA()
    log.Printf("Color.RGBA= %v", col)
    hex := fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
    log.Printf("Color.HEX= %v", hex)
    return hex
}

func Analyze(img image.Image) color.Color {
    return mostUsedColor(img)
}

func mostUsedColor(img image.Image) color.Color {
    histogram := make(map[color.Color]int)
    var bestColor color.Color
    bestCount := 0

    bounds := img.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {

            currColor := img.At(x,y)

            histogram[currColor]++
            currCount := histogram[currColor]
            if currCount > bestCount && !ignoreColor(currColor){
                bestCount = currCount
                bestColor = currColor
            }
        }
    }
    return bestColor
}

func ignoreColor(col color.Color) bool {
    return tooDark(col) || tooLight(col)
}

func tooDark(col color.Color) bool {
    r,g,b,_ := col.RGBA()
    return r < DarkThreshold && g < DarkThreshold && b < DarkThreshold
}

func tooLight(col color.Color) bool {
    r,g,b,_ := col.RGBA()
    return r > LightThreshold && g > LightThreshold && b > LightThreshold
}
