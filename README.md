# Pholcus [![GitHub release](https://img.shields.io/github/release/henrylee2cn/pholcus.svg?style=flat-square)](https://github.com/henrylee2cn/pholcus/releases) [![report card](https://goreportcard.com/badge/github.com/henrylee2cn/pholcus?style=flat-square)](http://goreportcard.com/report/henrylee2cn/pholcus) [![github issues](https://img.shields.io/github/issues/henrylee2cn/pholcus.svg?style=flat-square)](https://github.com/henrylee2cn/pholcus/issues?q=is%3Aopen+is%3Aissue) [![github closed issues](https://img.shields.io/github/issues-closed-raw/henrylee2cn/pholcus.svg?style=flat-square)](https://github.com/henrylee2cn/pholcus/issues?q=is%3Aissue+is%3Aclosed) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/henrylee2cn/pholcus) [![view Go大数据](<https://img.shields.io/badge/官方QQ群-Go大数据(42731170)-27a5ea.svg?style=flat-square>)](http://jq.qq.com/?_wv=1027&k=XnGGnc)

Pholcus is a purely high-concurrency, heavyweight crawler software written in pure Go language. It is targeted at Internet data collection and provides a function that only requires attention to rule customization for people with a certain Go or JS programming foundation. A powerful reptile tool.

It supports three operating modes: stand-alone, server, and client. It has three operation interfaces: Web, GUI, and command line. The rules are simple and flexible, batch tasks are concurrent, and output methods are rich (mysql/mongodb/kafka/csv/excel, etc.). There is a large amount of Demo sharing; in addition, it supports two horizontal and vertical capture modes, supporting a series of advanced functions such as simulated login and task pause and cancel.

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/icon.png)

# Reptile principle

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/module.png)

&nbsp;

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/project.png)

&nbsp;

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/distribute.png)

#Frame Features
 1. Provide heavyweight reptile tools that only need to pay attention to custom rules and functions for users with a certain Go or JS programming foundation;

2. Support three operating modes: stand-alone, server, and client;

3. GUI (Windows), Web, Cmd three operation interfaces, which can be controlled by parameters;

4. Support state control, such as pause, resume, stop, etc.

5. Control the amount of collection;

6. Control the number of concurrent associations;

7. Support multiple acquisition tasks concurrently;

8. Support proxy IP list, which can control the frequency of replacement;

9. Support the collection process to stop randomly and simulate artificial behavior;

10. Provide custom configuration input interface according to rule requirements

11. There are five output modes: mysql, mongodb, kafka, csv, excel, and original file download.

12. Support batch output, and the quantity of each batch is controllable;

13. Supports both static Go and dynamic JS collection rules, supports both horizontal and vertical capture modes, and has a large number of demos;

14. Persistence of successful records for automatic de-duplication;

15. Serialization failure request, support deserialization automatic reload processing;

16. Surfer high concurrent downloader, support GET / POST / HEAD method and http / https protocol, support fixed UserAgent automatic save cookie and random large number of UserAgent disable cookie two modes, highly simulate browser behavior, can achieve analog login, etc. Features;

17. The server/client mode adopts the Teleport high concurrent SocketAPI framework, which is connected by full-duplex long-length communication, and the internal data transmission format is JSON.

&nbsp;

# Go version requirements

≥Go1.6

&nbsp;

# Download and install

```
go get -u -v github.com/henrylee2cn/pholcus
```

Note: Pholcus publicly maintained spider rule base address <https://github.com/henrylee2cn/pholcus_lib>

&nbsp;

# Create project

```
package main

import (
    "github.com/henrylee2cn/pholcus/exec"
    _ "github.com/henrylee2cn/pholcus_lib" // This is a publicly maintained spider rule base
     // _ "pholcus_lib_pte" // You can also freely add your own rule base
)

func main() {
     // Set the runtime default action interface and start running
     // Before running the software, you can set the -a_ui parameter to "web", "gui" or "cmd" to specify the operation interface for this run.
     // where "gui" only supports Windows systems
    exec.DefaultRun("web")
}
```

&nbsp;

# Compile and run

Normal compilation method

```
cd {{replace your gopath}}/src/github.com/henrylee2cn/pholcus
go install or go build
```

Hidden cmd window compilation method under Windows

```
cd {{replace your gopath}}/src/github.com/henrylee2cn/pholcus
go install -ldflags="-H windowsgui" or go build -ldflags="-H windowsgui"
```

View optional parameters:

```
pholcus -h
```

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/help.jpg)

&nbsp;

> _<font size="2">The screenshot of the web version of the operation interface is as follows:_

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/webshow_1.png)

&nbsp;

> _<font size="2">GUI mode operation interface mode selection interface screenshot is as follows_

![image](https://github.com/henrylee2cn/pholcus/raw/master/doc/guishow_0.jpg)

&nbsp;

> _<font size="2">Cmd version of the operating parameter settings example is as follows_

```
$ pholcus -_ui=cmd -a_mode=0 -c_spider=3,8 -a_outtype=csv -a_thread=20 -a_dockercap=5000 -a_pause=300
-a_proxyminute=0 -a_keyins="<pholcus><golang>" -a_limit=10 -a_success=true -a_failure=true
```

&nbsp;

\*Note: If you use the proxy IP function under Mac, be sure to obtain the root user right, otherwise you can't get it through `ping`!

&nbsp;

# Runtime catalog file

```
├─pholcus software
│
├─pholcus_pkg runtime file directory
│  ├─config.ini configuration file
│  │
│  ├─proxy.lib proxy IP list file
│  │
│  ├─spiders dynamic rule directory
│  │  └─xxx.pholcus.html Dynamic Rules File
│  │
│  ├─phantomjs program file
│  │
│  ├─text_out text data file output directory
│  │
│  ├─file_out result output directory
│  │
│  ├─logs directory
│  │
│  ├─history history directory
│  │
└─└─cache temporary cache directory
```

&nbsp;

# Dynamic rule example

Features: Dynamic loading rules, no need to recompile the software, simple writing, free to add, suitable for lightweight collection projects.
<br/>
xxx.pholcus.html

```
<Spider>
    <Name>HTML dynamic rule example</Name>
    <Description>HTML dynamic rule example [Auto Page] [http://xxx.xxx.xxx]</Description>
    <Pausetime>300</Pausetime>
    <EnableLimit>false</EnableLimit>
    <EnableCookie>true</EnableCookie>
    <EnableKeyin>false</EnableKeyin>
    <NotDefaultField>false</NotDefaultField>
    <Namespace>
        <Script></Script>
    </Namespace>
    <SubNamespace>
        <Script></Script>
    </SubNamespace>
    <Root>
        <Script param="ctx">
        console.log("Root");
        ctx.JsAddQueue({
            Url: "http://xxx.xxx.xxx",
            Rule: "LoginPage"
        });
        </Script>
    </Root>
    <Rule name="LoginPage">
        <AidFunc>
            <Script param="ctx,aid">
            </Script>
        </AidFunc>
        <ParseFunc>
            <Script param="ctx">
            console.log(ctx.GetRuleName());
            ctx.JsAddQueue({
                Url: "http://xxx.xxx.xxx",
                Rule: "AfterLogin",
                Method: "POST",
                PostData: "username=44444444@qq.com&amp;password=44444444&amp;login_btn=login_btn&amp;submit=login_btn"
            });
            </Script>
        </ParseFunc>
    </Rule>
    <Rule name="AfterLogin">
        <ParseFunc>
            <Script param="ctx">
            console.log(ctx.GetRuleName());
            ctx.Output({
                "All": ctx.GetText()
            });
            ctx.JsAddQueue({
                Url: "http://accounts.xxx.xxx/member",
                Rule: "PersonalCenter",
                Header: {
                    "Referer": [ctx.GetUrl()]
                }
            });
            </Script>
        </ParseFunc>
    </Rule>
    <Rule name="PersonalCenter">
        <ParseFunc>
            <Script param="ctx">
            console.log("PersonalCenter: " + ctx.GetRuleName());
            ctx.Output({
                "All": ctx.GetText()
            });
            </Script>
        </ParseFunc>
    </Rule>
</Spider>
```

# Static rule example

Features: Compiled with the software, more customized, more efficient, suitable for heavyweight acquisition projects.
<br/>
xxx.go

```
func init() {
    Spider{
        Name:        "Static rule example",
        Description: "Static rule example [Auto Page] [http://xxx.xxx.xxx]",
        // Pausetime: 300,
        // Limit:   LIMIT,
        // Keyin:   KEYIN,
        EnableCookie:    true,
        NotDefaultField: false,
        Namespace:       nil,
        SubNamespace:    nil,
        RuleTree: &RuleTree{
            Root: func(ctx *Context) {
                ctx.AddQueue(&request.Request{Url: "http://xxx.xxx.xxx", Rule: "LoginPage"})
            },
            Trunk: map[string]*Rule{
                "LoginPage": {
                    ParseFunc: func(ctx *Context) {
                        ctx.AddQueue(&request.Request{
                            Url:      "http://xxx.xxx.xxx",
                            Rule:     "AfterLogin",
                            Method:   "POST",
                            PostData: "username=123456@qq.com&password=123456&login_btn=login_btn&submit=login_btn",
                        })
                    },
                },
                "AfterLogin": {
                    ParseFunc: func(ctx *Context) {
                        ctx.Output(map[string]interface{}{
                            "All": ctx.GetText(),
                        })
                        ctx.AddQueue(&request.Request{
                            Url:    "http://accounts.xxx.xxx/member",
                            Rule:   "PersonalCenter",
                            Header: http.Header{"Referer": []string{ctx.GetUrl()}},
                        })
                    },
                },
                "PersonalCenter": {
                    ParseFunc: func(ctx *Context) {
                        ctx.Output(map[string]interface{}{
                            "All": ctx.GetText(),
                        })
                    },
                },
            },
        },
    }.Register()
}
```

&nbsp;

# Proxy IP

- Proxy IP written in `/pholcus_pkg/proxy.lib` The file is in the following format, one IP per line:

```
http://183.141.168.95:3128
https://60.13.146.92:8088
http://59.59.4.22:8090
https://180.119.78.78:8090
https://222.178.56.73:8118
http://115.228.57.254:3128
http://49.84.106.160:9000
```

- In the operation interface, select "Proxy IP replacement frequency" or set the `-a_proxyminute` parameter on the command line to use.

- \*Note: If you use the proxy IP function under Mac, be sure to obtain the root user permission. Otherwise, you can not get the proxy through ping!

&nbsp;

# FAQ

In the request queue, will duplicate URLs be automatically de-duplicated?

```
The url is de-duplicated by default, but it can be ignored by setting Request.Reloadable=true.
```

If the content of the page pointed to by the URL is updated, does the framework have a mechanism for judging?

```
The content of the url page is updated. The framework cannot directly support the judgment, but the user can customize the support in the rules.
```

The success of the request is based on the status code of the web header.

```
Instead of judging the state, it is determined whether the server has a response stream or not. That is, the 404 page is also a success.
```

The re-request mechanism after the request fails?

```
After each url attempts to download a specified number of times, if it still fails, the request is appended to a special queue like defer.
After the current task ends normally, it will be automatically added to the download queue and downloaded again. If there are still no downloads successfully, save to the failure history.
The next time you execute the crawler rule, you can automatically add these failed requests to the special queue of the defer property by selecting the inheritance history failure record... (following the steps)
```

&nbsp;

# Third-party dependencies

```
"github.com/henrylee2cn/teleport"
"golang.org/x/net/html/charset"
"gopkg.in/mgo.v2"
"github.com/robertkrimen/otto"
"github.com/Shopify/sarama"
"github.com/go-sql-driver/mysql"
"github.com/lxn/walk"
"github.com/elazarl/go-bindata-assetfs"
"github.com/henrylee2cn/pholcus_lib" // This is a publicly maintained spider rule base
```

> _<font size="2">(Thanks for the support of the above open source project!)</font>_

# Open source agreement

Pholcus(Ghost Spider) project is commercially friendly[Apache License v2](https://github.com/henrylee2cn/pholcus/raw/master/LICENSE).release
