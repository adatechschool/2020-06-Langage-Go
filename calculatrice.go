package main 

import "fmt"

func main(){
	var operator string  
	var number1, number2 float64 
	fmt.Print("Please enter First number: ")  
	fmt.Scanln(&number1)  
	fmt.Print("Please enter Second number: ")  
	fmt.Scanln(&number2)  
	fmt.Print("Please enter Operator (+,-,/,%,*):")  
	fmt.Scanln(&operator)  
	output := 0.0  
	switch operator {  
	case "+":  
	 output = number1 + number2  
	 break  
	case "-":  
	 output = number1 - number2  
	case "*":  
	 output = number1 * number2  
	case "/":  
	 output = number1 / number2  
	default:  
	 fmt.Println("Invalid Operation")  
	}  
	fmt.Printf("%.2f %s %.2f = %.2f\n", number1, operator, number2, output)  
   }  


