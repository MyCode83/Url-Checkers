import sys, httpx
from colorama import Fore, init
from fake_useragent import UserAgent
init()

try:
    _urls=sys.argv[1]
except IndexError:
    print(Fore.RED+
    """[USO DEL SCRIPT - COMPROBADOR DE URLs]

    Este programa comprueba el estado (200, 404, etc.) de múltiples URLs.

    ▶ Cómo usarlo:

    1. Crea un archivo .txt con una URL por línea. Ejemplo:

        https://google.com  
        https://github.com  
        https://midominiofalso123.com  

    2. Ejecuta el script pasando el nombre del archivo como parámetro, por ejemplo:

        python url_checker.py urls.txt

    3. El script leerá las URLs y te mostrará el estado de cada una.

    ▶ Si no pasas ningún archivo, no se podrá hacer la comprobación.

    """
    )
    exit()
x=UserAgent()
UA=x.chrome()
_headers = {
    "User-Agent": UA,
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
    "Accept-Language": "es-ES,es;q=0.9,en;q=0.8",
    "Accept-Encoding": "gzip, deflate, br",
    "Connection": "keep-alive",
    "Upgrade-Insecure-Requests": "1",
    "Sec-Fetch-Dest": "document",
    "Sec-Fetch-Mode": "navigate",
    "Sec-Fetch-Site": "none",
    "Sec-Fetch-User": "?1"
}
try:
    with open(_urls, "r") as f:
         urls=f.read()
    for _url in urls.splitlines():
        if not _url:
            continue
        
        peticion=httpx.get(url=_url, headers=_headers, follow_redirects=True)
        print(Fore.GREEN+f"{_url} -->  {peticion.status_code}")
except FileNotFoundError:
    print(Fore.RED+"""
[ERROR: ARCHIVO NO ENCONTRADO]

El archivo que has especificado no existe o no se pudo abrir.

▶ Posibles causas:
  - Escribiste mal el nombre o la ruta del archivo.
  - El archivo no está en la misma carpeta que este script.
  - La ruta es incorrecta o contiene errores.
  - El archivo aún no ha sido creado.

▶ Soluciones:
  - Verifica que el archivo exista y que la ruta sea correcta.
  - Puedes usar un nombre directo si está en la misma carpeta (ej: urls.txt),
    o una ruta completa si está en otra ubicación.

▶ Ejemplos correctos de uso:
  python url_checker.py urls.txt
  python url_checker.py ./datos/mis_urls.txt
  python url_checker.py "C:/Users/usuario/Escritorio/urls_web.txt"
""")
    exit()
