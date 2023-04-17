from .timer import TimerNs
import json
import sys

def native_serialize(input : dict, need_size : bool = False):
    tm = TimerNs()
    my_dict_str = str(input)
    ser_time = tm.finish()
    if need_size:
        return sys.getsizeof(my_dict_str)
    tm.start()
    new_dict = eval(my_dict_str)
    des_time = tm.finish()

    return (ser_time, des_time)
