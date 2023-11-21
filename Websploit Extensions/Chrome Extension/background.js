const browserName = "Chrome";

let socket = null;

function connectToServer() {
    socket = new WebSocket('ws://localhost:8081/ws');

    socket.onopen = function () {
        socket.send(`BrowserConnected|${browserName}`);
    };

    socket.onerror = function (error) {
        console.error("WebSocket Error:", error);
    };

    socket.onclose = function (event) {
        if (event.wasClean) {
            console.log("WebSocket closed cleanly, code=" + event.code + " reason=" + event.reason);
        } else {
            console.error("Connection died");
            setTimeout(connectToServer, 1000);
        }
    };
}

chrome.webNavigation.onCompleted.addListener(function (details) {
    if (socket && socket.readyState === WebSocket.OPEN && details.tabId > -1 && details.url.startsWith("http")) {
        const message = 'VisitedURL:' + details.url;
        socket.send(message);
    }
});

chrome.runtime.onSuspend.addListener(function () {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send('BrowserDisconnected');
    }
});

connectToServer();
