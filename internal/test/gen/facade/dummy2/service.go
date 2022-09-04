// Code generated by thriftgo (0.2.1). DO NOT EDIT.

package dummy2

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type Dummy2Exception struct {
}

func NewDummy2Exception() *Dummy2Exception {
	return &Dummy2Exception{}
}

var fieldIDToName_Dummy2Exception = map[int16]string{}

func (p *Dummy2Exception) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(ctx); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin(ctx)
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(ctx, fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(ctx); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(ctx); err != nil {
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

func (p *Dummy2Exception) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin(ctx, "Dummy2Exception"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(ctx); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(ctx); err != nil {
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

func (p *Dummy2Exception) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy2Exception(%+v)", *p)
}
func (p *Dummy2Exception) Error() string {
	return p.String()
}

type Dummy2Service interface {
	Ping(ctx context.Context) (err error)
}

type Dummy2ServiceClient struct {
	c thrift.TClient
}

func NewDummy2ServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *Dummy2ServiceClient {
	return &Dummy2ServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewDummy2ServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *Dummy2ServiceClient {
	return &Dummy2ServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewDummy2ServiceClient(c thrift.TClient) *Dummy2ServiceClient {
	return &Dummy2ServiceClient{
		c: c,
	}
}

func (p *Dummy2ServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *Dummy2ServiceClient) Ping(ctx context.Context) (err error) {
	var _args Dummy2ServicePingArgs
	var _result Dummy2ServicePingResult
	if _, err = p.Client_().Call(ctx, "ping", &_args, &_result); err != nil {
		return
	}
	return nil
}

type Dummy2ServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Dummy2Service
}

func (p *Dummy2ServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *Dummy2ServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *Dummy2ServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewDummy2ServiceProcessor(handler Dummy2Service) *Dummy2ServiceProcessor {
	self := &Dummy2ServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("ping", &dummy2ServiceProcessorPing{handler: handler})
	return self
}
func (p *Dummy2ServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x
}

type dummy2ServiceProcessorPing struct {
	handler Dummy2Service
}

func (p *dummy2ServiceProcessorPing) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := Dummy2ServicePingArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "ping", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}

	iprot.ReadMessageEnd(ctx)
	result := Dummy2ServicePingResult{}
	if err2 = p.handler.Ping(ctx); err2 != nil {
		switch v := err2.(type) {
		case *Dummy2Exception:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: "+err2.Error())
			oprot.WriteMessageBegin(ctx, "ping", thrift.EXCEPTION, seqId)
			x.Write(ctx, oprot)
			oprot.WriteMessageEnd(ctx)
			oprot.Flush(ctx)
			return true, thrift.WrapTException(err2)
		}
	}
	if err2 = oprot.WriteMessageBegin(ctx, "ping", thrift.REPLY, seqId); err2 != nil {
		goto reply_err
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		goto reply_err
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		goto reply_err
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		goto reply_err
	}
	if err != nil {
		return
	}
	return true, err
reply_err:
	x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error on response ping: "+err2.Error())
	oprot.WriteMessageBegin(ctx, "ping", thrift.EXCEPTION, seqId)
	x.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	err = thrift.WrapTException(err2)
	return
}

type Dummy2ServicePingArgs struct {
}

func NewDummy2ServicePingArgs() *Dummy2ServicePingArgs {
	return &Dummy2ServicePingArgs{}
}

var fieldIDToName_Dummy2ServicePingArgs = map[int16]string{}

func (p *Dummy2ServicePingArgs) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(ctx); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin(ctx)
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(ctx, fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(ctx); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(ctx); err != nil {
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

func (p *Dummy2ServicePingArgs) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin(ctx, "ping_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(ctx); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(ctx); err != nil {
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

func (p *Dummy2ServicePingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy2ServicePingArgs(%+v)", *p)
}

type Dummy2ServicePingResult struct {
	E *Dummy2Exception `thrift:"e,1,optional" json:"e,omitempty"`
}

func NewDummy2ServicePingResult() *Dummy2ServicePingResult {
	return &Dummy2ServicePingResult{}
}

var Dummy2ServicePingResult_E_DEFAULT *Dummy2Exception

func (p *Dummy2ServicePingResult) GetE() (v *Dummy2Exception) {
	if !p.IsSetE() {
		return Dummy2ServicePingResult_E_DEFAULT
	}
	return p.E
}

var fieldIDToName_Dummy2ServicePingResult = map[int16]string{
	1: "e",
}

func (p *Dummy2ServicePingResult) IsSetE() bool {
	return p.E != nil
}

func (p *Dummy2ServicePingResult) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(ctx); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin(ctx)
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(ctx, iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(ctx, fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(ctx, fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(ctx); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(ctx); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Dummy2ServicePingResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Dummy2ServicePingResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = NewDummy2Exception()
	if err := p.E.Read(ctx, iprot); err != nil {
		return err
	}
	return nil
}

func (p *Dummy2ServicePingResult) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin(ctx, "ping_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(ctx, oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(ctx); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(ctx); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Dummy2ServicePingResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err = oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(ctx); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Dummy2ServicePingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy2ServicePingResult(%+v)", *p)
}