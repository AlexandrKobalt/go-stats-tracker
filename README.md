# GO-STATS-TRACKER
## framework for collecting server operation statistics
### How to use?
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
