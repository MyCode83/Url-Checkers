import customtkinter as ctk
import httpx
import datetime
from tkinter import filedialog
from tkinter import messagebox
from fake_useragent import UserAgent
ua=UserAgent()
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

root = ctk.CTk()
#settings
ctk.set_appearance_mode("Light")
root.minsize(900, 512)
root.title("Url-Checker")


#frame principal
frame=ctk.CTkFrame(master=root)
frame.pack(side="right", expand=True)
#frame
barra=ctk.CTkFrame(master=frame)
barra.pack(fill="x", padx=10, pady=10, side="bottom")
#frame arriba
barra_a = ctk.CTkFrame(master=frame)
barra_a.pack(side="top", fill="x")
#dark/light
def modo():
    if ctk.get_appearance_mode() == "Light":
        
        ctk.set_appearance_mode("Dark")
        _modo.configure(text="Modo Claro")
    else:
        
        ctk.set_appearance_mode("Light")
        _modo.configure(text="Modo oscuro")
#dark/light button
_modo=ctk.CTkButton(master=barra_a, text="Modo oscuro", command=modo)
_modo.pack(pady=20, padx=3, side="right")

#textbox & file
urls=ctk.CTkTextbox(master=frame, width=500)
urls.pack(pady=80, fill="both", expand=True, padx=(99, 99))

def txt():
    archivo=filedialog.askopenfilename(title="Urls TXT", filetypes=[("Archivos txt", "*.txt")])
    if archivo:
        with open(archivo, "r", encoding="utf-8") as f:
            texto=f.read()
            urls.delete("1.0", "end")
            urls.insert("1.0", texto)
txtboton=ctk.CTkButton(master=barra, command=txt, text="Urls(txt)")

txtboton.pack(side="right", pady=5, padx=5)

#output
output= ctk.CTkScrollableFrame(master=root, corner_radius=1)
output.pack(side="left", fill="both", expand=True, padx=15, pady=15)
auto_guardar= ctk.BooleanVar(value=True)
def nombre():
    fecha = datetime.datetime.now().strftime("%Y-%m-%d")
    return f"logs-{fecha}"
def _output(_texto):
    out=ctk.CTkLabel(master=output, text=_texto, anchor="w", justify="left")
    out.pack(anchor="w", pady=4, padx=6)
    if auto_guardar.get():
        archivo=nombre()
        with open(archivo, "a", encoding="utf-8") as f:
            f.write(_texto + "\n")
def poner():
    contenido=urls.get("1.0", "end")
    if not contenido:
        _output("El contenido esta vacio.")
    for _url in contenido.splitlines():
        try:
            peticion=httpx.get(url=_url, follow_redirects=True, headers=_headers)
            _output(f"{_url} --> {peticion.status_code}")
        except Exception as e:
            _output(f"Error en la peticion: {_url} --> {e}")
enviar=ctk.CTkButton(master=barra, text="enviar", command=poner)
enviar.pack(side="left", pady=5, padx=5)

def borrar():
    confirmacion=messagebox.askyesno("Confirmación", "¿Seguro que quieres borrar los logs?")
    if confirmacion == True:
        for label in output.winfo_children():
            label.destroy()
def guardar():
    archivo= filedialog.asksaveasfilename(title="Guardar logs como ...", defaultextension=".txt", filetypes=[("Archivos de texto", "*.txt"), ("logs", "*.log")])
    if archivo:
        with open(archivo, "w", encoding="utf-8") as f:
            for log in output.winfo_children():
                if isinstance(log, ctk.CTkLabel):
                    f.write(log.cget("text") + "\n")
        messagebox.showinfo("Logs guardados", f"Los logs se han guardado {archivo}")
guardar_boton=ctk.CTkButton(text="Guardar logs", command=guardar, master=barra)
guardar_boton.pack(side="right", pady=5, padx=5)
borrar_boton=ctk.CTkButton(master=barra, text="Borrar logs", command=borrar)
borrar_boton.pack(side="right", padx=5, pady=10)
def esconder():

    if auto_guardar.get():
        guardar_boton.pack_forget()
    else:
        guardar_boton.pack(side="right", pady=5, padx=5)
auto_guardar_checkbox = ctk.CTkCheckBox(
    master=barra_a,
    text="Guardar automáticamente",
    variable=auto_guardar,
    onvalue=True,
    offvalue=False,
    command=esconder
)
auto_guardar_checkbox.pack(side="right", pady=10)

auto_guardar_checkbox.select()
guardar_boton.pack_forget() 
root.mainloop()
