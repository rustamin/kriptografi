package kriptoalgo

// import "fmt"
// import "reflect"
import "sort"

// var p = fmt.Println



// func main() {
//   var coba string = "coobadibaca"
//   // p(string(coba[1])) // ascii only
//   // p( string( []rune(coba)[1])  ) // utf8 only
//
//   // var newKalimat string
//   //
//   // p(newKalimat)
//
//   p(Scytale("sayarustam"))
//
//   p(len(coba))
//
//   p(Transposisi("rustamin"))
//   p(Caesar("rustamin21"))
// }

func Scytale(kalimat string) string {
  var baris int = 3
  var newKalimat string
  // sayarustam
  // sasm
  // art
  // yua
  for i := 0; i < baris; i++ {
    for j := 0; j < len(kalimat); j += baris {
      if j == 0 {
        newKalimat += string( []rune(kalimat)[i])
      } else {
        if j+i < len(kalimat) {
          newKalimat += string( []rune(kalimat)[j+i])
        }
      }
    }
  }

  return newKalimat
}

func Transposisi(kalimat string) string {
  var kunci string = "gundar"
  // kode ascii dari gundar 103 117 110 100 97 114

  dict := make(map[uint8][]string)

  var temp string

  // looping selama huruf kalimat masih ada
  // looping sepanjang kunci
  for len(kalimat) > 0 {
    for i := 0; i < len(kunci); i++ {
      // harusnya cek dulu masih ada huruf di kalimat atau tidak
      if len(kalimat) > 0 {
        temp = string([]rune(kalimat)[0])
        kalimat = string(kalimat[1:len(kalimat)])
      } else {
        temp = "x"
      }
      dict[kunci[i]] = append(dict[kunci[i]], temp)
    }
  }

  // for k, v := range dict {
  //   fmt.Println("k:", k, "v:", v)
  // }

  // sort key of the dict
  var keys []int
  // get the key first
  for k := range dict {
      keys = append(keys, int(k))
  }

  sort.Ints(keys)

  var chiper string

  for _, k := range keys {
      for _, v := range dict[uint8(k)] {
        chiper += v
      }
  }

  // fmt.Println(chiper)

  return chiper

}


func Caesar(text string) string {
  // shift -> number of letters to move to right or left
  // offset -> size of the alphabet, in this case the plain ASCII
  shift, offset := rune(3), rune(26)

  // string->rune conversion
  runes := []rune(text)

  for index, char := range runes {
    if char >= 'a'+shift && char <= 'z' || char >= 'A'+shift && char <= 'Z' {
      char = char - shift
      } else if char >= 'a' && char < 'a'+shift || char >= 'A' && char < 'A'+shift {
        char = char - shift + offset
      }
      runes[index] = char
  }

  return string(runes)
}
