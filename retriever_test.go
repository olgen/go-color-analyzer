package main

import (
    "testing"
)


var (
    imageUrl = "http://cdn3.txtr.com/delivery/img?type=DOCUMENTIMAGE&documentID=qmtsu89&size=LARGE"
)
func TestHttpRetrieval(t *testing.T) {
    image, err:= RetrieveImageFromUrl(imageUrl)
    if err!=nil {
            panic(err)
    }

    if image == nil {
        panic( "Image is nil!")
    }

    if !(  image.Bounds().Dx() == 320 ) {
            t.Error("Wrong image width, should == 320")
    }

}
