# SSL Labs Go Checker
CLI tool written in Go that analyzes the TLS security of a domain using the SSL Labs API.

## How to run

```bash
go run main.go google.com

---
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

#Error Response Status Codes
The following status codes are used:

400 - invocation error (e.g., invalid parameters)
429 - client request rate too high or too many new assessments too fast
500 - internal error
503 - the service is not available (e.g., down for maintenance)
529 - the service is overloaded

# SSL LABS RATING REMINDER
A+,A,A- -> Excellent security
B -> Good
C/D -> Fair / Poor
F -> Very poor
T -> Certificate not trusted
M -> Certificate name does not match