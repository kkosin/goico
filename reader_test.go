package ico

import (
	"bytes"
	"io/ioutil"
	"testing"
)

const testImage = "/vagrant_data/text.ico"

func TestDecodeAll(t *testing.T) {
	data, err := ioutil.ReadFile(testImage)
	if err != nil {
		t.Error(err)
	}
	r := bytes.NewReader(data)
	_, err = DecodeAll(r)
	if err != nil {
		t.Error(err)
	}
}
