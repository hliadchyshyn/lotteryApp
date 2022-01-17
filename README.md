# lotteryApp

to start app run command
```
go run .
```

Simple web service that receives email to identify user and gives response:
- with HTTP 200 if the request is successful;
- with HTTP 410 if out of tickets;
- with HTTP 403 if users try to issue a ticket again.

Example of request will be validated in app
```
{
    "email": *Valid email*
}
```
Request should contain headers:
```
"Content-Type": "application/json"
```
In case of deploy check ```.env``` file to identify which vars should be set.
