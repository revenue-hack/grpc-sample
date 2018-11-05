package service

import (
	"sync"
	"github.com/revenue-hack/grpc-sample/person/protocol"
	"context"
)

type PersonService struct {
	customers []*protocol.Person
	m         sync.Mutex
}

func (s *PersonService) ListPerson(p *protocol.RequestType, stream protocol.PersonService_ListPersonServer) error {
	s.m.Lock()
	defer s.m.Unlock()
	for _, p := range s.customers {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func (s *PersonService) AddPerson(c context.Context, p *protocol.Person) (*protocol.ResponseType, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.customers = append(s.customers, p)
	return new(protocol.ResponseType), nil
}