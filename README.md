# GO-STATS-TRACKER
framework for collecting server operation statistics
## How to use?
#### 1. Import package
```golang
import "github.com/AlexandrKobalt/go-stats-tracker/pkg"
```
#### 2. Add a middleware handler to the Mux middleware stack
```golang
// example of using with chi
r := chi.NewRouter()
r.Use(pkg.StatsMiddleware)
```
#### 3. Get the collected statistics
```golang
stats, err := pkg.GetAllStats()
if err != nil {
	panic("it's just impossible, I can't believe it")
}
```
## Result example
```json
{
    "/getStats": {
        "totalRequestsCount": "14",
        "requestsFrequency": "0.54 per second",
        "averageProcessTime": "0 ms",
        "lastRequestTime": "2023-07-25 15:24:23"
    },
    "/post": {
        "totalRequestsCount": "508",
        "requestsFrequency": "6.52 per second",
        "averageProcessTime": "0 ms",
        "lastRequestTime": "2023-07-25 15:25:15"
    }
}
```
