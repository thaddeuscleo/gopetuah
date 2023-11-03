package proxy

import (
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

type Server struct {
	listenAddr string
	ln         net.Listener
}

func New(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	s.acceptConnectionLoop()

	return nil
}

func (s *Server) acceptConnectionLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.readConnectionLoop(conn)
	}
}

func (s *Server) readConnectionLoop(downstreamConn net.Conn) {
	defer downstreamConn.Close()

	upstreams := []string{
		"localhost:9000",
		"localhost:8000",
	}

	buf := make([]byte, 1024)
	n, err := downstreamConn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(upstreams))

	for idx, upstreamUrl := range upstreams {
		upstreamConn, err := net.Dial("tcp", upstreamUrl)
		if err != nil {
			log.Println(err)
		}
		go func(i int) {
			defer upstreamConn.Close()
			defer wg.Done()
			go io.Copy(upstreamConn, strings.NewReader(string(buf[:n])))
			log.Println("Sended to upstream ", i)
			_, err := io.Copy(downstreamConn, upstreamConn)
			if err != nil {
				log.Println(err)
			}
		}(idx)
	}

	wg.Wait()
}
