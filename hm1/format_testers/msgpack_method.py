from .timer import TimerNs
import msgpack
import sys

def msgpack_serialize(input : dict, need_size : bool = False):
    tm = TimerNs()
    data = msgpack.packb(input)
    ser_time = tm.finish()
    if need_size:
        return sys.getsizeof(data)
    tm.start()
    data = msgpack.unpackb(data)
    des_time = tm.finish()

    return (ser_time, des_time)
