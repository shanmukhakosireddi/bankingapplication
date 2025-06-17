package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)
type Bank struct{
	name string
	password string 
	enterpassword string 
	balance int 
	newbalance int 
	ammount int 
	update string 
	check string 


	

	
}




func bankfunc(details *Bank)(int){
	
	

	if details.password == details.enterpassword {
		
	
	
		fmt.Printf(" welcome back %v \n",details.name)
		check := checkbalance(details)
		if check == "yes"{
		fmt.Printf(" your  bank balance is %v \n",details.balance)
		
	}
	up := updatebalance(details)
	if up == "yes"{
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("pleas enter the amount:")
		amount,_ := reader.ReadString('\n')
		amount = strings.TrimSpace(amount)
		upamount, _ := strconv.Atoi(amount)
		
		details.balance = details.balance + upamount
		fmt.Printf("you updated balance is %v \n",details.balance)

	}
		
	}else{
		fmt.Println("incorrect password")

	}
	
	return details.balance
	


}
func verify(v *Bank) (string,string){
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	nameinput,_ := reader.ReadString('\n')
	fmt.Println("Enter your password:")
	passwordinput,_ := reader.ReadString('\n')
	passwordinput = strings.TrimSpace(passwordinput)

	return nameinput,passwordinput
}
func checkbalance(c *Bank)(string){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("DO YOU WANT TO CHECK BALANCE:")
	check,_:= reader.ReadString('\n')
	check = strings.TrimSpace(check)
	strings.ToLower(check)
	return check 
}
func updatebalance(u *Bank)(string){
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("do you want to update balance:")
	up,_:= reader.ReadString('\n')
	up = strings.TrimSpace(up)
	up = strings.ToLower(up)
	return up
	
}

func main() {
	var info Bank
	
   name,enterpassword := verify(&info)
   
   

	info = Bank{name:name,
		password:"123",
		enterpassword: enterpassword}
	bankfunc(&info)
	

}




/*

{
name
password
}

checkBalance()
UpdateBalnce()
sendBlance()
*/

func (b *Bank) checkBalance(){
	
}