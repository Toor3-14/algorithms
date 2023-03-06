package abc

/**
* Default alphabet of priority (the smaller the array cell, the more priority) 
* The alphabet can also take a string of two or more characters. This means they have the same priority.
*/
var ABC = []string{"A","B","C","D","E","F","G",
                   "H","I","J","K","L","M","N",
                   "O","P","Q","R","S","T","U",
                   "V","W","X","Y","Z",

                   "a","b","c","d","e","f","g",
                   "h","i","j","k","l","m","n",
                   "o","p","q","r","s","t","u",
                   "v","w","x","y","z"}
                  

func GetPosition(char string, abc []string) int {
  for i := 0; i < len(abc); i++  {
     if(len(abc[i]) > 1) {
      for j := 0; j < len(abc[i]); j++  {
        if(char == abc[i][j:j+1]) {
          return i
        }
      }
    }else {
      if(char == abc[i][0:1]) {
        return i
      }
    }
  }
  return len(abc)
}