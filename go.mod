module madr

go 1.14

replace github.com/pramineni01/madr/configsvc => ./configsvc

require (
	github.com/pramineni01/madr/configSvc v1.0.0
	google.golang.org/grpc v1.28.1 // indirect
)
