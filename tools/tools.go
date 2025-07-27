package tools

import (
        "fmt"
        "bufio"
        "strconv"
        "os"
        )//java.util.Scanner;




func GetMonth(m int) string{
	month:= []string {"jan","feb","mar","apr","may","jun","jul","aug","sep","oct","nov","dec"};
	sm:="unknoun"
	if m>-1 && m <12 {
		sm=month[m]
	}
	return sm;
}

func Menu(list []string) int {
	b:=true
	var x int
	for b{
	for i:=0;i<len(list);i++ {
				fmt.Println(list[i]);
			}
            x=ReadInt(); 
            if x>0 && x <len(list){
				b=false
				}
            }           
          return x;
          }
		  
func ToInt(st string) int{
	var x int ;
	x,err := strconv.Atoi(st)
	if err !=nil{
	
                fmt.Println("incorrect data");
                x=-1
            }
	return x;			
}
func ReadInt() int{
	var e int
	b:=true
	for b{
		_,err:=fmt.Scan(&e)
		if err ==nil {
			b=false
			} else {
				fmt.Println("Error! Please enter digit ")
			}
	}
	return e
}


func ToFloat(st string) float64{
	var  x float64 =-1.0;
	x,err := strconv.ParseFloat(st, 64)
	if err !=nil{
	
                fmt.Println("incorrect data");
                x=-1
            }
	return x;		
}


func ReadFloat() float64 {
	var d float64
	b:=true
	for b{
		_,err:=fmt.Scan(&d)
		if err ==nil {
			b=false
			} else {
				fmt.Println("Error! Please enter digit ")
			}
	}
	return d;
}



func St() string{
    var s string	
	scanner := bufio.NewScanner(os.Stdin) 
    _ = scanner.Scan()
    s= scanner.Text()
    return s
}




func GetWork(gr int) string{
	   work:="Unknown";
	   switch (gr) {
		  case 0:
		     work="day off";
          case 1: 
             work="09-18"; 
          case 2:
             work="13-22"; 
          case 3: 
             work="09-17"; 
          case 4: 
             work="13-21"; 
          case 5:
             work="business trip"; 	  
	   }
    return work;
}

//returrn day of week firstDay -1st day of month day-current day
func GetDay(firstDay,day int) string{
	day=(day+firstDay)-1;
	if day>6 {
		 day=day%7;
	 }
    dayWeek:="Unknown";
	switch (day){
	   case 1: 
		dayWeek="Sun";
       case 2: 
          dayWeek="Mon"
       case 3:
		  dayWeek="Tue"
       case 4: 
		  dayWeek="Wed"
       case 5:
		  dayWeek="Thu"
       case 6: 
          dayWeek="Fri"
       case 0: 
          dayWeek="Sat"		
	  }
     return dayWeek;	  
}

func ExistFile(path string) bool{
	b:=false
	file,err:=os.Open(path)
	if err==nil {
		b=true
		}
	 defer file.Close()   
	 return b	
	}

 func UstToInt(u float64) int{
		i:= int(u*100);
		return i;
		}
		
	func UstToFloat(u int) float64{
		var d float64
		d=(float64(u)/100.0);
		return d;
		}	
		


