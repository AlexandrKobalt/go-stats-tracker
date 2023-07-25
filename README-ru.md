# GO-STATS-TRACKER
Утилита для сбора статистики работы сервера
## Как использовать?
#### 1. Импортируйте пакет
```golang
import "github.com/AlexandrKobalt/go-stats-tracker/pkg"
```
#### 2. Добавьте обработчик middleware в стек Mux middleware
```golang
// example of using with chi
r := chi.NewRouter()
r.Use(pkg.StatsMiddleware)
```
#### 3. Получите собранную статистику
```golang
stats, err := pkg.GetAllStats()
if err != nil {
	panic("it's just impossible, I can't believe it")
}
```
## Пример результата
```json
{
    "/getStats": {
        "totalRequestsCount": "14",
        "requestsFrequency": "0.54 per second",
        "averageProcessTime": "1 ms",
        "lastRequestTime": "2023-07-25 15:24:23"
    },
    "/post": {
        "totalRequestsCount": "508",
        "requestsFrequency": "6.52 per second",
        "averageProcessTime": "2.4 ms",
        "lastRequestTime": "2023-07-25 15:25:15"
    }
}
```
