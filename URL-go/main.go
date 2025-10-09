package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"flag"
	"os"
	"encoding/json"
	"bytes"
	
)

func main() {
	var (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
	)
	url := flag.String("url", "", "URL to check (GET)")
	flag.Parse()
	if *url == "" {
		fmt.Println(Red+"Usage: check -url https://example.com"+Reset)
		os.Exit(1)

	}
	peticion:=resty.New()
	response, err :=  peticion.R().Get(*url)
	if err != nil  {
		panic(err)
	}
		
	var codigosActivos = map[int]string{
		200: "OK – Solicitud exitosa.",
		201: "Created – Recurso creado correctamente.",
		202: "Accepted – Solicitud aceptada, aún no procesada.",
		203: "Non-Authoritative Information – Información modificada por intermediario.",
		204: "No Content – Solicitud exitosa sin contenido.",
		205: "Reset Content – Solicitud exitosa, reiniciar la vista.",
		206: "Partial Content – Se envió solo parte del recurso.",
	}


	var codigosError = map[int]string{
		400: "Bad Request – Solicitud malformada.",
		401: "Unauthorized – No autenticado.",
		402: "Payment Required – Se requiere pago.",
		403: "Forbidden – Acceso denegado.",
		404: "Not Found – Recurso no encontrado.",
		405: "Method Not Allowed – Método no permitido.",
		406: "Not Acceptable – No cumple requisitos de respuesta.",
		407: "Proxy Authentication Required – Autenticación del proxy requerida.",
		408: "Request Timeout – Tiempo de espera agotado.",
		409: "Conflict – Conflicto con el estado actual.",
		410: "Gone – Recurso eliminado permanentemente.",
		411: "Length Required – Falta 'Content-Length'.",
		412: "Precondition Failed – Fallaron las precondiciones.",
		413: "Payload Too Large – Cuerpo demasiado grande.",
		414: "URI Too Long – URL demasiado larga.",
		415: "Unsupported Media Type – Tipo no soportado.",
		416: "Range Not Satisfiable – Rango no válido.",
		417: "Expectation Failed – Falló la expectativa.",
		418: "I'm a teapot – Broma oficial del protocolo.",
		421: "Misdirected Request – Servidor incorrecto.",
		422: "Unprocessable Entity – Error de validación.",
		423: "Locked – Recurso bloqueado.",
		424: "Failed Dependency – Dependencia fallida.",
		425: "Too Early – Solicitud prematura.",
		426: "Upgrade Required – Requiere actualizar protocolo.",
		428: "Precondition Required – Se requieren precondiciones.",
		429: "Too Many Requests – Demasiadas solicitudes.",
		431: "Request Header Fields Too Large – Encabezados muy grandes.",
		451: "Unavailable For Legal Reasons – Bloqueado legalmente.",
		500: "Internal Server Error – Error interno.",
		501: "Not Implemented – Función no soportada.",
		502: "Bad Gateway – Error entre servidores.",
		503: "Service Unavailable – Servicio no disponible.",
		504: "Gateway Timeout – Tiempo de espera agotado.",
		505: "HTTP Version Not Supported – Versión no soportada.",
		506: "Variant Also Negotiates – Error de negociación.",
		507: "Insufficient Storage – Almacenamiento insuficiente.",
		508: "Loop Detected – Bucle infinito.",
		510: "Not Extended – Requiere extensiones.",
		511: "Network Authentication Required – Autenticación de red requerida.",
	}

	body := response.Body()
	var pretty bytes.Buffer
	err = json.Indent(&pretty, body, "", "  ")
	if err != nil {
		fmt.Println("Invalid Json, the returned value will be printed (Probably html+css): ")
		fmt.Println(string(body))
	} else {
		fmt.Println(pretty.String())
	}
	for clave, valor := range  codigosError {
		if clave == response.StatusCode() {
			fmt.Printf("\n%s %d %s %s\n", Red, clave, valor, Reset)
			fmt.Println("final url: ", response.Request.URL)
		}
	}
	for clave, valor := range  codigosActivos {
		if clave == response.StatusCode() {
			fmt.Printf("\n%s %d %s  %s\n", Green, clave, valor, Reset)
			fmt.Println("final url: ", response.Request.URL)

		}
	}
}
