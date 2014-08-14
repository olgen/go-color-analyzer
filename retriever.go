package color_analyzer

import (
    "net/http"
    "image"
    "bufio"
)

func RetrieveImageFromUrl(url string) (*image.Image, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    img, _, err := image.Decode(bufio.NewReader(response.Body))
    return &img, nil
}
