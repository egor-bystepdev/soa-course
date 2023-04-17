import socket

def send_message(sock, message, host, port) -> None:
    sock.sendto(message.encode(), (host, port))
