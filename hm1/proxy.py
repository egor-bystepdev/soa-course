import socket
import asyncio
import os, random
import logging
from udp_tools.messages import *
from udp_tools import send_message
from udp_tools import parse_adrr
from format_testers import available_methods
from get_env import get_env_var

logging.basicConfig(level=logging.INFO)

def multicast_send(addr):
    pass


class UdpProtocol(asyncio.DatagramProtocol):
    def __init__(self, workers_port, multicast_addr, multicast_port):
        super().__init__()
        self.worker_port = workers_port
        self.multicast_addr = multicast_addr
        self.multicast_port = multicast_port

    def connection_made(self, transport):
        self.transport = transport

    def datagram_received(self, data, addr):
        data = data.decode("utf-8")
        logging.info('receive %s', data)
        request = deserialize_json(data)
        if (request is None or "type" not in request):
            return
        if (request["type"] == "get_result_all"):
            data_to_send = serialize_json({"type" : "get_result", "addr" : addr[0] + ":" + str(addr[1])})
            send_message.send_message(self.transport, data_to_send, self.multicast_addr, self.multicast_port) # change

        if (request["type"] == "get_result"):
            if ("methods" not in request or not isinstance(request["methods"], list)):
                return
            for method in request["methods"]:
                if not isinstance(method, str):
                    continue
                method = method.upper()
                if (method not in available_methods.avaliable_methods):
                    return
                logging.info("Send to %d", self.worker_port)
                data_to_send = serialize_json({"type" : "get_result", "addr" : addr[0] + ":" + str(addr[1])})
                send_message.send_message(self.transport, data_to_send, method, self.worker_port) # change

        if (request["type"] == "result"):
            logging.info("hello %s", addr[0])
            if ("addr" not in request or "result" not in request):
                return
            
            client_addr = parse_adrr.parse_address_port(request["addr"])
            if (client_addr is None):
                logging.warn('skip %s', data)
                return
            if (not isinstance(request["result"], str)):
                return
            logging.info("send to %s:%d", client_addr[0], client_addr[1])
            send_message.send_message(self.transport, request["result"], client_addr[0], client_addr[1])
        

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    PORT = int(get_env_var('PORT'))
    MULTICAST_ADDR = get_env_var('MULTICAST_ADDR')
    MULTICAST_PORT = int(get_env_var('MULTICAST_PORT'))
    logging.info('port %d', PORT)
    t = loop.create_datagram_endpoint(lambda: UdpProtocol(PORT, MULTICAST_ADDR, MULTICAST_PORT), local_addr=('0.0.0.0', PORT))
    loop.run_until_complete(t)
    loop.run_forever()

