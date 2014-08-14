package main

import (
    "net/http"
    "os"
    "log"
    "fmt"

    "image"
    "image/color"
    "bufio"

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
    log.Printf("Got url: %v", url)
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
    fmt.Println("Got color: %v", color)
    /* fmt.Fprintf(writer, "Color=%v", color) */
}

func portSetting() string {
    port := os.Getenv("PORT")
    if port == "" {
        panic("No PORT env-var given!")
    }
    return ":" + port
}

func RetrieveImageFromUrl(url string) (*image.Image, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    img, _, err := image.Decode(bufio.NewReader(response.Body))
    return &img, nil
}

func Analyze(img *image.Image) color.Color {
    /* return img.At(2,2) */
    /* return color.RGBA{0,0,0,0} */
    return mostUsedColor(img)
}

func mostUsedColor(img *image.Image) color.Color {
    counts := make(map[color.Color]int)
    var bestColor color.Color
    bestCount := 0

    b := (*img).Bounds()
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {

            currColor := (*img).At(x,y)
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
