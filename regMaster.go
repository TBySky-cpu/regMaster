  package main
 import(
    "fmt"
    "log"
    "os"
    "strings"
    "time"
    "strconv"  
    "regMaster/tools"
    "io/ioutil"
    "io/fs"
    "encoding/json"
    "errors"
    )

 type myObject struct{
	Name string;
	Adress string;
	Zone int;
	Ust float64;
	RegInMonth int;
	Index int
	Note  string
	}
	
type dataReg struct {
	Month        int
	FirstDay     int
	Year         int
	TotalObjects int
	MyObjects    []myObject
	Grafik [32]  int
	Reglament    [][]int
	}	



 var  day, 
      year,monthN,
	  modeMenu,
	  modeComand,
	  mode int
var   arhive bool
var	  fileData, //= "data.dpm";
	  filesDir,//,="arhive";
	  fileArhive string

var   calendar time.Time
var   gData dataReg	  
 
	
	
 func main(){
	  modeMenu=1;
	  modeComand=2;
	  mode=modeComand;
	 calendar=time.Now()
	  fileData="data.bin"
	  filesDir="arhive"
	  day=calendar.Day()
	  monthN= int(calendar.Month())
	  year=calendar.Year()
	  initr()  
	  typeLine();         
	  fmt.Println("*      RegMaster v.7.0      *");
	  fmt.Println("*Takhir Bairashevski aug2025*");
	  typeLine();
	  fmt.Println("Today ",day,monthN,year);
	  
	  if gData.Month+1 != monthN {
		   typeLine()
		   fmt.Println ("DATABASE IS OUT OF DATE !!");
		   typeLine()
		   }
	  regByDay(day);
	  typeLine();
	  if mode==modeMenu {
		fmt.Println("1-menu  0-exit"); 	   
	    s:=tools.ReadInt();
        if s==1 {
			fmt.Println("This version does not support menu mode.")
			}
		}
       if mode==modeComand{ 
		   comandLine()
		   } 
         }
	 
	 
func initr(){
	  
	  v:=0;
	  exists := tools.ExistFile(fileData);
      if (!exists) {
		  
          fmt.Println("File doesn't exist!!  "+fileData);
		  fmt.Println("Create ? 1-yes other-no");
		  v=tools.ReadInt();
		  if (v==1) {
			  fmt.Println("enter the num of myObjects  ");
		      gData.TotalObjects =tools.ReadInt();
		      makeArray()
		      setAllObjects();
		      setAllReglament();
		      setAllGrafik();
			  writeData(fileData);
			} else {
				os.Exit(0)
				}
		} else {
		readData(fileData)
		
		}
	}	 
 
func makeArray(){
	gData.MyObjects=make([]myObject,gData.TotalObjects)
	gData.Reglament=make([][]int,gData.TotalObjects)
	for i:=0;i<len(gData.MyObjects);i++{
		gData.Reglament[i]=make([]int,32)
	 }
	
}




 func readData(fileData string) {
	 file, err := os.Open(fileData)
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
    bytes, err := os.ReadFile(fileData);
    if err != nil {
     log.Fatal(err);
     } else {
		  _=json.Unmarshal(bytes, &gData)
		  
		 }
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

  func typeAll(){
		  
		  typeLine();
		  sum:=0.0;
		  for i:=0;i<len(gData.MyObjects);i++ {
			u:=fmt.Sprintf("%.2f",gData.MyObjects[i].Ust)  
			fmt.Println(i,gData.MyObjects[i].Name," ",gData.MyObjects[i].Adress," Z",gData.MyObjects[i].Zone," U",u);
			regByObject(i);
			sum+=gData.MyObjects[i].Ust;
			typeLine();		
          }
	    u:=fmt.Sprintf("%.2f",sum) 
	    fmt.Println("total usl ust ",u);
	}
	

func typeLine(){
		fmt.Println("-----------------------------");
    }	
					
 func regByObject(index int){
	  exist:=false;
		  fmt.Println("reglaments: ");
		  for d:=0;d<len(gData.Reglament[index]);d++ { 
			if (gData.Reglament[index][d]==1){ 
	            fmt.Print(d,",");
				exist=true;
			}
		}
		  fmt.Println();
		  typeLine();	
		  if !exist {
			  fmt.Println("reglaments not found");
			  }
     }
  

func typeZone(z int){
	 typeLine();
	 fmt.Println("Зона № ",z)
		  sum:=0.0;
		  k:=1;
		  for i:=0;i<len(gData.MyObjects);i++ {
			if gData.MyObjects[i].Zone==z{
			u:=fmt.Sprintf("%.2f",gData.MyObjects[i].Ust)  
			fmt.Println(k,gData.MyObjects[i].Name," ",gData.MyObjects[i].Adress," Z",gData.MyObjects[i].Zone," U",u);
			regByObject(i);
			sum+=gData.MyObjects[i].Ust;
			k++
			typeLine();		
          }
	  }
	    u:=fmt.Sprintf("%.2f",sum) 
	    fmt.Println("total usl ust ",u);
	}
func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
	    
func comandLine(){
		
		str:="";
		str1:="";
		
		
		var d1,d2,ind int
		
		
		var err bool
		    work:=true
		
		fmt.Println("enter command or help");
		for work {
			err=true;
			if arhive {
				typeLine()
				fmt.Println("**working in arhive -",fileArhive)
				fmt.Println("archive for",tools.GetMonth(gData.Month),gData.Year)
				} 
			fmt.Print("> ");
			str1=tools.St();
			str=strings.Trim(str1, " ")
			if str=="" || len(str) ==0 {
				continue
				}
			if !checkComand(str) { //если это не команда
					d, err := strconv.Atoi(str)
					if err==nil{ //  и если это число  
						regByDay(d) // показываем день
						} else { // иначе
							_=search(str,false);// ищем объект
							}
			
				} else {  // если это команда то обрабатываем ее
				
			parts:=strings.Split(str," ")
            if len(parts)==1 {
				switch parts[0]{
					case "exit": {
						err=false;
					    work=false
					    }
					case "notes" : typeNotes()
					               err = false
					case "newmonth": {
						err=false;
					    fmt.Println("all data will be overwritten.Continue ? 1-yes");
		                d:=tools.ReadInt();
		                if (d==1){
							saveArhive();
							setAllGrafik();
							setAllReglament();
							writeData(fileData);
							}
					    
					}
					
					case "q": {
						err=false;
					    work=false
					}
					case "check": {
					 veryfy();
					  err=false;
					}
					case "menu": {
						mode=modeMenu;
						//mainMenu();
						err=false;
						}
					case "add": {
						addObject();
						err=false;
						}
					case "help": {
						typeHelp();
						err=false;
						}
					case  "arhr":  {
						err=false;
					    readArhive();
						arhive=true;
						comandLine();
						arhive=false;
						}
					case "arhs": {
						saveArhive();
						err=false;
						}
					case "grk": {
						typeGrafik();
						err=false;
						}
					case "greg" :{
						typeDays();
						err=false;
						}
					case "gust": {
						typeUst();
						err=false;
						}
					case  "obs": {
						typeAll();
						err=false;
						}
					case  "obsf": {
						writeToFile()
						err=false;
						}	
					case "restore": {
					restore();
					err=false;
					}
				}
				} 
			  
             if len(parts)==2 {
				switch parts[0] {
					
				    case "note" :{
						ind=search(parts[1],true)
						if ind >-1 {
							editNote(ind)
							err=false
							}
						}
				    
				    case "zone": {
						typeZone(tools.ToInt(parts[1]));
						err=false;
						}
			
					case "del" :{
						 ind =search(parts[1],true);
						 if  ind>-1  {
							 delObject(ind);
							 err=false;}
							  }
					}
				} 
			   if len(parts)==3 {
				switch parts[0] {
					case "arh" :{
						err =false
						m:=tools.ToInt(parts[1]);
						y:=tools.ToInt(parts[2]);
						err:=findArhive(m,y)
						if err == nil { 
							arhive=true;
							comandLine();
							arhive=false;
						} else {
							fmt.Println(err)
							readData(fileData)
							} 
						}
					
					case "rep" :{
						d1=tools.ToInt(parts[1]);
						d2=tools.ToInt(parts[2]);
						if d1 >-1 && d2>-1  {
							replaceDays(d1,d2);
							 err=false;
							 } 
					    }
					case  "ed" :{
						switch parts[1] {
							case "ob":{
								if  parts[2]=="all"  {
									setAllObjects();
									err=false;
								} else {
									 ind=search(parts[2],true);
							          if ind>-1 {
										  setObject(ind) 
										  err=false;
							        }	
							     }
						        }
						        
							case "reg": {
								if  parts[2]=="all" {
									fmt.Println("all data will be overwritten.Continue ? 1-yes");
									d:=tools.ReadInt();
									if (d==1){
										setAllReglament();
									}
									err=false;
								} else { 
									ind=search(parts[2],true);
									 if (ind>-1) {
											setReglament(ind);
											err=false;
											}
								}
							}
						    case "gr" : {
								if parts[2]=="all" {
									fmt.Println("all data will be overwritten.Continue ? 1-yes");
									d:=tools.ReadInt();
									if (d==1){
										setAllGrafik();
									}
									err=false;
								} else { 
							         ind=tools.ToInt(parts[2]);
							         if (ind>-1) {
										editOne(ind)
										}
							         err=false;
							         }	
								}
														
							}
					if (arhive){
						 writeData(fileArhive);
						 } else {
							 writeData(fileData);
						 } 	
					}
				}
			
			} 
		if err {
			fmt.Println("error in parameters");
			}	
		} 
	}
	    if !arhive {
			 os.Exit(0);
			 } else {
				 fmt.Println("exit arhive mode");
		         readData(fileData);
		         }
  }
		    
func checkComand(s string) bool {
		b:=false;
		
		 mainComand:= []string{"q","add","check","menu","help","exit","rep","obs","obsf",
			                   "grk","greg","gust","del","arhr","arhs","ed","restore","zone","newmonth","arh","note","notes"};
		parts:=strings.Split(s," ")
		
		s1:=parts[0];
		for i:=0;i<len(mainComand);i++{ 
		    if s1== mainComand[i] {
				b=true;
				}
		     }
		return b;
		}	    	  


func typeHelp(){
		fmt.Println("add -add object");
		fmt.Println("arh <month> <year>  - find and read arhive file");
		fmt.Println("arhr -read arhive file");
		fmt.Println("arhs -save in arhive");
		fmt.Println("check   -data base check");
		fmt.Println("<day> -type data for the day");
		fmt.Println("del <name object> -delete object");
		fmt.Println("help -type this help");
		fmt.Println("menu   -mode with menu");
		fmt.Println("<name object>  - type data for the object");
		fmt.Println("obs  -type all objects");
		fmt.Println("obsf  -type all objects in file");
		fmt.Println("grk-type grafik");
		fmt.Println("greg-type regl and grafik");
		fmt.Println("gust-type grafik and usl ust");
		fmt.Println("ed ob all -edit all objects");
		fmt.Println("ed ob <name object> -edit one object");
		fmt.Println("ed reg all  -edit all reglaments");
		fmt.Println("ed reg <name object> -edit one reglament");
		fmt.Println("ed gr all -edit all grafik");
		fmt.Println("ed gr <day>  -edit one day");
		fmt.Println("newmonth  -entering data for a new month");
		fmt.Println("note <object> - entering a note for an object ")
		fmt.Println("notes - search all notes ")
		fmt.Println("rep <day1> <day2>  -replacing reg from day1 to day2");
		fmt.Println("restore -restoring a database from an archive");
		fmt.Println("zone <number zone> -print zone objects");
}	
	
func typeUst(){
	fmt.Println("type Ust")
	sum:=0.0;
	 min:=100.0;
	 max:=0.0;
	 minDay:=0
	 maxDay:=0
		for i :=1;i<len(gData.Grafik);i++{
			if gData.Grafik[i] !=0 && gData.Grafik[i] !=5 {
			for ob:=0;ob<len(gData.MyObjects);ob++ { 
				if gData.Reglament[ob][i]==1 {
					sum+=gData.MyObjects[ob].Ust/float64(gData.MyObjects[ob].RegInMonth)
					 }
					}
					if sum<min {
						min=sum
						minDay=i
						}
					if sum>max {
						max=sum
						maxDay=i
						} 
					} else{
						 sum=0.0} 
			u:=fmt.Sprintf("%.2f",sum) 			 
			fmt.Println(i,tools.GetDay(gData.FirstDay,i)," ",tools.GetWork(gData.Grafik[i])," U",u);
			sum=0.0;
			}
			
		typeLine();
		fmt.Printf("max ust %.2f in day %v \n",max,maxDay)
		fmt.Printf("min ust %.2f in day %v \n",min,minDay)
		typeLine();	
	
	
	}	


func veryfy(){
	typeLine();
		e:=0;
		for i:=0;i<len(gData.MyObjects);i++ {
		  r:=0;
	      for d:=1;d<len(gData.Reglament[i]); d++ {
				if gData.Reglament[i][d]>0 {
					r++;
					if  gData.Grafik[d]==0  {
						fmt.Println("day ",d,"-reglament in day off ",gData.MyObjects[i].Name," I",i)
						e++;
				    }
				} 
			}
		  
			if r==0 {
				fmt.Println(gData.MyObjects[i].Name+" I",i," -no reglaments");
				e++;
			}
		   
			if r>0 && r<gData.MyObjects[i].RegInMonth {
				 fmt.Println(gData.MyObjects[i].Name," I",i," -few days ",r," but need ",gData.MyObjects[i].RegInMonth);
		      e++;
			}		
		}
        for d:=1;d<len(gData.Grafik);d++ {
			r:=0;
			for b:=0;b<len(gData.MyObjects);b++ { 
			  if gData.Reglament[b][d]>0 {
				  r++;
			  }	
			}
		  if r==0 && gData.Grafik[d] !=0 && gData.Grafik[d] !=5 {
			  fmt.Println("day ",d," -no reglaments");
				e++;
			}
		  
		}
		
		fmt.Println("errors found ",e);
       	fmt.Println("verification completed");
		typeLine();		
	}


func addObject(){
	   gData.TotalObjects++
	   newarr:=make([]int,32)
	   var newObject myObject
	  gData.MyObjects=append(gData.MyObjects,newObject)
	   gData.Reglament=append(gData.Reglament,newarr)
	   setObject(len(gData.MyObjects)-1);
	  if arhive {
		    writeData(fileArhive);
		    } else {
			writeData(fileData);
			} 
    }
  ////////////////
func setObject(num int){
	  
	  if num<len(gData.MyObjects) && num>-1 {
	  typeLine();
	  fmt.Println("data for the object  ",num);
	   
	  fmt.Println("enter name 0-next ",gData.MyObjects[num].Name);
	  s:=tools.St();
	  if s !="0" {
		   gData.MyObjects[num].Name=s
		   }
	  
	  fmt.Println("enter adress 0-next ",gData.MyObjects[num].Adress);
	  s=tools.St();
	  if s !="0" { 
		   gData.MyObjects[num].Adress=s
	   }
	
	  fmt.Println("enter zone  ",gData.MyObjects[num].Zone);
	  gData.MyObjects[num].Zone=tools.ReadInt();
	  fmt.Println("enter the uslUst ",gData.MyObjects[num].Ust);
	  gData.MyObjects[num].Ust=tools.ReadFloat()
      fmt.Println("enter the num of reglaments ",gData.MyObjects[num].RegInMonth);
	  gData.MyObjects[num].RegInMonth=tools.ReadInt();   
	   
    } else {
		fmt.Println("incorrect index");
		}
    }


	

	
	
	
 func saveArhive(){
	  fmt.Println("saving the database in an archive");
	  fmt.Println("1-continue 0-cancel");
	  n:=tools.ReadInt();
	  if n==1 {
	  typeLine();
	  ds:=strconv.Itoa(day);
	  if day<10 {
		  ds ="0"+strconv.Itoa(day);
	  }
	  ms:=strconv.Itoa(monthN);
	  if monthN<10 {
		  ms="0"+strconv.Itoa(monthN);
		  }
	  name:=ds+ms+strconv.Itoa(year)+".bin"
	  fmt.Println("file name ",name,"1 -yes 2-other name")
	  n:=tools.ReadInt();
	  if n==2 {
		fmt.Println("Enter filename to save");
		name=tools.St();
	  }
	  name=filesDir+"/"+name;
	  writeData(name);
	  typeLine();}
  };	
	
func typeGrafik(){
		  
		  typeLine();
		  for i:=1;i<len(gData.Grafik);i++{
			  fmt.Println(i,tools.GetDay(gData.FirstDay,i)," ",tools.GetWork(gData.Grafik[i]));
		  }
	      typeLine();
	    }
	    
func writeData(fileName string){
		file, err := os.Create(fileName)
		if err != nil{
        fmt.Println("Unable to create file:", err) 
          os.Exit(1) 
        return 
         }
    defer file.Close() 
    bytes,err := json.Marshal(gData)
    file.Write(bytes)
    fmt.Println("write Done.")
	}
	
 func search(s string, result bool) int {
	   	
	   	countSearch:=0;
	   	n:=-1;
	   	for i:=0;i<len(gData.MyObjects);i++ {
			str :=gData.MyObjects[i].Name;
			
			str= strings.ToLower(str)
			
			b:=strings.Contains(str,s)
					
			if (b) {
				countSearch++;
				n=i 
				u:=fmt.Sprintf("%.2f",gData.MyObjects[i].Ust) 
				fmt.Println(i,gData.MyObjects[i].Name,gData.MyObjects[i].Adress," Z",gData.MyObjects[i].Zone," U",u," I",i);
				regByObject(n);
			  }
		
		}
		 if n==-1 {
			 fmt.Println("not found ",s)
			}
		 if countSearch>1 && result {
			fmt.Println("enter index object");
			n=tools.ReadIntRange(0,len(gData.MyObjects)-1);}
		return n;
		}


func setAllObjects(){
	    fmt.Println("all data will be overwritten.Continue ? 1-yes");
		d:=tools.ReadInt();
		 if d==1 {
	    for i:=0;i<len(gData.MyObjects);i++ {
			 setObject(i) 
			 }
			 }
       
    }  


 func regByDay(data int){
			if data>0 && data<32  {
			exist:=false;
			n:=1;
			fmt.Println("Data:",data,tools.GetMonth(gData.Month),",",tools.GetDay(gData.FirstDay,data));
			s:=tools.GetWork(gData.Grafik[data])
			fmt.Println("grafik:",s);
			if s=="09-17" || s=="13-21" {
				typeLine()
				fmt.Println("SHORTENED WORKING HOURS !!")
				typeLine()
				}
			fmt.Println("reglaments: ");
			sum:=0.0;
			for ob:=0;ob<len(gData.MyObjects);ob++ { 
				if  gData.Reglament[ob][data]==1 { 
					u:=fmt.Sprintf("%.2f",gData.MyObjects[ob].Ust) 
					fmt.Println(n,gData.MyObjects[ob].Name," ",gData.MyObjects[ob].Adress," Z",gData.MyObjects[ob].Zone," U",u," I",ob);
					if gData.MyObjects[ob].Note !=""{
						fmt.Println("note -",gData.MyObjects[ob].Note)
						}
					
					sum+=gData.MyObjects[ob].Ust/float64 (gData.MyObjects[ob].RegInMonth);
					typeLine();
					n++;
					exist=true;
				}
			}
			if exist { 
				fmt.Printf("total usl ust %.2f \n",sum);
				} else { 
					fmt.Println("reglaments not found");
					}	 
	    } else {
			fmt.Println("incorrect day");
		     }
	}
	
	
	
func setAllReglament(){
	    
		for ob:=0;ob<len(gData.MyObjects);ob++{
			for d:=0;d<32;d++{ 
				gData.Reglament[ob][d]=0
			}
		}
		
	    for ob:=0;ob<len(gData.MyObjects);ob++{
			setReglament(ob);
		}
	  }
		
	
	
 func setReglament(ob int){
	   if ob>-1 && ob<len(gData.MyObjects) {
	   d:=0;
	   t:=0;
	   fmt.Println("object:",gData.MyObjects[ob].Name);
		
		for i:=0;i<len(gData.Reglament[ob]);i++ {
			gData.Reglament[ob][i]=0;
			}
		
		t=gData.MyObjects[ob].RegInMonth;
		for i:=1;i<t+1;i++ {
			fmt.Print("day",i,": ");
			y:=false;
			for !y {
				d=tools.ReadInt();
				if (d<32) {
					gData.Reglament[ob][d]=1;
					y=true;
					} else{
						 fmt.Println("wrong date");
						}
			}
   
			} 
		} else {
		fmt.Println("incorrect index"); 
			}
	}
	
func setAllGrafik(){
		  
		  typeLine();
		   fmt.Println("enter month 1-jan 2-feb..12-dec");
		  gData.Month=tools.ReadInt()-1;
		  fmt.Println("enter year ");
		 
		 year=tools.ReadInt();
		  fut:= time.Date(year,time.Month(gData.Month+1),1,1,1,1,1, time.Local)
		  gData.FirstDay=int(fut.Weekday())+1
		  gData.Year = year 
		  for i:=1;i<len(gData.Grafik);i++ {
			  editOne(i);
		    }
	    }
	    
	    
	    
func editOne(day int){
		  if day>0 && day<32 { 
		  typeLine();
		  fmt.Println("0-day off 1-9:18");
		  fmt.Println("2-13:22 3-9:17");
		  fmt.Println("4-13:21 5-business trip");
		   fmt.Println("grafik for the ",day,"st",gData.Grafik[day]);
		  gData.Grafik[day]=tools.ReadInt();
		} else {
			 fmt.Println("incorrect day");
		 }
	}



func typeDays(){
	   for i :=1;i<32;i++ {
		   regByDay(i);
		   typeLine();
		   typeLine();
		   }
	   
	   }
	   


func getFile(dir string) string {
	lst, err := ioutil.ReadDir(filesDir)
	
	if err != nil {
		//panic(err)
		fmt.Println("error can.t open  dir ",dir)
		return "error"
	}
	
	for i:=0;i<len(lst);i++{
	 if lst[i].IsDir()==false {
		 fmt.Println(i+1,lst[i].Name())
		 }
	 }		
		
	fmt.Println("Enter number file ");
	nf:=tools.ReadIntRange(1,len(lst));
	return lst[nf-1].Name();
	
	}



func readArhive() {
	name:=getFile(filesDir)
	if name=="error"{
		return 
		}
	fileArhive=filesDir+"/"+name
	readData(fileArhive);
	//typeAll();
	typeLine();
    }
    
    
    
 func restore(){
	if arhive {
		return
		}
	fmt.Println("restoring a database from an archive");
	fmt.Println("1-continue 0-cancel");
	nf:=tools.ReadInt();
	if nf==1{
		name:=getFile(filesDir)
		if name=="error" {
			return
			}
		fileArhive=filesDir+"/"+name;
		fmt.Println("the current database will be replaced 1-yes 0-cancel");
		nf=tools.ReadInt();
		if (nf==1){
		readData(fileArhive);
		writeData(fileData);
		}
    }
        
	}
		 
func delObject(d int){
		  if d>-1 && d<len(gData.MyObjects) {
		  fmt.Println("object ",gData.MyObjects[d].Name," will be deleted");
		  fmt.Println("continue ? 1-yes 2-cancel");
		  v:=tools.ReadInt();
		  if v==1{
			  gData.MyObjects = append(gData.MyObjects[:d],gData.MyObjects[d+1:]...)
          			
			    gData.Reglament = append(gData.Reglament[:d],gData.Reglament[d+1:]...)
			
			   gData.TotalObjects-=1;
			 if arhive {
				 writeData(fileArhive);
				 readData(fileArhive);
				 } else {
					 writeData(fileData);
					 readData(fileData);
				 }
			
		    }
	    } else{
			 fmt.Println("incorrect index");
		 }
	}



func replaceDays(d1,d2 int){
		var tmp int;
		if  d1>0 && d1<32 && d2>0 && d2<32 {
		for i:=0;i<len(gData.Reglament);i++ {
			tmp=gData.Reglament[i][d1];
			gData.Reglament[i][d1]=gData.Reglament[i][d2];
			gData.Reglament[i][d2]=tmp;
			}
		}
		if arhive {
			 writeData(fileArhive);
			 } else {
				 writeData(fileData);
			 }
		}
		
func findArhive(m int, y int) error {
	
	if m < 1 || m > 12 {
        return errors.New("month must be between 1 and 12")
    }
    if y < 2023 {
        return errors.New("year must be 2023 or later")
    }
    
    lst, err := os.ReadDir(filesDir)
	
	if err != nil {
		return errors.New("can.t open  dir "+filesDir)
	}
	
	countresult:=0
	number:=0
	
	for i:=0;i<len(lst);i++{
	 if lst[i].IsDir()==false {
		  fileArhive=filesDir+"/"+lst[i].Name()
	      readData(fileArhive);
	      if gData.Month +1  ==m && gData.Year == y {
			
			 data,err:=formatFileInfo(lst[i])
			 if err !=nil{
				 fmt.Println(err)
				 } else {
					 fmt.Println(i+1,data) // i+1 это сдвиг чтобы счет шел с 1 а не с нуля при отображении
					 }
				 
			 number=i
			 countresult++
			  }
		 }
	 }
	 	 
	 if countresult==0 {
		 return errors.New("no arhive for "+strconv.Itoa(m)+" "+strconv.Itoa(y))	
		 }
		 
	 if countresult>1{
		 	fmt.Println("enter number file");
			number=tools.ReadIntRange(1,len(lst)) -1	// для чтения сдвигаем обратно		
	
	}
	fileArhive=filesDir+"/"+lst[number].Name()
	readData(fileArhive);
	return nil
}
	 	
			
func formatFileInfo(entry fs.DirEntry) (string, error) {
	// Получаем детальную информацию о файле
	info, err := entry.Info()
	if err != nil {
		return "", err
	}

	// Получаем время изменения
	modTime := info.ModTime()

	// Форматируем строку: "ИмяФайла - ДД.ММ.ГГГГ ЧЧ:ММ"
	formattedTime := modTime.Format("02.01.2006 15:04")
	result := fmt.Sprintf("%s - %s", entry.Name(), formattedTime)

	return result, nil
}

func editNote(ind int){
	typeLine()
	
	s:=1
	for s !=0 {
		fmt.Println("note for ",gData.MyObjects[ind].Name)
	    note := gData.MyObjects[ind].Note
	    fmt.Println(note)
		fmt.Println("1 -add note")
		fmt.Println("2 -replace note")
		fmt.Println("3 -delete note")
		fmt.Println("0 -exit")
		s = tools.ReadIntRange(0,3)
		switch s {
			case 1: fmt.Println("type note for ",gData.MyObjects[ind].Name)
			        note2:=tools.St()
			        note+=note2
			        
			case 2: fmt.Println("type note for ",gData.MyObjects[ind].Name)
			        note2:=tools.St()
			        note=note2
			        
			case 3: note=""
			}
		gData.MyObjects[ind].Note = note
		writeData(fileData)	
		}
typeLine()		
}

func typeNotes(){
	typeLine()
	for _,ob :=range (gData.MyObjects) {
		if ob.Note !="" {
			fmt.Println(ob.Name,ob.Adress)
			fmt.Println("note-",ob.Note)
			typeLine()
			}
		}
	
	}


func getMaxLen() (int,int){
	maxName:=0
	maxAdress:=0
	for _,ob :=range (gData.MyObjects){
		nameLen := len([]rune(ob.Name))
		adrLen:=  len([]rune(ob.Adress))
        if nameLen > maxName {
			maxName = nameLen
			
			}
		if adrLen > maxAdress {
			maxAdress = adrLen
			
			}
		}
	return maxName,maxAdress	
	}
	
func writeToFile() {
	n:=1
	ds:=strconv.Itoa(day);
	  if day<10 {
		  ds ="0"+strconv.Itoa(day);
	  }
	ms:=strconv.Itoa(monthN);
	  if monthN<10 {
		  ms="0"+strconv.Itoa(monthN);
		  }
	filename:="objects"+ds+ms+strconv.Itoa(year)+".txt"
	
	file, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        return 
    }
    defer file.Close()
	maxName,maxAdress:=getMaxLen()
	maxName+=2
	maxAdress+=2
	for _,ob :=range (gData.MyObjects){
		fmt.Fprintf(file,"%3d %-*s %-*s %3d %5.1f\n",n,maxName,ob.Name,maxAdress,ob.Adress,ob.Zone,ob.Ust)
		n++
    }
    fmt.Println("data written to file ",filename)
	return 
	}
