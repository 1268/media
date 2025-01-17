proposal: cmd/covdata: html subcommand

if you need multi run coverage, it seems you currently have two options. first option:

~~~sh
go test -cover '-test.gocoverdir' cover -run Zero
go test -cover '-test.gocoverdir' cover -run One
go tool covdata textfmt -i cover -o cover.txt
go tool cover -html cover.txt
~~~

second option:

~~~sh
go test -coverprofile zero.txt -run Zero
go test -coverprofile one.txt -run One
# manually merge files
go tool cover -html merge.txt
~~~

with the first option, you are having to convert to the legacy format [1] before you can get HTML output. with the second option, is seems no current method is available for multiple runs using the text format, meaning user needs to manually combined the resulting files somehow. to that end, I propose adding a new subcommand:

~~~
go tool covdata html -i cover -o cover.html
~~~

1. https://go.dev/doc/build-cover#converting-to-legacy-text-format
