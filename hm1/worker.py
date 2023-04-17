import socket
import asyncio
import os, random
from format_testers.get_method import get_method_for_testing
from udp_tools import parse_adrr
from format_testers.available_methods import avaliable_methods
from udp_tools import send_message
from udp_tools.messages import *
from get_env import get_env_var

data = { 
    "number" : 42,
    "float" : 0.1345,
    "bool" : True,
    "ask" : "answer",
    "array_str": ["c++", "code", "matter"],
    "dict": {"one" : 1, "five" : 5, "three" : 3}
}

def do_test(name, method):
    times = []
    for _ in range(1000):
        times.append(method(data))
    size = method(data, True)
    result_times = [sum(x) for x in zip(*times)]
    serialize_time = int(result_times[0] / len(times))
    deserialize_time = int(result_times[1] / len(times))

    return f"{name} - {size} - {serialize_time}ns - {deserialize_time}ns;"


HOST, PORT = 'localhost', 514

def MessageHandler(transport, method, type_name, data, addr):
    data = data.decode("utf-8")
    request = deserialize_json(data)
    if (request is None or "addr" not in request or "type" not in request
        or request["type"] != "get_result"):
        return
    
    parsed = parse_adrr.parse_address_port(request["addr"])
    if (parsed is not None):
        data_to_send = do_test(type_name, method)
        send_message.send_message(transport, serialize_json({"type": "result",
            "addr" : request["addr"], "result" : data_to_send}),addr[0], addr[1])

class UdpProtocol(asyncio.DatagramProtocol):
    def __init__(self, format_type):
        super().__init__()
        self.method = get_method_for_testing(format_type) 
        self.type_name = format_type

    def connection_made(self, transport):
        self.transport = transport

    def datagram_received(self, data, addr):
        MessageHandler(self.transport, self.method, self.type_name, data, addr)
    
class MulticastProtocol(asyncio.DatagramProtocol):
    def __init__(self, group, port, format_type):
        self.transport = None
        self.group = group
        self.port = port
        self.method = get_method_for_testing(format_type) 
        self.type_name = format_type

    def connection_made(self, transport):
        self.transport = transport
        sock = transport.get_extra_info('socket')
        sock.setsockopt(socket.IPPROTO_IP, socket.IP_ADD_MEMBERSHIP, socket.inet_aton(self.group) + socket.inet_aton('0.0.0.0'))
        print(f"Started listening on multicast address {self.group}:{self.port}")

    def datagram_received(self, data, addr):
        MessageHandler(self.transport, self.method, self.type_name, data, addr)

    def connection_lost(self, exc):
        print("Connection lost")

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    PORT = int(get_env_var('PORT'))
    TYPE = get_env_var('TYPE')
    if TYPE not in avaliable_methods:
        raise f"No format type {TYPE}, methods {avaliable_methods}"
    MULTICAST_ADDR = get_env_var('MULTICAST_ADDR')
    MULTICAST_PORT = int(get_env_var('MULTICAST_PORT'))

    t = loop.create_datagram_endpoint(lambda: UdpProtocol(TYPE), local_addr=('0.0.0.0', PORT))
    loop.run_until_complete(t)
    print("Hello")
    protocol = MulticastProtocol(MULTICAST_ADDR, MULTICAST_PORT, TYPE)
    t = loop.create_datagram_endpoint(lambda: protocol, local_addr=("0.0.0.0", 9999))
    loop.run_until_complete(t)
    loop.run_forever()

