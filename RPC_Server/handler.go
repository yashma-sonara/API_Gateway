package main

import (
	"context"
	"log"
)

// ServiceAImpl implements the last service interface defined in the IDL.
type ServiceAImpl struct{}

// MethodA implements the ServiceAImpl interface.
func (s *ServiceAImpl) MethodA(ctx context.Context) (err error) {
	// TODO: Your code here...
	log.Println("Connected to serviceA, methodA")
	return nil
}

// MethodB implements the ServiceAImpl interface.
// Oneway methods are not guaranteed to receive 100% of the requests sent by the client.
// And the client may not perceive the loss of requests due to network packet loss.
// If possible, do not use oneway methods.
func (s *ServiceAImpl) MethodB(ctx context.Context) (err error) {
	// TODO: Your code here...
	log.Println("Connected to serviceA, methodB")
	return nil
}

// MethodC implements the ServiceAImpl interface.
func (s *ServiceAImpl) MethodC(ctx context.Context) (err error) {
	// TODO: Your code here...
	log.Println("Connected to serviceA, methodC")
	return nil
}
