from urllib.parse import urlparse, parse_qs
import json

def serialize_json(dct):
    return json.dumps(dct)

def deserialize_json(s):
    try:
        dictionary = json.loads(s)
        return dictionary
    except:
        return None
