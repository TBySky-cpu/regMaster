 package main
 import(
 "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"
    "strconv"
    "regMaster/tools"
    "io/ioutil"
    )
 var  allObject, //
	  day, 
      firstDay,//	  1st day of week 0-sun 1-mon...6-sat
	  year,
	  month,
	  modeMenu,
	  modeComand,
	  mode int
var  arhive bool
var	  fileData, //= "data.dpm";
	  filesDir,//,="arhive";
	  fileArhive string
var	  myObjects []myObject;
var	  grafik [32]int
var	  reglament [][]int
var   calendar time.Time
	  
  type myObject struct{
	name string;
	adress string;
	zone int;
	ust float64;
	regInMonth int;
	}
	
	
 func main(){
	  modeMenu=1;
	  modeComand=2;
	  mode=modeComand;
	 calendar=time.Now()
	fileData="data.dpm"
	  filesDir="arhive"
	  day=calendar.Day()
	 var  monthN int
	  monthN= int(calendar.Month())
	  year=calendar.Year()
	  initr()  
	  typeLine();         
	  fmt.Println("*      RegMaster v.5.4      *");
	  fmt.Println("*Takhir Bairashevski nov2024*");
	  typeLine();
	  fmt.Println("Today ",day,monthN,year);
	  
	  if month+1 != monthN {
		   fmt.Println ("database is out of date");
		   }
	  regByDay(day);
	  typeLine();
	  if mode==modeMenu {
		fmt.Println("1-menu  0-exit"); 	   
	    s:=tools.ReadInt();
        if s==1 {
			//mainMenu()}
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
		      allObject =tools.ReadInt();
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
	myObjects=make([]myObject,allObject+5)
	reglament=make([][]int,allObject+5)
	for i:=0;i<allObject+5;i++{
		reglament[i]=make([]int,32)
	 }
	
}




 func readData(nameFile string){
	  ob:=0;
	file, err := os.Open(nameFile)
    if err != nil {
        log.Fatal(err)
    }
   
	 
	 scanner := bufio.NewScanner(file)
	 for scanner.Scan() {
        str:=(scanner.Text())
        strs:=strings.Split(str," ")
        if len(strs)>1 {
			switch strs[0]{
				case "month":
					month, _ = strconv.Atoi(strs[1])
			    case "firstDay":
					firstDay, _ = strconv.Atoi(strs[1])
				case "total_objects": {
					allObject, _ = strconv.Atoi(strs[1])
					makeArray()
					}
				case "index":
					ob,_= strconv.Atoi(strs[1])
				case "uslUst": {
					u, _ := strconv.Atoi(strs[1])
					myObjects[ob].ust=UstToFloat(u)
					}	
				case "zone":
					myObjects[ob].zone, _ = strconv.Atoi(strs[1])
				case "name":
					myObjects[ob].name = strs[1]
				case "adress":
					myObjects[ob].adress= strs[1]
				case "regInMonth":
					myObjects[ob].regInMonth, _ = strconv.Atoi(strs[1])
			}
		}
		if strs[0]=="Grafik" {
			if scanner.Scan(){
			str=(scanner.Text())
			strg:=strings.Split(str," ")
			for i:=0;i<len(grafik);i++{
				grafik[i],_=strconv.Atoi(strg[i])
				}
			}
		}
		if strs[0]=="Reglament"{
			for obj:=0;obj<len(reglament);obj++ {
				 if scanner.Scan(){
					str=(scanner.Text())
			        strr:=strings.Split(str," ")
				for d:=0;d<len(reglament[obj]);d++{
					reglament[obj][d],_=strconv.Atoi(strr[d])
					}
				  }
				}
			
			} 
			 
    }
    file.Close()
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
		  for i:=0;i<allObject;i++ {
		  u:=fmt.Sprintf("%.2f",myObjects[i].ust)  
		  fmt.Println(i,myObjects[i].name," ",myObjects[i].adress," Z",myObjects[i].zone," U",u);
		  regByObject(i);
		  sum+=myObjects[i].ust;
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
		  for d:=0;d<len(reglament[index]);d++ { 
			if (reglament[index][d]==1){ 
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
				fmt.Println("**working in arhive**")
			    } 
			fmt.Print("> ");
			str1=tools.St();
			str=strings.Trim(str1, " ")
			if str=="" || len(str) ==0 {
				continue
				}
			if checkComand(str) {
				
			parts:=strings.Split(str," ")
            if len(parts)==1 {
				switch parts[0]{
					case "exit": {
						err=false;
					    work=false
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
					case "restore": {
					restore();
					err=false;
					}
				}
				} 
			  
             if len(parts)==2 {
				switch parts[0] {
					case "day": {
						regByDay(tools.ToInt(parts[1]));
						err=false; 
						}
					case "obt": {
						ind=search(parts[1]);
						err=false;
						}
				
			
					case "del" :{
						 ind =search(parts[1]);
						 if  ind>-1  {
							 delObject(ind);
							 err=false;}
							  }
					}
				} 
			   if len(parts)==3 {
				switch parts[0] {
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
									 ind=search(parts[2]);
							          if ind>-1 {
										  setObject(ind) 
										  err=false;
							        }	
							     }
						        }
						        
							case "reg": {
								if  parts[2]=="all" {
									setAllReglament();
									err=false;
								} else { 
									ind=search(parts[2]);
									 if (ind>-1) {
											setReglament(ind);
											err=false;
											}
								}
							}
						    case "gr" : {
								if parts[2]=="all" {
									setAllGrafik();
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
		
		 mainComand:= []string{"q","add","check","menu","help","exit","rep","day","obt","obs",
			                   "grk","greg","gust","del","arhr","arhs","ed","restore"};
		parts:=strings.Split(s," ")
		
		s1:=parts[0];
		for i:=0;i<len(mainComand);i++{ 
		    if s1== mainComand[i] {
				b=true;
				}
		     }
		
		 if !b { 
			 fmt.Println("command not found  "+s1)
			 }
		return b;
		}	    	  


func typeHelp(){
		fmt.Println("add -add object");
		fmt.Println("arhr -read arhive file");
		fmt.Println("arhs -save in arhive");
		fmt.Println("check   -data base check");
		fmt.Println("del <name object> -delete object");
		fmt.Println("day <number day> -type grafik and objects for the current day");
		fmt.Println("help -type this help");
		fmt.Println("menu   -mode with menu");
		fmt.Println("obt <name object>  - type data for the object");
		fmt.Println("obs  -type all objects");
		fmt.Println("grk-type grafik");
		fmt.Println("greg-type regl and grafik");
		fmt.Println("gust-type grafik and usl ust");
		fmt.Println("ed ob all -edit all objects");
		fmt.Println("ed ob <name object> -edit one object");
		fmt.Println("ed reg all  -edit all reglaments");
		fmt.Println("ed reg <name object> -edit one reglament");
		fmt.Println("ed gr all -edit all grafik");
		fmt.Println("ed gr <day>  -edit one day");
		fmt.Println("rep <day1> <day2>  -replacing reg from day1 to day2");
		fmt.Println("restore -restoring a database from an archive");
		}	
	
func typeUst(){
	fmt.Println("type Ust")
	sum:=0.0;
	 min:=100.0;
	 max:=0.0;
	 minDay:=0
	 maxDay:=0
		for i :=1;i<len(grafik);i++{
			if grafik[i] !=0 && grafik[i] !=5 {
			for ob:=0;ob<allObject;ob++ { 
				if reglament[ob][i]==1 {
					sum+=myObjects[ob].ust/float64(myObjects[ob].regInMonth)
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
			fmt.Println(i,tools.GetDay(firstDay,i)," ",tools.GetWork(grafik[i])," U",u);
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
		for i:=0;i<allObject;i++ {
		  r:=0;
	      for d:=1;d<len(reglament[i]); d++ {
				if reglament[i][d]>0 {
					r++;
					if  grafik[d]==0  {
						fmt.Println("day ",d,"-reglament in day off ",myObjects[i].name," I",i)
						e++;
				    }
				} 
			}
		  
			if r==0 {
				fmt.Println(myObjects[i].name+" I",i," -no reglaments");
				e++;
			}
		   
			if r>0 && r<myObjects[i].regInMonth {
				 fmt.Println(myObjects[i].name," I",i," -few days ",r," but need ",myObjects[i].regInMonth);
		      e++;
			}		
		}
        for d:=1;d<len(grafik);d++ {
			r:=0;
			for b:=0;b<allObject;b++ { 
			  if reglament[b][d]>0 {
				  r++;
			  }	
			}
		  if r==0 && grafik[d] !=0  {
			  fmt.Println("day ",d," -no reglaments");
				e++;
			}
		}
		
		fmt.Println("errors found ",e);
       	fmt.Println("verification completed");
		typeLine();		
	}


func addObject(){
	   allObject++;
	   newarr:=make([]int,32)
	   reglament=append(reglament,newarr)
	   setObject(allObject-1);
	  if arhive {
		    writeData(fileArhive);
		    } else {
			writeData(fileData);
			} 
    }
  ////////////////
func setObject(num int){
	  
	  if num<allObject && num>-1 {
	  typeLine();
	  fmt.Println("data for the object  ",num);
	   
	  fmt.Println("enter name 0-next ",myObjects[num].name);
	  s:=tools.St();
	  if s !="0" {
		   myObjects[num].name=s
		   }
	  
	  fmt.Println("enter adress 0-next ",myObjects[num].adress);
	  s=tools.St();
	  if s !="0" { 
		   myObjects[num].adress=s
	   }
	
	  fmt.Println("enter zone  ",myObjects[num].zone);
	  myObjects[num].zone=tools.ReadInt();
	  fmt.Println("enter the uslUst ",myObjects[num].ust);
	  myObjects[num].ust=tools.ReadFloat()
      fmt.Println("enter the num of reglaments ",myObjects[num].regInMonth);
	  myObjects[num].regInMonth=tools.ReadInt();   
	   
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
	  fmt.Println("Enter filename to save");
	  fn:=tools.St();
	  fn=filesDir+"/"+fn;
	  writeData(fn);
	  typeLine();}
  };	
	
func typeGrafik(){
		  
		  typeLine();
		  for i:=1;i<len(grafik);i++{
			  fmt.Println(i,tools.GetDay(firstDay,i)," ",tools.GetWork(grafik[i]));
		  }
	      typeLine();
	    }
	    
func writeData(fileName string){
		file, err := os.Create(fileName)
		if err != nil{
        fmt.Println("Unable to create file:", err) 
        return 
         }
		
			file.WriteString("month "+strconv.Itoa(month)+" \n");
			file.WriteString("firstDay "+ strconv.Itoa(firstDay)+" \n");
			file.WriteString("total_objects "+ strconv.Itoa(allObject)+" \n");
			for i:=0;i<allObject;i++ {
				file.WriteString("index "+strconv.Itoa(i)+"\n");
				file.WriteString("uslUst "+strconv.Itoa(tools.UstToInt(myObjects[i].ust))+"\n");
				file.WriteString("zone "+strconv.Itoa(myObjects[i].zone)+"\n");
				file.WriteString("name "+myObjects[i].name+"\n");
				file.WriteString("adress "+myObjects[i].adress+"\n");
				file.WriteString("regInMonth "+strconv.Itoa(myObjects[i].regInMonth)+"\n");		
			}
		
			file.WriteString("Grafik \n"); 
			for i:=0;i<len(grafik);i++ {
				  file.WriteString(strconv.Itoa(grafik[i])+" "); 
				    
			}
			file.WriteString("\n");
			file.WriteString("Reglament \n");    
            for i:=0;i<len(reglament);i++ {
				for b:=0;b<len(reglament[i]);b++ { 
					v:=reglament[i][b];
					file.WriteString(strconv.Itoa(v)+" ");
				}
					file.WriteString("\n");
			}
	        file.Close() 
			
			fmt.Println("The file has been written");
		
		
    }
	
 func search(s string) int {
	   	
	   	countSearch:=0;
	   	n:=-1;
	   	for i:=0;i<len(myObjects);i++ {
			str :=myObjects[i].name;
			
			str= strings.ToLower(str)
			
			b:=strings.Contains(str,s)
					
			if (b) {
				countSearch++;
				n=i 
				u:=fmt.Sprintf("%.2f",myObjects[i].ust) 
				fmt.Println(i,myObjects[i].name,myObjects[i].adress," Z",myObjects[i].zone," U",u," I",i);
				regByObject(n);
			  }
		
		}
		 if n==-1 {
			 fmt.Println("not found ",s)
			}
		 if (countSearch>1) {
			fmt.Println("enter index object");
			n=tools.ReadInt();}
		return n;
		}


func setAllObjects(){
	    fmt.Println("all data will be overwritten.Continue ? 1-yes");
		d:=tools.ReadInt();
		 if d==1 {
	    for i:=0;i<allObject;i++ {
			 setObject(i) 
			 }
			 }
       
    }  


 func regByDay(data int){
			if data>0 && data<32  {
			exist:=false;
			n:=1;
			fmt.Println("Data:",data,tools.GetMonth(month),",",tools.GetDay(firstDay,data));
			fmt.Println("grafik:",tools.GetWork(grafik[data]));
			fmt.Println("reglaments: ");
			sum:=0.0;
			for ob:=0;ob<allObject;ob++ { 
				if  reglament[ob][data]==1 { 
					u:=fmt.Sprintf("%.2f",myObjects[ob].ust) 
					fmt.Println(n,myObjects[ob].name," ",myObjects[ob].adress," Z",myObjects[ob].zone," U",u," I",ob);
					sum+=myObjects[ob].ust/float64 (myObjects[ob].regInMonth);
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
	    fmt.Println("all data will be overwritten.Continue ? 1-yes");
		d:=tools.ReadInt();
		if (d==1){
		for ob:=0;ob<allObject+5;ob++{
			for d:=0;d<32;d++{ 
				reglament[ob][d]=0
			}
		}
		
	    for ob:=0;ob<allObject;ob++{
			setReglament(ob);
		}
	  }
		
	}
	
 func setReglament(ob int){
	   if ob>-1 && ob<allObject {
	   d:=0;
	   t:=0;
	   fmt.Println("object:",myObjects[ob].name);
		
		//for i:=0;i<len(reglament[ob]);i++ {
			//reglament[ob][i]=0;
			//}
		
		t=myObjects[ob].regInMonth;
		for i:=1;i<t+1;i++ {
			fmt.Print("day",i,": ");
			y:=false;
			for !y {
				d=tools.ReadInt();
				if (d<32) {
					reglament[ob][d]=1;
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
		  month=tools.ReadInt()-1;
		  fmt.Println("enter year ");
		 
		 year=tools.ReadInt();
		  fut:= time.Date(year,time.Month(month+1),1,1,1,1,1, time.Local)
		  firstDay=int(fut.Weekday())+1
		
		  for i:=1;i<len(grafik);i++ {
			  editOne(i);
		    }
	    }
	    
	    
	    
func editOne(day int){
		  if day>0 && day<32 { 
		  typeLine();
		  fmt.Println("0-day off 1-9:18");
		  fmt.Println("2-13:22 3-9:17");
		  fmt.Println("4-13:21 5-business trip");
		   fmt.Println("grafik for the ",day,"st",grafik[day]);
		  grafik[day]=tools.ReadInt();
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
	nf:=tools.ReadInt();
	return lst[nf-1].Name();
	
	}



func readArhive(){
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
		  if d>-1 && d<allObject {
		  fmt.Println("object ",myObjects[d].name," will be deleted");
		  fmt.Println("continue ? 1-yes 2-cancel");
		  v:=tools.ReadInt();
		  if v==1{
			  for i:=d;i<allObject+1;i++ {
				  myObjects[i].name=myObjects[i+1].name;
				  myObjects[i].adress=myObjects[i+1].adress;
				  myObjects[i].zone=myObjects[i+1].zone; 
				  myObjects[i].ust=myObjects[i+1].ust;
                  myObjects[i].regInMonth=myObjects[i+1].regInMonth;
          			
			    }
			  
			  for b:=d;b<allObject;b++ {
				  for day:=0;day<len(reglament[b]);day++{
					  reglament[b][day]=reglament[b+1][day];
				    }
				}
			  allObject -= 1;
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
		for i:=0;i<len(reglament);i++ {
			tmp=reglament[i][d1];
			reglament[i][d1]=reglament[i][d2];
			reglament[i][d2]=tmp;
			}
		}
		if arhive {
			 writeData(fileArhive);
			 } else {
				 writeData(fileData);
			 }
		}
