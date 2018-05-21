go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/ksa/.cache/go-build"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/ksa/go"
GORACE=""
GOROOT="/usr/lib/go-1.10"
GOTMPDIR=""
GOTOOLDIR="/usr/lib/go-1.10/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build658993777=/tmp/go-build -gno-record-gcc-switches"

/go$ cd src

git init

git clone -b developer https://gitlab.com/protoplan/spider.git

go get github.com/gin-gonic/gin
go get github.com/golang-collections/collections/stack
go get github.com/gorilla/sessions
go get github.com/joho/godotenv
go get golang.org/x/net/html/charset 
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson


