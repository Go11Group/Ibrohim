package handler

import pb "api-gateway/genproto"

type Handler struct {
	Weather   pb.WeatherServiceClient
	Transport pb.TransportServiceClient
}

func NewHandler(W pb.WeatherServiceClient, T pb.TransportServiceClient) *Handler {
	return &Handler{Weather: W, Transport: T}
}
