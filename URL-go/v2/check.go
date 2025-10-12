/*
Check v2
usage: check -url <url>
example with flags: check -threads 25 -list domains.txt -ignore 404,403 > log.txt
usage:
	- flag -url to send the request to that url
	- flag -ignore to ignore customs status code
	- flag -threads to control max goroutines
	- flag -ua to put custom user-agent
	- flag -list to put a list of urls
	- flag -see to see the body (if you use -list)
	- flag -timeout to custom the timeout
*/
package main
// dependecies
import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"flag"
	"os"
	"encoding/json"
	"bytes"
	"bufio"
	"sync"
	"time"
	"strings"
	"strconv"
	
	
)	// maps
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
// Main
func main() {
	//colors
	var (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
	)
	
	var ListaUrl []string // a slice of urls
	// flags
	url := flag.String("url", "", "URL to check (GET)")
	ua := flag.String("ua", "", "An optional flag to put custom UserAgent")
	list := flag.String("list", "", "An optional flag to check all the domains in a list")
	see := flag.Bool("see", false, "It only can be used with -list. It is for see all the html (not recomended)")
	threads := flag.Int("threads", 5, "To limite goroutines (default 5)")
	timeout := flag.Int("timeout", 5, "An optional flag to put custom timeout (default: 5s)")
	ignore := flag.String("ignore", "", "Comma-separated list of HTTP status codes to ignore (e.g. 404,403,500")
	flag.Parse()
	limiter := make(chan struct{}, *threads) // limiter
	var wg sync.WaitGroup //Wait group
	htmlSee := *see
	wordList := *list
	UserAgent := *ua
	ignoreCodes := make(map[int]bool) // the ignore codes (if you used -ignore)
	// all the flags scripts
	if *ignore != "" {
		for _, number := range strings.Split(*ignore, ",") {
			code, err := strconv.Atoi(strings.TrimSpace(number))
			if err == nil {
				ignoreCodes[code] = true
			}
		}
	}
	if *url == "" && wordList == "" {
		fmt.Println(Red+"Usage: check -url https://example.com"+Reset)
		os.Exit(1)

	}
	valid:= false
	if *url != "" {
		ListaUrl = append(ListaUrl, *url)
		htmlSee = true
		valid = true
	}
	if UserAgent == "" {
		UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0"
	}
	if wordList != "" {
		file, err := os.Open(wordList)
		if err != nil {
			fmt.Println(Red+"[X]Error opening the file. It might not exist." + Reset)
			os.Exit(1)
		
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ListaUrl = append(ListaUrl, scanner.Text())
		}
	}
	if htmlSee && wordList == "" && valid == false {
		fmt.Println(Red+"You can't use see without list."+Reset)
		os.Exit(1)
	}
	
	start := time.Now() // start timer
	request:=resty.New() // client
	request.SetTimeout(time.Duration(*timeout) * time.Second) // Timeout
	//all the source code
	if len(ListaUrl) > 0 {
		// request
		for _, link := range ListaUrl { 
			wg.Add(1)
			limiter <- struct{}{}
			go func(link string) {
				defer wg.Done()
				defer func ()  { <-limiter}()	
				response, err :=  request.R().SetHeader("User-Agent", UserAgent).Get(link)
				if err != nil  {
					fmt.Printf("%s[!] Error fetching %s: %v%s\n", Red, link, err, Reset)
					return
				}
		// ignore status code
		if ignoreCodes[response.StatusCode()] {
			return
		}
		// -see
		if htmlSee {
			body := response.Body()
			var pretty bytes.Buffer
			err = json.Indent(&pretty, body, "", "  ")
			if err != nil {
				fmt.Printf("%s[?] Non-JSON body from %s — showing raw output:%s\n", Red, link, Reset)
				fmt.Println(string(body))
			} else {
				fmt.Println(pretty.String())
			}
		}
		// bad status code like: 404, 429
		for clave, valor := range  codigosError {
			if clave == response.StatusCode() {
				fmt.Printf("\n%s %d %s %s\n", Red, clave, valor, Reset)
				fmt.Println("final url: ", response.Request.URL)
				fmt.Println("UserAgent: ", UserAgent)
			}
		}
		// good status codes
		for clave, valor := range  codigosActivos {
			if clave == response.StatusCode() {
				fmt.Printf("\n%s %d %s  %s\n", Green, clave, valor, Reset)
				fmt.Println("final url: ", response.Request.URL)
				fmt.Println("UserAgent: ", UserAgent)

			}
		}
			}(link) // to pass the link argument
		}
			
	}
	wg.Wait() // to wait untilall the goroutines finished
	finish := time.Since(start) 
	fmt.Printf("\nDone! Completed %d urls with %d threads in %s! \n", len(ListaUrl), *threads, finish)
}
