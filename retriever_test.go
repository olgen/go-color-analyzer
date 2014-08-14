package main

import (
    "testing"
    "fmt"
)


var (
    imageUrl = "http://cdn3.txtr.com/delivery/img?type=DOCUMENTIMAGE&documentID=qmtsu89&size=LARGE"
)
func TestHttpRetrieval(t *testing.T) {
    fmt.Println("Retrieving image from url: ", imageUrl)
    img, err:= RetrieveImageFromUrl(imageUrl)
    if err!=nil {
            panic(err)
    }

    if !( ( *img ).Bounds().Dx() == 320 ) {
            t.Error("Wrong image width, should == 320")
    }

}
