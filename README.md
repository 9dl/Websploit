
# Websploit
![](/image.png)
A new exploit that could be used to gather links that are Vulnerable by using **RATTED** Users and using a Server scanning these Links with SQLMap (SQLI) or other Scanners like LFI.
## How it works

When someone uses the Browser Extension, it sends all links to a WebSocket. 
With these links, you can create a comprehensive database, filter out known sites, and scan sites for vulnerabilities such as LFI, SQLI, and more. These sites could be used for Database Attacks and much more.
## FAQ

#### How to protect against it?

Keep your Server up to date and search about how to prevent it:
- https://brightsec.com/blog/sql-injection-attack/
- https://brightsec.com/blog/lfi-attack-real-life-attacks-and-attack-examples/

## Tech Stack

**Client:** JS (Chrome Extension)

**Server:** Golang

## Info
- Websploit RAT => Server
- Websploit Extensions => Client (only Chromium Browsers Supported | Tested on: Brave/Chrome)
- Websploit UI => Server and Client (isnt working but was for Testing with UI)
This Project was coded Poorly and isnt Completed (Half done)


## Authors

- [@FaxHack](https://github.com/Faxhack) - Tester


## PoC Notice
This repository is a Proof of Concept (PoC) and is not intended for use in production environments. The provided content is for demonstration purposes only and may lack completeness. Users are cautioned to use it with care, and the creators make no warranties regarding its accuracy or reliability.

By accessing this repository, you acknowledge the experimental nature of the PoC and agree to use it at your own risk. The creators are not liable for any damages arising from its use. If you do not agree with these terms, refrain from using the PoC.
