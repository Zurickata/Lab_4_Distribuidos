module github.com/Zurickata/Lab_4_Distribuidos

go 1.22.2

replace (
	github.com/Zurickata/Lab_4_Distribuidos/mercenarios => ../mercenarios
	github.com/Zurickata/Lab_4_Distribuidos/proto => ../proto
)

require (
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240515191416-fc5f0ca64291 // indirect
)
