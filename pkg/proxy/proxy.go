package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/thaddeuscleo/gopetuah/config"
)

type Server struct {
	ln net.Listener
	c  config.Config
}

func New(c config.Config) *Server {
	return &Server{
		c: c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", s.c.Proxy.Port))
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

	buf := make([]byte, 1024)
	n, err := downstreamConn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(s.c.Upstreams))

	for name, upstream := range s.c.Upstreams {
		log.Println("Connecting to: ", name)
		upstreamConn, err := net.Dial("tcp", net.JoinHostPort(upstream.Host, upstream.Port))
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			defer upstreamConn.Close()
			defer wg.Done()
			go io.Copy(upstreamConn, strings.NewReader(string(buf[:n])))
			_, err := io.Copy(downstreamConn, upstreamConn)
			if err != nil {
				log.Println(err)
			}
		}()
	}

	wg.Wait()
}
