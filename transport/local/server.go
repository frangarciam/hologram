// Copyright 2014 AdRoll, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package local

import (
	"github.com/AdRoll/hologram/protocol"
	"net"
)

type server struct {
	s       net.Listener
	handler protocol.ConnectionHandlerFunc
}

/*
listen accepts incoming connections and hands them off to the
connection handler.
*/
func (us *server) listen() {
	for {
		conn, err := us.s.Accept()
		if err != nil {
			continue
		}

		smc := protocol.NewMessageConnection(conn)
		go us.handler(smc)
	}
}

/*
Close stops all further processing and releases the socket.
*/
func (us *server) Close() error {
	return us.s.Close()
}

/*
NewServer returns a server that listens on a UNIX socket, and
automatically starts that server.
*/
func NewServer(address string, handler protocol.ConnectionHandlerFunc) (retServer *server, err error) {
	socket, err := net.Listen("unix", address)
	if err != nil {
		return
	}

	retServer = &server{
		s:       socket,
		handler: handler,
	}

	go retServer.listen()
	return
}
