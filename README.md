
# image-previewer

Сервис предназначен для изготовления preview (создания изображения с новыми размерами на основе имеющегося изображения).

## Возможности
- Обрезка изображений по заданным размерам
- Выбор одного из возможного типа для обработки `fill`, `fit`, `anchor`,`resize`

## Развертывание
```sh
make run
```
Сервис будет доступен по адресу:
```sh
127.0.0.1:8080
```

## Использование

Пример для запроса:

http://localhost:8080/fill/200/300/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg

Где:
- `fill` один из возможных типов обработки изображений (`fill`, `fit`, `anchor`,`resize`)
- `200/300` ширина и высота изображения соответственно
- `raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg` исходное изображение


## Конфигурация docker

Пример конфигурации для запуска сервиса `example.env`

## Конфигурация  сервиса

Конфигурационный файл находится по следующему пути: `configs/config.toml`

Где `capacity` количество сохраненных оригиналов изображений в кэш

## Зависимости

| Пакет | Репозиторий |
| ------ | ------ |
| Imaging | https://github.com/disintegration/imaging |
| toml | https://github.com/BurntSushi/toml |


