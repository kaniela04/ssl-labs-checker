# SSL Labs Go Checker
CLI tool written in Go that analyzes the TLS/SSL security of a domain using the SSL Labs API.

## How to run

```bash
go run main.go google.com


---

###  4. Ejemplo de salida

```md
## Example output

Estado actual: IN_PROGRESS  
Análisis completado...

Dominio analizado: google.com  
Estado: READY  

Endpoint IP: 172.217.215.139  
Grade: B  
Status: Ready
## Technical decisions

- Polling every 15 seconds to avoid overloading the SSL Labs API
- Retry mechanism for temporary network or rate-limit errors
- Minimal JSON parsing focused on required fields
- Compatible with long-running SSL Labs analyses

## Known limitations

- Long analyses may take several minutes for large domains
- SSL Labs may close connections during long polling
- Only basic endpoint data is displayed (IP, grade, status)

