package main 
import ("fmt"
		"os")

func main(){
	if len(os.Args)<2 {
	fmt.Println("Uso: go run main.go <dominio>")
	return 
	}
	domain := os.Args[1]
	fmt.Println("Dominio a analizar", domain)
}