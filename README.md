# Usage

```
	func main(){

		app := gi.New()

		app.Handler(func(ctx *gi.Context){
			// handler 
		})

		app.Go("https://baidu.com")
	}
```
