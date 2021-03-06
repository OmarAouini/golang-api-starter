# golang-api-starter

## scripts

- running all tests with coverage:

```sh

./build/test.sh
```

it will also generate a .out coverage file.

- build project

```sh
./build/build.sh
```

- generate swagger docs

```sh
./build/gen_swagger.sh
```

to access the docs, run the app and go to: 
[swagger docs](http://localhost:8080/swagger/index.html)

- generate test Mocks

```sh
./build/gen_mock.sh
```

it is also needed to add the interface that you want to mock in the script everytime a new one is being created, like this:

```sh
mockery --recursive --name=InterfaceToMock
```

it will generate the interface mocks in the "mocks" folder.
