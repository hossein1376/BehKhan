package config

type services struct {
	Catalogue catalogue `json:"catalogue"`
}

type catalogue struct {
	Address  string `json:"address"`
	HttpPort string `json:"http_port"`
	GrpcPort string `json:"grpc_port"`
}
