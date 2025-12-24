package main 
import ("fmt" //print to console
		"os" // read args
		"net/http" // make http requests
		"io" // read response body
		"encoding/json"// parse json response
		"time"
		
	)
/**
//structure of /info 
	type InfoResponse struct {
		EngineVersion string `json:"engineVersion"`
		
	}
**/

//struc of endpoint /analyze
type endpoint struct {
	IPAddress string `json:"ipAddress"`
	Grade string `json:"grade"`
	StatusMessage string `json:"statusMessage"`
}
//struc of /analyze 
type AnalyzeResponse struct {
	Host string `json:"host"`
	Status string `json:"status"`
	Endpoints []endpoint `json:"endpoints"`
}

func main(){ // entry point
	// validate that the user enters a domain as an argument
	if len(os.Args)<2 {
	fmt.Println("Uso: go run main.go <dominio>")
	return 
	}

	// get the domain from the command line arguments 
	domain := os.Args[1]
	baseURL := "https://api.ssllabs.com/api/v2/analyze?host=" + domain 
	/**
	//call the SSL Labs API to analyze the domain on /info endpoint
	url := "https://api.ssllabs.com/api/v2/info"
	**/

	//call the SSL Labs API to analyze the domain on /analyze endpoint
	//url := "https://api.ssllabs.com/api/v2/analyze?host=" + domain + "&startNew=on"
	
	// start analysis 
	url := baseURL + "&startNew=on"
	//the response from the API
	response, err := http.Get(url)
	if err !=nil {
		fmt.Println("Error al iniciar el analisis: ", err)
		return 
	}

	//close the coneccion when the function ends
	body, _ := io.ReadAll(response.Body)
	response.Body.Close()
	
	var analyze AnalyzeResponse
	json.Unmarshal(body, &analyze)

	//read the response body
	// Polling
	for analyze.Status != "READY" {
		fmt.Println("Estado actual:", analyze.Status)
		time.Sleep(15 * time.Second)

		response, err = http.Get(baseURL)
		if err != nil {
			fmt.Println("Error en polling temporalmente, reintentando en unos segundos...:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		body, _ = io.ReadAll(response.Body)
		response.Body.Close()

		json.Unmarshal(body, &analyze)
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
	fmt.Println("Análisis completado...")
	fmt.Println("Dominio analizado (host):", analyze.Host)
	fmt.Println("Estado del análisis(status):", analyze.Status)
	fmt.Println("\n Resultados de segutidad TLS/SSL por endpoint:")
	for _, ep := range analyze.Endpoints {
		fmt.Println("IP del endpoint analizado:", ep.IPAddress)
		fmt.Println("Calificación del endpoint (grade):", ep.Grade)
		fmt.Println("Mensaje de estado (statusMessage):", ep.StatusMessage)

	}
}