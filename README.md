## Installation

1. Install Go 1.7.x or greater, godep, git, setup `$GOPATH` and `PATH=$PATH:$GOPATH/bin`, sqlite3

2. Install database
    ```
    cat migrations/db.sql | sqlite3 instagram.db
    ```

3. Run the server
    ```
    cd $GOPATH/src/github.com/maddevsio/instagram-agent
    godep restore
    go build .
    ./instagram-agent -httpAddr=:8090 -dashboardURL=http://localhost:8080/dashboard/v1/register
    ```

## Env usage
```
export PORT=8090
export DASHBOARD_URL="http://localhost:8080/dashboard/v1/register"
export CLIEND_ID="{INSTAGRAM CLIENT_ID}"
export ACCESS_TOKEN="{INSTAGRAM ACCESS_TOKEN}
```

## Flag usage
```
Usage of ./instagram-agent:
  -dashboardURL string
       	Dashboard service URL (default "http://localhost:8080/dashboard/v1/register")
  -httpAddr strin
       	HTTP listen address (default ":8090")
```
