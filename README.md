# graphQL_sample

This repository is simple golang web application for practice of graphQL.

**Note:** This application don't have database.
So application save data to temp memory.

## Setup

Please install [dep](https://github.com/golang/dep) and run following command.

```
dep init
dep ensure
```

## Run

```
go run main.go
```

## Sample request data

```console
# fetch userList
$ curl -H 'Content-Type:application/json' -X POST -d '{userList{userId, userName, description, photoURL, email}}' 'http://localhost:8080/graphql'

# createUser
$ curl -H 'Content-Type:application/json' -X POST -d 'mutation{createUser(userName:"mituba", description: "des", photoURL: "photo", email: "email"){userId, userName, description, photoURL, email}}' 'http://localhost:8080/graphql'

# fetch eventList
$ curl -H 'Content-Type:application/json' -X POST -d '{eventList{eventId}}' 'http://localhost:8080/graphql'

# createEvent
$ curl -H 'Content-Type:application/json' -X POST -d 'mutation{createEvent(user: {userId: "1", userName:"mituba", description: "des", photoURL: "photo", email: "email"}, eventName: "event", description: "des", location: "loc", startTime: "start", endTime: "end",){eventId, user{userName}}}' 'http://localhost:8080/graphql'
```
