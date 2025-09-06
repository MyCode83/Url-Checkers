import httpx, sys, json
from colorama import Fore, init
from fake_useragent import UserAgent
ua = UserAgent()
init()
_headers=None
if len(sys.argv) > 1:
    try:
        datos=sys.argv[1]
        
        with open(datos, "r") as f:
            _headers=json.load(f)

    except FileNotFoundError:
        print(Fore.RED+"Error comprueba la ruta del archivo json, cabeceras por defecto")
        _headers = None
    except json.JSONDecodeError:
        print(Fore.RED+"El json tiene un error, usando cabeceras por defecto.")
        _headers = None


if not _headers:
    _headers = {
            "User-Agent": ua.random,
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
            "Accept-Language": "es-ES,es;q=0.9,en;q=0.8",
            "Accept-Encoding": "gzip, deflate, br",
            "Connection": "keep-alive",
            "Upgrade-Insecure-Requests": "1",
            "Sec-Fetch-Dest": "document",
            "Sec-Fetch-Mode": "navigate",
            "Sec-Fetch-Site": "none",
            "Sec-Fetch-User": "?1",
            "Cache-Control": "max-age=0"
        }

while True:
    URL = input("Que url desea ver si esta activo: ").lower()
    try:
        http = httpx.get(url=f"https://{URL}", follow_redirects=True, headers=_headers)
    except httpx.ConnectError as hc:
        print(Fore.RED+f"Error en la conexion: {hc}")
    activos = {
    200: "200 OK – Solicitud exitosa.",
    201: "201 Created – Recurso creado correctamente.",
    202: "202 Accepted – Solicitud aceptada, aún no procesada.",
    203: "203 Non-Authoritative Information – Información modificada por intermediario.",
    204: "204 No Content – Solicitud exitosa sin contenido.",
    205: "205 Reset Content – Solicitud exitosa, reiniciar la vista.",
    206: "206 Partial Content – Se envió solo parte del recurso (ej. descarga parcial)."
}


    error = {
    300: "300 Multiple Choices – Múltiples opciones posibles.",
    301: "301 Moved Permanently – Redirección permanente.",
    302: "302 Found – Redirección temporal (antes: Moved Temporarily).",
    303: "303 See Other – Consultar otro recurso mediante GET.",
    304: "304 Not Modified – No ha cambiado desde la última solicitud.",
    305: "305 Use Proxy – Debe accederse mediante proxy (obsoleto).",
    306: "306 (Unused) – Código reservado, ya no se usa.",
    307: "307 Temporary Redirect – Redirección temporal (mismo método).",
    308: "308 Permanent Redirect – Redirección permanente (mismo método).",
    400: "400 Bad Request – Solicitud malformada.",
    401: "401 Unauthorized – No autenticado.",
    402: "402 Payment Required – Se requiere pago (reservado).",
    403: "403 Forbidden – Acceso denegado.",
    404: "404 Not Found – Recurso no encontrado.",
    405: "405 Method Not Allowed – Método no permitido.",
    406: "406 Not Acceptable – No cumple con requisitos de respuesta.",
    407: "407 Proxy Authentication Required – Autenticación del proxy requerida.",
    408: "408 Request Timeout – Tiempo de espera agotado.",
    409: "409 Conflict – Conflicto con el estado actual del recurso.",
    410: "410 Gone – Recurso eliminado permanentemente.",
    411: "411 Length Required – Falta el encabezado 'Content-Length'.",
    412: "412 Precondition Failed – Fallaron las precondiciones.",
    413: "413 Payload Too Large – Cuerpo de solicitud demasiado grande.",
    414: "414 URI Too Long – URL demasiado larga.",
    415: "415 Unsupported Media Type – Tipo de archivo no soportado.",
    416: "416 Range Not Satisfiable – Rango solicitado no válido.",
    417: "417 Expectation Failed – Falló la expectativa definida.",
    418: "418 I'm a teapot – (Broma del protocolo HTCY) 🫖",
    421: "421 Misdirected Request – Solicitud dirigida al servidor incorrecto.",
    422: "422 Unprocessable Entity – Error de validación en contenido.",
    423: "423 Locked – Recurso bloqueado.",
    424: "424 Failed Dependency – Dependencia falló.",
    425: "425 Too Early – Solicitud enviada demasiado pronto.",
    426: "426 Upgrade Required – Requiere actualizar el protocolo.",
    428: "428 Precondition Required – Se requieren precondiciones.",
    429: "429 Too Many Requests – Demasiadas solicitudes (rate limit).",
    431: "431 Request Header Fields Too Large – Encabezados demasiado grandes.",
    451: "451 Unavailable For Legal Reasons – Bloqueado por motivos legales.",
    500: "500 Internal Server Error – Error interno del servidor.",
    501: "501 Not Implemented – Función no soportada.",
    502: "502 Bad Gateway – Comunicación entre servidores falló.",
    503: "503 Service Unavailable – Servidor temporalmente no disponible.",
    504: "504 Gateway Timeout – Tiempo de espera agotado en el servidor.",
    505: "505 HTTP Version Not Supported – Versión HTTP no soportada.",
    506: "506 Variant Also Negotiates – Error de negociación de contenido.",
    507: "507 Insufficient Storage – Almacenamiento insuficiente.",
    508: "508 Loop Detected – Bucle infinito detectado.",
    510: "510 Not Extended – Requiere extensiones adicionales.",
    511: "511 Network Authentication Required – Autenticación de red requerida."
    }
    try:
        
        if http.status_code in activos:
            print(Fore.GREEN+activos[http.status_code])

        elif http.status_code in error:
            print(Fore.RED+error[http.status_code])
        else:
            print("codigo desconocido.")
            pass     
    except:
        print("Error en el request.")
        pass

    continuar = input("Y/N? ").strip().lower()
    if continuar not in ["y", "yes", "si", "s", "ok"]:
        break
