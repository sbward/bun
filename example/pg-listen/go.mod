module github.com/uptrace/bun/example/pg-listen

go 1.17

replace github.com/uptrace/bun => ../..

replace github.com/uptrace/bun/extra/bundebug => ../../extra/bundebug

replace github.com/uptrace/bun/driver/pgdriver => ../../driver/pgdriver

replace github.com/uptrace/bun/dialect/pgdialect => ../../dialect/pgdialect

require (
	github.com/uptrace/bun v1.1.3
	github.com/uptrace/bun/dialect/pgdialect v1.1.3
	github.com/uptrace/bun/driver/pgdriver v1.1.3
	github.com/uptrace/bun/extra/bundebug v1.1.3
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
	golang.org/x/sys v0.0.0-20220328115105-d36c6a25d886 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	mellium.im/sasl v0.2.1 // indirect
)
