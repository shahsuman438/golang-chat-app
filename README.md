# GoLang CLI based Chat Application 

## Git

Clone this repository

```sh
$ git clone https://github.com/shahsuman438/golang-chat-app.git
```

## Server Setup

Move to the server folder

```sh
 $ cd server
```

Start Application

```sh
 $ go run main.go
```
Now you should be able to access backend server on http://localhost:8000/.


## Client Setup

Move to the Client folder

```sh
 $ cd client
```

Build Application

```sh
 $ go build
```

Now you can login and register to start chating

## Use cmd to LOGIN

```sh
 $ ./client login <email> <password> --room=<room name>
```
default room name would be "public"
## Use cmd to REGISTER

```sh
 $ ./client register <email> <user name> <password>
```




