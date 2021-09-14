**imgoptimizer** - утилита создания изображений, оптимизированных для web

Собирать следует такой командой:
```
go build -a --tags netgo --ldflags '-extldflags "-static -lm"'
```

Зависит от [liwebp](https://developers.google.com/speed/webp)

Может записать результат в БД (MySQL)
