package kriptoalgo

import "testing"

func TestScytale(t *testing.T) {

  var kalimat string = "rustamin"
  var cipher string = Scytale(kalimat)

  // rustamin
  // r t i
  // u a n
  // s m

  if cipher != "rtiuansm" {
    t.Error("Expected rtiuansm, got ", cipher)
  }

}


func TestTransposisi(t *testing.T) {

  var kalimat string = "rtiuansm"
  var cipher string = Transposisi(kalimat)

  if cipher != "axuxrsixnxtm" {
    t.Error("Expected axuxrsixnxtm, got ", cipher)
  }

}

func TestCaesar(t *testing.T) {

  var kalimat string = "axuxrsixnxtm"
  var cipher string = Caesar(kalimat)

  if cipher != "xuruopfukuqj" {
    t.Error("Expected xuruopfukuqj, got ", cipher)
  }

}
