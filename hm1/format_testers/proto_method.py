from .timer import TimerNs
from . import dict_pb2 as dict__pb2
import sys

# $ protoc --python_out=. mydata.proto

def fill_data(data):
    mydata = dict__pb2.MyData()
    mydata.number = data["number"]
    mydata.float = data["float"]
    mydata.bool = data["bool"]
    mydata.ask = data["ask"]
    mydata.array_str.extend(data["array_str"])
    for key, value in data["dict"].items():
        mydata.dict[key] = value
    return mydata

def proto_serialize(input : dict, need_size : bool = False):
    input = fill_data(input)
    tm = TimerNs()
    serialized_data = input.SerializeToString()
    ser_time = tm.finish()
    if need_size:
        return sys.getsizeof(serialized_data)
    mydata = dict__pb2.MyData()
    tm.start()
    mydata.ParseFromString(serialized_data)
    des_time = tm.finish()

    return (ser_time, des_time)
