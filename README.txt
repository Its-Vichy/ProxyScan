🔎 scan the internet to find "private" proxies.

🧠 HTTP/SOCKS4/SOCKS5 Proxies.

📌 Installation:
  - sudo apt-get install git zmap golang-go -y (debian based distribution)
  - git clone https://github.com/Its-Vichy/ProxyScan && cd ProxyScan
  
  👉 Scan for proxies:
  - zmap -p <proxy_port> | go run ./main.go <proxy_port> <output_file>
  
  👉 Check ips from file (without port):
  - cat <proxy_file> | go run ./main.go <proxy_port> <output_file>

📌 Default port:
   - Socks: 1080
   - Http:  3128 | 8080


🚨 You need working proxies ?

  We are providing the best proxies at cheap price !

✅ https://docs.proxies.gay
✅ https://vu.fr/rca-discord
