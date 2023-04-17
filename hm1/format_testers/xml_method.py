from .timer import TimerNs
import xmltodict
import dicttoxml
import sys

def xml_serialize(input : dict, need_size : bool = False):
    tm = TimerNs()
    xml_data =  dicttoxml.dicttoxml(input)
    ser_time = tm.finish()
    if need_size:
        return sys.getsizeof(xml_data)
    tm.start()
    dict_data = xmltodict.parse(xml_data)
    des_time = tm.finish()

    return (ser_time, des_time)
