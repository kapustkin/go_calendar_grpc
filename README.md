# GRPC сервис

## Задание

1. Создать GRPC спецификацию для сервиса-календаря из дз 2.7
2. Сгенерировать GRPC сервер и клиент, проверить работу сервиса
3. Отделить модель данных от Protobuf структур(?)

## Памятка

* Чтобы работал gomod вместе с gRPC

```console
go get github.com/golang/protobuf@master
```

* Для хранения всех пакетов в папке вендор
  
```console
go mod vendor
```

* Для генерации протофайлов

```console
PS D:\Documents\kapustkin\protoc\bin> .\protoc -I=d:\Documents\kapustkin\go_calendar_grpc\calendarpb\ --go_out=plugins=grpc:d:\Documents\kapustkin\go_calendar_grpc\calendarpb\ calendar.proto
```
