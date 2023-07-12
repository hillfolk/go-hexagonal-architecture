package web

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}
func (s *HttpServer) Run() error {
	r := NewRouter()
	if err := r.Run(":" + "8080"); err != nil {
		return err
	}
	return nil
}
