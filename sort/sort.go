package sort 

import "gihub.com/Toor3-14/algorithms/abc"

type gen interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}


func Buble[T gen](array []T) {
  for i := 0; i < len(array)-1; i++ {
     if ((array)[i] > (array)[i+1]) {
       (array)[i], (array)[i+1] = (array)[i+1], (array)[i]
       i=-1;
     }
   }
}

func BubleCustom(array []string) {
  for i := 0; i < len(array)-1; i++ {
     if abc.GetPosition((array)[i][0:1],abc.ABC) > abc.GetPosition((array)[i+1][0:1],abc.ABC)  {
      (array)[i], (array)[i+1] = (array)[i+1], (array)[i]
       i=-1;
     }else if abc.GetPosition((array)[i][0:1],abc.ABC) == abc.GetPosition((array)[i+1][0:1],abc.ABC) {

        var min int = 0
        if len((array)[i]) > len((array)[i+1]) {
          min = len((array)[i+1])
        }else {
          min = len((array)[i])
        }
        for j := 1; j < min-1; j++ {
          if abc.GetPosition((array)[i][j:j+1],abc.ABC) > abc.GetPosition((array)[i+1][j:j+1],abc.ABC)  {
            (array)[i], (array)[i+1] = (array)[i+1], (array)[i]
            i=0;
          }else if abc.GetPosition((array)[i][j:j+1],abc.ABC) == abc.GetPosition((array)[i+1][j:j+1],abc.ABC) {
            continue
          }else {
            break
          }
        }

     }
   }
}

