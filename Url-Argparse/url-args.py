import argparse, time
import httpx, json
from colorama import Fore, init
init(autoreset=True)
red=Fore.RED
yellow=Fore.YELLOW
blue=Fore.BLUE
green=Fore.GREEN
parser = argparse.ArgumentParser(description="Url-Checker with arguments")
parser.add_argument("url", help="")
parser.add_argument("--status", "-S", help="an argument to view the status code", action="store_true")
parser.add_argument("--header", "-H", help="an argument for setting custom headers (json)")
parser.add_argument("--html", help="an argument to see the html response", action="store_true")
parser.add_argument("--Not-follow-redirect", "-nfr", help="an argument to not follow redirects (may give 301 in the status code)", action="store_true")
args= parser.parse_args()
if args.header:
    try:
        with open(args.header, "r") as f:
            _headers=json.load(f)
    except FileNotFoundError:
        print(red+"file not found check the path")
        pass
elif not args.header:
    _headers={
        
  "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
  "Accept-Encoding": "gzip, deflate, br, zstd",
  "Accept-Language": "es-ES,es;q=0.9,en;q=0.8",
  "Cache-Control": "no-cache",
  "Connection": "keep-alive",
  "Upgrade-Insecure-Requests": "1",
  "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
  "Sec-Ch-Ua": "\"Chromium\";v=\"128\", \"Google Chrome\";v=\"128\", \"Not;A=Brand\";v=\"24\"",
  "Sec-Ch-Ua-Mobile": "?0",
  "Sec-Ch-Ua-Platform": "\"Windows\"",
  "Sec-Fetch-Dest": "document",
  "Sec-Fetch-Mode": "navigate",
  "Sec-Fetch-Site": "none",
  "Sec-Fetch-User": "?1"


    }
redirects=True
if args.Not_follow_redirect:
    redirects = False
request=httpx.get(url=args.url, headers=_headers, follow_redirects=redirects)
if args.status == True:
    print(green+f"{args.url} --> {request.status_code}")
    time.sleep(0.01)
if args.html == True:
    request.encoding = "utf-8"
    time.sleep(0.01)
    print(request.text)

elif not args.html and not args.status:
    print(green+f"{args.url} --> {request}")
