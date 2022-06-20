let socket = null;

document.addEventListener('DOMContentLoaded', () => {
    socket = new WebSocket("ws://192.168.1.155:8080/ws");
    socket.onopen = () => {
        console.log("Successfully connected");
    }
    socket.onclose = () => {
        console.log("Connection closed");
    }
    socket.onerror = (error) => {
        console.log("Error: " + error);
    }
    socket.onmessage = (message) => {
        console.log("Message: " + message.data);
    }
});
