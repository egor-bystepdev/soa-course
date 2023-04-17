from .json_method import json_serialize
from .native_method import native_serialize
from .xml_method import xml_serialize
from .avro_method import avro_serialize
from .yaml_method import yaml_serialize
from .msgpack_method import msgpack_serialize
from .proto_method import proto_serialize
import os

def get_method_for_testing(name):
    if (name == "JSON"):
        return json_serialize
    if (name == "NATIVE"):
        return native_serialize
    if (name == "XML"):
        return xml_serialize
    if (name == "AVRO"):
        return avro_serialize
    if (name == "YAML"):
        return yaml_serialize
    if (name == "MSGPACK"):
        return msgpack_serialize
    if (name == "PROTO"):
        return proto_serialize
    raise Exception(f"No method with name {name}")

