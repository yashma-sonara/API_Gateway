// Code generated by thriftgo (0.2.11). DO NOT EDIT.

package api

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type ServiceA interface {
	MethodA(ctx context.Context) (err error)

	MethodB(ctx context.Context) (err error)

	MethodC(ctx context.Context) (err error)
}

type ServiceAClient struct {
	c thrift.TClient
}

func NewServiceAClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ServiceAClient {
	return &ServiceAClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewServiceAClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ServiceAClient {
	return &ServiceAClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewServiceAClient(c thrift.TClient) *ServiceAClient {
	return &ServiceAClient{
		c: c,
	}
}

func (p *ServiceAClient) Client_() thrift.TClient {
	return p.c
}

func (p *ServiceAClient) MethodA(ctx context.Context) (err error) {
	var _args ServiceAMethodAArgs
	var _result ServiceAMethodAResult
	if err = p.Client_().Call(ctx, "methodA", &_args, &_result); err != nil {
		return
	}
	return nil
}
func (p *ServiceAClient) MethodB(ctx context.Context) (err error) {
	var _args ServiceAMethodBArgs
	if err = p.Client_().Call(ctx, "methodB", &_args, nil); err != nil {
		return
	}
	return nil
}
func (p *ServiceAClient) MethodC(ctx context.Context) (err error) {
	var _args ServiceAMethodCArgs
	var _result ServiceAMethodCResult
	if err = p.Client_().Call(ctx, "methodC", &_args, &_result); err != nil {
		return
	}
	return nil
}

type ServiceAProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ServiceA
}

func (p *ServiceAProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ServiceAProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ServiceAProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewServiceAProcessor(handler ServiceA) *ServiceAProcessor {
	self := &ServiceAProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("methodA", &serviceAProcessorMethodA{handler: handler})
	self.AddToProcessorMap("methodB", &serviceAProcessorMethodB{handler: handler})
	self.AddToProcessorMap("methodC", &serviceAProcessorMethodC{handler: handler})
	return self
}
func (p *ServiceAProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type serviceAProcessorMethodA struct {
	handler ServiceA
}

func (p *serviceAProcessorMethodA) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ServiceAMethodAArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("methodA", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ServiceAMethodAResult{}
	if err2 = p.handler.MethodA(ctx); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing methodA: "+err2.Error())
		oprot.WriteMessageBegin("methodA", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("methodA", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type serviceAProcessorMethodB struct {
	handler ServiceA
}

func (p *serviceAProcessorMethodB) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ServiceAMethodBArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.MethodB(ctx); err2 != nil {
		return true, err2
	}
	return true, nil
}

type serviceAProcessorMethodC struct {
	handler ServiceA
}

func (p *serviceAProcessorMethodC) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ServiceAMethodCArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("methodC", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ServiceAMethodCResult{}
	if err2 = p.handler.MethodC(ctx); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing methodC: "+err2.Error())
		oprot.WriteMessageBegin("methodC", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("methodC", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ServiceAMethodAArgs struct {
}

func NewServiceAMethodAArgs() *ServiceAMethodAArgs {
	return &ServiceAMethodAArgs{}
}

func (p *ServiceAMethodAArgs) InitDefault() {
	*p = ServiceAMethodAArgs{}
}

var fieldIDToName_ServiceAMethodAArgs = map[int16]string{}

func (p *ServiceAMethodAArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ServiceAMethodAArgs) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("methodA_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ServiceAMethodAArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceAMethodAArgs(%+v)", *p)
}

func (p *ServiceAMethodAArgs) DeepEqual(ano *ServiceAMethodAArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type ServiceAMethodAResult struct {
}

func NewServiceAMethodAResult() *ServiceAMethodAResult {
	return &ServiceAMethodAResult{}
}

func (p *ServiceAMethodAResult) InitDefault() {
	*p = ServiceAMethodAResult{}
}

var fieldIDToName_ServiceAMethodAResult = map[int16]string{}

func (p *ServiceAMethodAResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ServiceAMethodAResult) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("methodA_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ServiceAMethodAResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceAMethodAResult(%+v)", *p)
}

func (p *ServiceAMethodAResult) DeepEqual(ano *ServiceAMethodAResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type ServiceAMethodBArgs struct {
}

func NewServiceAMethodBArgs() *ServiceAMethodBArgs {
	return &ServiceAMethodBArgs{}
}

func (p *ServiceAMethodBArgs) InitDefault() {
	*p = ServiceAMethodBArgs{}
}

var fieldIDToName_ServiceAMethodBArgs = map[int16]string{}

func (p *ServiceAMethodBArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ServiceAMethodBArgs) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("methodB_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ServiceAMethodBArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceAMethodBArgs(%+v)", *p)
}

func (p *ServiceAMethodBArgs) DeepEqual(ano *ServiceAMethodBArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type ServiceAMethodCArgs struct {
}

func NewServiceAMethodCArgs() *ServiceAMethodCArgs {
	return &ServiceAMethodCArgs{}
}

func (p *ServiceAMethodCArgs) InitDefault() {
	*p = ServiceAMethodCArgs{}
}

var fieldIDToName_ServiceAMethodCArgs = map[int16]string{}

func (p *ServiceAMethodCArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ServiceAMethodCArgs) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("methodC_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ServiceAMethodCArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceAMethodCArgs(%+v)", *p)
}

func (p *ServiceAMethodCArgs) DeepEqual(ano *ServiceAMethodCArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type ServiceAMethodCResult struct {
}

func NewServiceAMethodCResult() *ServiceAMethodCResult {
	return &ServiceAMethodCResult{}
}

func (p *ServiceAMethodCResult) InitDefault() {
	*p = ServiceAMethodCResult{}
}

var fieldIDToName_ServiceAMethodCResult = map[int16]string{}

func (p *ServiceAMethodCResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ServiceAMethodCResult) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("methodC_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ServiceAMethodCResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceAMethodCResult(%+v)", *p)
}

func (p *ServiceAMethodCResult) DeepEqual(ano *ServiceAMethodCResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}
