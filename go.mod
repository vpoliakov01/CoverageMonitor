module github.com/vpoliakov01/CoverageMonitor

go 1.13

require (
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.1
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
)

replace github.com/vpoliakov01/CoverageMonitor/back_end/server => ./back_end/server
