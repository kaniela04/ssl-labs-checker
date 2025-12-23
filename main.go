package main 
import ("fmt" //print to console
		"os" // read args
		"net/http" // make http requests
		"io" // read response body
		"encoding/json"// parse json response
	)
/**
//structure of /info 
	type InfoResponse struct {
		EngineVersion string `json:"engineVersion"`
		
	}
**/

//struc of /analyze 
type AnalizeResponse struct {
	Host string `json:"host"`
	Status string `json:"status"`
}

func main(){ // entry point
	// validate that the user enters a domain as an argument
	if len(os.Args)<2 {
	fmt.Println("Uso: go run main.go <dominio>")
	return 
	}

	// get the domain from the command line arguments 
	domain := os.Args[1]
	
	/**
	//call the SSL Labs API to analyze the domain on /info endpoint
	url := "https://api.ssllabs.com/api/v2/info"
	**/

	//call the SSL Labs API to analyze the domain on /analyze endpoint
	url := "https://api.ssllabs.com/api/v2/analyze?host=" + domain + "&startNew=on&all=done"
	
	//the response from the API
	response, err := http.Get(url)
	if err !=nil {
		fmt.Println("Error al hacer la peticion: ", err)
		return 
	}

	//close the coneccion when the function ends
	defer response.Body.Close()

	//read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

    /**
	//parse the JSON response
	var info InfoResponse
	err = json.Unmarshal(body, &info)
	if err !=nil{
		fmt.Println("Error al pasear el JSON: ", err)
		return 
	}

	fmt.Println("Versión de la API:", info.EngineVersion)
	**/

	//parse the JSON response whit /analyze structure
	var analyze AnalizeResponse
	err = json.Unmarshal(body, &analyze)
	if err !=nil {
		fmt.Println("Error al parsear el JSON: ", err)
		return
	}

	/**
	//show the response 
	fmt.Println("Respuesta de la API de SSL Labs/info:")
	//convert text to readable format
	fmt.Println(string(body))

	//warning of variable not used
	_=domain
	//print the domain to analyze
	fmt.Println("Dominio a analizar", domain)
	**/
	//print the analyze results
	fmt.Println("Dominio analizado(host):", analyze.Host)
	fmt.Println("Estado del análisis:", analyze.Status)
}