🔎 scan the internet to find "private" proxies.

Installation:
  - sudo apt-get install git zmap golang-go -y (debian based distribution)
  - git clone https://github.com/Its-Vichy/ProxyScan && cd ProxyScan
  
  👉 Scan for proxies:
  - zmap -p <proxy_port> | go run ./main.go <proxy_port> <output_file>
  
  👉 Check ips from file (without port):
  - cat <proxy_file> | go run ./main.go <proxy_port> <output_file>
