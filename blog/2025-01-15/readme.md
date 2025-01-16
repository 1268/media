# cover

~~~
go test -cover '-test.gocoverdir' cover -run Zero
go test -cover '-test.gocoverdir' cover -run One
go tool covdata textfmt -i cover -o cover.txt
go tool cover -html cover.txt

go test -coverprofile cover.txt
go tool cover -html cover.txt
~~~

https://go.dev/doc/build-cover#converting-to-legacy-text-format
