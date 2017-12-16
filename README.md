# agenda-cs
[![Build Status](https://travis-ci.org/dzc15331066/agenda-cs.svg?branch=master "Travis CI status")](https://travis-ci.org/dzc15331066/agenda-cs)

an agenda project with cli and service

## use cli to test mock server

```
​```
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli 
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  agenda [command]

Available Commands:
  cm          A brief description of your command
  delUser     A brief description of your command
  help        Help about any command
  login       A brief description of your command
  logout      A brief description of your command
  qm          A brief description of your command
  query       A brief description of your command
  register    A brief description of your command

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ clear

sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli register -h
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  agenda register [flags]

Flags:
  -c, --contact string    Help message for phone
  -e, --email string      Help message for email
  -h, --help              help for register
  -p, --password string   Help message for password
  -u, --user string       Help message for username

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli register -u fangfang -p pass -c 123456 -e fangfang@qq.com
{
    "id": 1,
    "username":"zhang3",
    "password":"zhang",
    "email":"zhang3@mail2.sysu.edu.cn",
    "phone":"13256984563"
}
[success]:OK! register successfully!
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli login -h
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  agenda login [flags]

Flags:
  -h, --help              help for login
  -p, --password string   user -password [password] or -p [password]
  -u, --user string       use -user [username] or -u [username]

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli login -u fangfang -p pass
<nil>
Post https://private-ff790a-agenda25.apiary-mock.com/v1/user/state/login?key=1e3576bt: dial tcp 23.21.55.239:443: i/o timeout
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli login -u fangfang -p pass
<nil>
Post https://private-ff790a-agenda25.apiary-mock.com/v1/user/state/login?key=1e3576bt: dial tcp 204.236.236.192:443: i/o timeout
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli login -u fangfang -p pass
&{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Max-Age:[10] Content-Length:[60] Date:[Tue, 12 Dec 2017 15:05:32 GMT] X-Apiary-Ratelimit-Remaining:[119] Access-Control-Allow-Origin:[*] Access-Control-Allow-Methods:[OPTIONS,GET,HEAD,POST,PUT,DELETE,TRACE,CONNECT] Content-Type:[application/json] X-Apiary-Transaction-Id:[5a2ff03ccc154907008e9ce5] Via:[1.1 vegur] Server:[Cowboy] Connection:[keep-alive] X-Apiary-Ratelimit-Limit:[120]] 0xc4201ba0c0 60 [] false false map[] 0xc42000ad00 0xc4200b1080}
    {
        "success":"true"
        "state":"login"
    }
[success]: login successfully!
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli query -h
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  agenda query [flags]

Flags:
  -h, --help   help for query

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli query 
200 OK
[
    {
        "id":1,
        "username":"zhang3",
        "password":"zhang",
        "email":"zhang3@mail2.sysu.edu.cn",
        "phone":"13256984563"
    },{
        "id":2,
        "username":"li4",
        "password":"li",
        "email":"li4@mail2.sysu.edu.cn",
        "phone":"13256984563"
    }
]
[{1 zhang3 zhang zhang3@mail2.sysu.edu.cn 13256984563} {2 li4 li li4@mail2.sysu.edu.cn 13256984563}]
List all users:
1 zhang3 zhang3@mail2.sysu.edu.cn 13256984563
2 li4 li4@mail2.sysu.edu.cn 13256984563
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli cm -h
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  agenda cm [flags]

Flags:
  -e, --end string     use -end or -e [year-month-day] (default "2017-Jan-25")
  -h, --help           help for cm
  -p, --part strings   use -part [participators] or -p [participators]
  -s, --start string   use -start or -s [year-month-day] (default "2017-Jan-25")
  -t, --title string   use -title [meeting's title] or -t [meeting's title]

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli cm -t math -p fangfang manman -s 2017-12-24/23:23 -e 2017-12-24/24:00
[error]: time format should be like yyyy-mm-dd/hh:mm
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli cm -t math -p fangfang manman -s 2017-12-24/23:23 -e 2017-12-24/23:40
201 Created
{
    "id": 1,
    "username":"zhang3",
    "password":"zhang",
    "email":"zhang3@mail2.sysu.edu.cn",
    "phone":"13256984563"
}
[success]: Adding meeting successfully!
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli -qm -h
Error: unknown shorthand flag: 'q' in -qm
Usage:
  agenda [command]

Available Commands:
  cm          A brief description of your command
  delUser     A brief description of your command
  help        Help about any command
  login       A brief description of your command
  logout      A brief description of your command
  qm          A brief description of your command
  query       A brief description of your command
  register    A brief description of your command

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.

unknown shorthand flag: 'q' in -qm
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli -qm
Error: unknown shorthand flag: 'q' in -qm
Usage:
  agenda [command]

Available Commands:
  cm          A brief description of your command
  delUser     A brief description of your command
  help        Help about any command
  login       A brief description of your command
  logout      A brief description of your command
  qm          A brief description of your command
  query       A brief description of your command
  register    A brief description of your command

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.

unknown shorthand flag: 'q' in -qm
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ ./cli qm
qm called
200 OK
[
    {
        "sponsor":"zhang3",
        "title":"service computing",
        "start":"2017-08-05T08:40:51.620Z",
        "end":"2017-08-05T08:40:51.620Z",
        "participators": [
                "name1",
                "name2",
                "name3"
            ]

    }, 
    {
        "sponsor":"zhang3",
        "title":"service computing",
        "start":"2017-08-05T08:40:51.620Z",
        "end":"2017-08-05T08:40:51.620Z",
        "participators": [
                "name1",
                "name2",
                "name3"
            ]

    }

]
results:
1. service computing | zhang3 | [name1 name2 name3] | 2017-08-05T08:40:51.620Z | 2017-08-05T08:40:51.620Z 
2. service computing | zhang3 | [name1 name2 name3] | 2017-08-05T08:40:51.620Z | 2017-08-05T08:40:51.620Z 
sysuygm@sysuygm:~/golang-workspace/src/agenda-cs/cli$ 
​```
```

## use cli to test service

```
​```
sysuygm@sysuygm:~/golang-workspace/src/github.com/dzc15331066/agenda-cs/cli/test$ go test -v
=== RUN   TestUserRegister
--- PASS: TestUserRegister (0.12s)
=== RUN   TestUserLogin
--- PASS: TestUserLogin (0.00s)
=== RUN   TestUserLogout
--- PASS: TestUserLogout (0.00s)
=== RUN   TestListAllUsers
--- PASS: TestListAllUsers (0.00s)
=== RUN   TestDeleteUser
--- PASS: TestDeleteUser (0.14s)
=== RUN   TestAddMeeting
200 OK
--- PASS: TestAddMeeting (0.53s)
=== RUN   TestListAllMeetings
200 OK
--- PASS: TestListAllMeetings (0.00s)
PASS
ok  	github.com/dzc15331066/agenda-cs/cli/test	0.790s
sysuygm@sysuygm:~/golang-workspace/src/github.com/dzc15331066/agenda-cs/cli/test$ 

​```
```
##  download images and start service

We build an image in docker hub called deng123/agenda-cs

###  download images

```
$ docker pull deng123/agenda-cs
```

###  start agenda-server

```
$ sudo docker run -p 8080:8080 --name agenda-cs -d deng123/agenda-cs:latest
```

### use agenda client

```
$ sudo docker run -it --net host agenda-cs "sh"
# agenda
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  agenda [command]

Available Commands:
  cm          A brief description of your command
  delUser     A brief description of your command
  help        Help about any command
  login       A brief description of your command
  logout      A brief description of your command
  qm          A brief description of your command
  query       A brief description of your command
  register    A brief description of your command

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.
```



