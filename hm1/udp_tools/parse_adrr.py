import socket

def parse_address_port(addr_port):
    try:
        address, port = addr_port.split(':')
        socket.inet_pton(socket.AF_INET, address)
        if not (int(port) >= 0 and int(port) <= 65535):
            return None
        return (address, int(port))
    except:
        return None
