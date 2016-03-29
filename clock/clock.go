// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "strconv"

const testVersion = 3

//Clock soporta la hroa y los minutos
type Clock struct {
    h,  m int
}
// Complete the type definition.  Pick a suitable data type.

//New - asignar a struct
func New(hour, minute int) Clock {
   //for negative minutes  , add 60 minutes and substract 1 
   for minute < 0 {
       minute += 60
       hour --
   } 
   //for minutes  , substract 60 minutes and add 1    
   for minute > 59 {
    minute -= 60
    hour++
   }
   
   //for negative hours, just make positive   
   for hour < 0 {
       hour +=  24  
    }
   hour = hour % 24
        
   return Clock {hour,minute}
    
}

func (c Clock) String() string {

   resultado :=    strconv.Itoa(c.h)
   if len(resultado) == 1 {
   resultado = "0" + resultado
   }
   resultado2 :=    strconv.Itoa(c.m)
   if len(resultado2) == 1 {
   resultado2 = "0" + resultado2
   }
    return resultado + ":" + resultado2
}

//Clock.add added minutes 
func (c Clock) Add(minutes int) Clock {
  //just add the minutes
   c.m += minutes
    
   for c.m < 0 {
       c.m += 60
       c.h --
   } 
      
   for c.m > 59 {
    c.m -= 60
    c.h++
   }
   
   for c.h < 0 {
       c.h +=  24  
    }
    c.h = c.h % 24
    return c
}

