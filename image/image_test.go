package image

import (
	"testing"
)

func TestImage(t *testing.T) {

	// Scale a image file by cuting 100*100
	err := ThumbnailF2F("../data/image.png", "../data/image100-100.png", 100, 100)
	if err != nil {
		t.Error("Test ThumbnailF2F:" + err.Error())
	}

	// Scale a image file by cuting width:200 (Equal scaling)
	err = ScaleF2F("../data/image.png", "../data/image200.png", 200)
	if err != nil {
		t.Error("Test ScaleF2F:" + err.Error())
	}

	// File Real name
	filename, err := RealImageName("../data/image.png")
	if err != nil {
		t.Error("Test RealImageName:" + err.Error())
	} else {
		t.Log("Test RealImageName::real filename" + filename)
	}
}
