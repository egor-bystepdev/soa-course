# Data serialization and deserialization

Сборка контейнера
``
    docker build . -f Dockerfile -t dofe71/coa_hm1
``

Или сразу запуск
``
    docker-compose up
    docker-compose down
``

Формат запросов :

Запрос всех методов
``
{"type" : "get_result_all"}
``

Запрос конкретного набора из списка ["NATIVE", "JSON", "XML", "AVRO", "PROTO", "MSGPACK", "YAML"] (не чувствительно к регистру)

``
{"type" : "get_result", "methods" : ["JSON", "XML"]}
``

Адрес прокси сервера 0.0.0.0:2000
Для удобства можно использовать 
``
spammer.py
``

