import asyncio
import json

class UDPServerProtocol(asyncio.DatagramProtocol):
    def connection_made(self, transport):
        self.transport = transport

    def datagram_received(self, data, addr):
        print(f"Received {data.decode()} from {addr}")

    async def send_message(self, message, host, port):
        self.transport.sendto(message.encode(), (host, port))

async def main():
    loop = asyncio.get_running_loop()
    transport, protocol = await loop.create_datagram_endpoint(
        lambda: UDPServerProtocol(),
        local_addr=("0.0.0.0", 1234)
    )
    print(f"Server started on {transport.get_extra_info('sockname')}")

    # Цикл для периодической отправки сообщений
    mt = ["NATIVE", "JSON", "XML", "AVRO", "PROTO", "MSGPACK", "YAML"]
    import random
    while True:
        m = random.choice(mt)
        await protocol.send_message(json.dumps({"type" : "get_result_all"}), "0.0.0.0", 5777)
        await asyncio.sleep(1)

asyncio.run(main())
