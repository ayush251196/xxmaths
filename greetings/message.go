package greetings

import "rsc.io/quote"
//GetHelloMessage returns a greeting message
func GetHelloMessage()(string){
	return quote.Hello();
}

