from .timer import TimerNs
import avro.schema
from avro.io import DatumWriter, BinaryEncoder, DatumReader, BinaryDecoder
from io import BytesIO
import sys

avro_schema = avro.schema.parse('''
{
    "namespace": "example.avro",
    "type": "record",
    "name": "dict",
    "fields": [
        {"name": "number", "type": "int"},
        {"name": "float", "type": "float"},
        {"name": "bool", "type": "boolean"},
        {"name": "ask", "type": "string"},
        {"name": "array_str", "type": {"type": "array", "items": "string"}},
        {"name": "dict", "type": {"type": "map", "values": "int"}}
    ]
}
''')

def avro_serialize(input : dict, need_size : bool = False):
    out = BytesIO()
    writer = DatumWriter(avro_schema)
    encoder = BinaryEncoder(out)
    tm = TimerNs()
    writer.write(input, encoder)
    data = out.getvalue()
    ser_time = tm.finish()
    if need_size:
        return sys.getsizeof(data)
    reader = DatumReader(avro_schema)
    tm.start()
    in_buf = BytesIO(data)
    decoder = BinaryDecoder(in_buf)
    data = reader.read(decoder)
    des_time = tm.finish()

    return (ser_time, des_time)
