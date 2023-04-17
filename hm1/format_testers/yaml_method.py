from .timer import TimerNs
import yaml
import sys

def yaml_serialize(input : dict, need_size : bool = False):
    tm = TimerNs()
    yaml_str = yaml.dump(input)
    ser_time = tm.finish()
    if need_size:
        return sys.getsizeof(yaml_str)
    tm.start()
    data = yaml.load(yaml_str, Loader=yaml.FullLoader)
    des_time = tm.finish()

    return (ser_time, des_time)
