package learn

import "fmt"

func contadorVocales(palabra string)(int, int, int, int, int){
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	// drop index, use value
	for _ , valor := range palabra { 
		variable := string(valor) // value is ASCII number, transform to str
		switch variable {
		case "a" :
			conta++
		case "e" :
			conte++
		case "i" :
			conti++
		case "o" :
			conto++
		case "u" :
			contu++
		}
	}
	return conta, conte , conti, conto, contu
}

func Switch_ex() {
	palabra := "Some phrase in which we will count its vowels"	
	a,e,i,o,u := contadorVocales(palabra)
	fmt.Printf("la frase '%s' tiene: \n",palabra)
	fmt.Printf("%d vocales a \n",a)
	fmt.Printf("%d vocales e \n",e)
	fmt.Printf("%d vocales i \n",i)
	fmt.Printf("%d vocales o \n",o)
	fmt.Printf("%d vocales u \n",u)	
}