# Url-Argparse
> A Python script that checks the status and content of a URL using `argparse`, with support for showing the status code, the **HTML** response, and using custom **JSON** headers.
## Usage
> python3 url-args.py **URL** *args*
- Use `--status` or `-S` to check the status code
- Use `--html` to view the html responde
- Use `--headers` or `-H` to load custom headers, for example: `python3 url-args.py https://example.com -H headers.json --status`
### Requirements
> `python 3`
> `httpx`
> `colorama`
`python3 -m pip install colorama httpx`
