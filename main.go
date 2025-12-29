package main

import (
	"encoding/json" // parse json response
	"fmt"           //print to console
	"io"            // read response body
	"net/http"      // make http requests
	"os"            // read args
	"time"
)

/**
 * Endpoint represents a single server endpoint returned by the SSL Labs
 * /analyze API. A domain can have multiple endpoints (IP addresses),
 * each evaluated independently.
 */
type endpoint struct {
	IPAddress     string `json:"ipAddress"` // IP address of the endpoint
	Grade         string `json:"grade"` // Overall security grade assigned by SSL Labs
	StatusMessage string `json:"statusMessage"` // Status message for the endpoint analysis
}

/**
 * AnalyzeResponse represents the main response structure returned by
 * the SSL Labs /analyze endpoint.
 */
type AnalyzeResponse struct {
	Host          string     `json:"host"` // Domain being analyzed
	Status        string     `json:"status"` // Current status of the analysis
	Endpoints     []endpoint `json:"endpoints"` // List of endpoints analyzed
	StatusMessage string     `json:"statusMessage"`// Message providing additional status information
}

/**
 * handleHTTPError processes HTTP status codes returned by the SSL Labs API.
 *
 * @param code HTTP status code returned by the API
 * @return bool Returns true if execution can continue, false otherwise
 */
func handleHTTPError(code int) bool {
	switch code {
	case 200:
		return true
	case 400:
		fmt.Println("400: Solicitud incorrecta (dominio incorrecto) )")
	case 429:
		fmt.Println("429: Demasiadas solicitudes (limite de tasa excedido) espere 1 minuto...")
		time.Sleep(60 * time.Second)
	case 500:
		fmt.Println("500: Error interno del servidor de SSL Labs, intente nuevamente mas tarde")
	case 503:
		fmt.Println("503: Servicio no disponible, intente nuevamente mas tarde")
	case 529:
		fmt.Println("529: servidor sobrecargado - el servicio esta siendo utilizado en exceso, espere un minuto...")
		time.Sleep(60 * time.Second)
	default:
		fmt.Println("Error desconocido:", code)
	}
	return false
}

/**
 * main is the entry point of the application.
 * It validates input, starts the SSL analysis, polls the API until
 * completion, and prints the final TLS/SSL security results.
 */
func main() { // entry point
	// validate that the user enters a domain as an argument
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <dominio>")
		return
	}

	// Retrieve domain from command-line arguments
	domain := os.Args[1]

	// base URL for SSL Labs API
	baseURL := "https://api.ssllabs.com/api/v2/analyze?host=" + domain + "&all=done"

	// start a new analysis
	url := baseURL + "&startNew=on"
	//the response from the API
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al iniciar el analisis: ", err)
		return
	}

	// Handle HTTP-level errors
	if !handleHTTPError(response.StatusCode) {
		return
	}

	// Read initial response body
	body, _ := io.ReadAll(response.Body)
	response.Body.Close()

	var analyze AnalyzeResponse
	json.Unmarshal(body, &analyze)

	/**
	 * Polling loop:
	 * Repeatedly queries the API until the analysis status is READY or ERROR.
	 * A delay is applied between requests to respect API rate limits.
	 */
	for {
		fmt.Println("Estado actual:", analyze.Status)

		// Exit loop when analysis is completed
		if analyze.Status == "READY" {
			break
		}

		// Stop execution if the analysis fails
		if analyze.Status == "ERROR" {
			fmt.Println("Error en el análisis:", analyze.StatusMessage)
			return
		}
		// Wait before the next polling attempt
		time.Sleep(15 * time.Second)

		// Poll the API for updated analysis status
		response, err = http.Get(baseURL)
		if err != nil {
			fmt.Println("Error en polling temporalmente, reintentando en unos segundos...:", err)
			time.Sleep(30 * time.Second)
			continue
		}
		// Handle HTTP-level errors
		if !handleHTTPError(response.StatusCode) {
			continue
		}
		// Read and parse the response body
		body, _ = io.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(body, &analyze)
	}

	//print the analyze results
	fmt.Println("Análisis completado...")
	fmt.Println("Dominio analizado (host):", analyze.Host)
	fmt.Println("Estado del análisis(status):", analyze.Status)
	fmt.Println("\n Resultados de seguridad TLS por endpoint:")
	// Iterate over each endpoint and print its results
	for _, ep := range analyze.Endpoints {
		fmt.Println("IP del endpoint analizado:", ep.IPAddress)
		fmt.Println("Calificación del endpoint (grade):", ep.Grade)
		fmt.Println("Estado del análisis (status):", ep.StatusMessage)

	}
}
