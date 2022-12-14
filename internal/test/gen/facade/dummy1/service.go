// Code generated by thriftgo (0.2.1). DO NOT EDIT.

package dummy1

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ii64/thrift-idl-builder/internal/test/gen/facade/dummy1/structs"
)

type Dummy1Exception struct {
}

func NewDummy1Exception() *Dummy1Exception {
	return &Dummy1Exception{}
}

var fieldIDToName_Dummy1Exception = map[int16]string{}

func (p *Dummy1Exception) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

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

func (p *Dummy1Exception) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin(ctx, "Dummy1Exception"); err != nil {
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

func (p *Dummy1Exception) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy1Exception(%+v)", *p)
}
func (p *Dummy1Exception) Error() string {
	return p.String()
}

type Dummy1Service interface {
	Ping(ctx context.Context) (err error)

	GetJobList(ctx context.Context) (r []*structs.Job, err error)
}

type Dummy1ServiceClient struct {
	c thrift.TClient
}

func NewDummy1ServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *Dummy1ServiceClient {
	return &Dummy1ServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewDummy1ServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *Dummy1ServiceClient {
	return &Dummy1ServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewDummy1ServiceClient(c thrift.TClient) *Dummy1ServiceClient {
	return &Dummy1ServiceClient{
		c: c,
	}
}

func (p *Dummy1ServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *Dummy1ServiceClient) Ping(ctx context.Context) (err error) {
	var _args Dummy1ServicePingArgs
	var _result Dummy1ServicePingResult
	if _, err = p.Client_().Call(ctx, "ping", &_args, &_result); err != nil {
		return
	}
	return nil
}
func (p *Dummy1ServiceClient) GetJobList(ctx context.Context) (r []*structs.Job, err error) {
	var _args Dummy1ServiceGetJobListArgs
	var _result Dummy1ServiceGetJobListResult
	if _, err = p.Client_().Call(ctx, "getJobList", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.E != nil:
		return r, _result.E
	}
	return _result.GetSuccess(), nil
}

type Dummy1ServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Dummy1Service
}

func (p *Dummy1ServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *Dummy1ServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *Dummy1ServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewDummy1ServiceProcessor(handler Dummy1Service) *Dummy1ServiceProcessor {
	self := &Dummy1ServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("ping", &dummy1ServiceProcessorPing{handler: handler})
	self.AddToProcessorMap("getJobList", &dummy1ServiceProcessorGetJobList{handler: handler})
	return self
}
func (p *Dummy1ServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type dummy1ServiceProcessorPing struct {
	handler Dummy1Service
}

func (p *dummy1ServiceProcessorPing) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := Dummy1ServicePingArgs{}
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
	result := Dummy1ServicePingResult{}
	if err2 = p.handler.Ping(ctx); err2 != nil {
		switch v := err2.(type) {
		case *Dummy1Exception:
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

type dummy1ServiceProcessorGetJobList struct {
	handler Dummy1Service
}

func (p *dummy1ServiceProcessorGetJobList) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := Dummy1ServiceGetJobListArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "getJobList", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}

	iprot.ReadMessageEnd(ctx)
	result := Dummy1ServiceGetJobListResult{}
	var retval []*structs.Job
	if retval, err2 = p.handler.GetJobList(ctx); err2 != nil {
		switch v := err2.(type) {
		case *Dummy1Exception:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getJobList: "+err2.Error())
			oprot.WriteMessageBegin(ctx, "getJobList", thrift.EXCEPTION, seqId)
			x.Write(ctx, oprot)
			oprot.WriteMessageEnd(ctx)
			oprot.Flush(ctx)
			return true, thrift.WrapTException(err2)
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin(ctx, "getJobList", thrift.REPLY, seqId); err2 != nil {
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
	x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error on response getJobList: "+err2.Error())
	oprot.WriteMessageBegin(ctx, "getJobList", thrift.EXCEPTION, seqId)
	x.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	err = thrift.WrapTException(err2)
	return
}

type Dummy1ServicePingArgs struct {
}

func NewDummy1ServicePingArgs() *Dummy1ServicePingArgs {
	return &Dummy1ServicePingArgs{}
}

var fieldIDToName_Dummy1ServicePingArgs = map[int16]string{}

func (p *Dummy1ServicePingArgs) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

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

func (p *Dummy1ServicePingArgs) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
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

func (p *Dummy1ServicePingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy1ServicePingArgs(%+v)", *p)
}

type Dummy1ServicePingResult struct {
	E *Dummy1Exception `thrift:"e,1,optional" json:"e,omitempty"`
}

func NewDummy1ServicePingResult() *Dummy1ServicePingResult {
	return &Dummy1ServicePingResult{}
}

var Dummy1ServicePingResult_E_DEFAULT *Dummy1Exception

func (p *Dummy1ServicePingResult) GetE() (v *Dummy1Exception) {
	if !p.IsSetE() {
		return Dummy1ServicePingResult_E_DEFAULT
	}
	return p.E
}

var fieldIDToName_Dummy1ServicePingResult = map[int16]string{
	1: "e",
}

func (p *Dummy1ServicePingResult) IsSetE() bool {
	return p.E != nil
}

func (p *Dummy1ServicePingResult) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Dummy1ServicePingResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Dummy1ServicePingResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = NewDummy1Exception()
	if err := p.E.Read(ctx, iprot); err != nil {
		return err
	}
	return nil
}

func (p *Dummy1ServicePingResult) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
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

func (p *Dummy1ServicePingResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
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

func (p *Dummy1ServicePingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy1ServicePingResult(%+v)", *p)
}

type Dummy1ServiceGetJobListArgs struct {
}

func NewDummy1ServiceGetJobListArgs() *Dummy1ServiceGetJobListArgs {
	return &Dummy1ServiceGetJobListArgs{}
}

var fieldIDToName_Dummy1ServiceGetJobListArgs = map[int16]string{}

func (p *Dummy1ServiceGetJobListArgs) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

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

func (p *Dummy1ServiceGetJobListArgs) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin(ctx, "getJobList_args"); err != nil {
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

func (p *Dummy1ServiceGetJobListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy1ServiceGetJobListArgs(%+v)", *p)
}

type Dummy1ServiceGetJobListResult struct {
	Success []*structs.Job   `thrift:"success,0,optional" json:"success,omitempty"`
	E       *Dummy1Exception `thrift:"e,1,optional" json:"e,omitempty"`
}

func NewDummy1ServiceGetJobListResult() *Dummy1ServiceGetJobListResult {
	return &Dummy1ServiceGetJobListResult{}
}

var Dummy1ServiceGetJobListResult_Success_DEFAULT []*structs.Job

func (p *Dummy1ServiceGetJobListResult) GetSuccess() (v []*structs.Job) {
	if !p.IsSetSuccess() {
		return Dummy1ServiceGetJobListResult_Success_DEFAULT
	}
	return p.Success
}

var Dummy1ServiceGetJobListResult_E_DEFAULT *Dummy1Exception

func (p *Dummy1ServiceGetJobListResult) GetE() (v *Dummy1Exception) {
	if !p.IsSetE() {
		return Dummy1ServiceGetJobListResult_E_DEFAULT
	}
	return p.E
}

var fieldIDToName_Dummy1ServiceGetJobListResult = map[int16]string{
	0: "success",
	1: "e",
}

func (p *Dummy1ServiceGetJobListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *Dummy1ServiceGetJobListResult) IsSetE() bool {
	return p.E != nil
}

func (p *Dummy1ServiceGetJobListResult) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

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
		case 0:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField0(ctx, iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(ctx, fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Dummy1ServiceGetJobListResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Dummy1ServiceGetJobListResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return err
	}
	p.Success = make([]*structs.Job, 0, size)
	for i := 0; i < size; i++ {
		_elem := structs.NewJob()
		if err := _elem.Read(ctx, iprot); err != nil {
			return err
		}

		p.Success = append(p.Success, _elem)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return err
	}
	return nil
}

func (p *Dummy1ServiceGetJobListResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = NewDummy1Exception()
	if err := p.E.Read(ctx, iprot); err != nil {
		return err
	}
	return nil
}

func (p *Dummy1ServiceGetJobListResult) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin(ctx, "getJobList_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(ctx, oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
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

func (p *Dummy1ServiceGetJobListResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin(ctx, "success", thrift.LIST, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.Success)); err != nil {
			return err
		}
		for _, v := range p.Success {
			if err := v.Write(ctx, oprot); err != nil {
				return err
			}
		}
		if err := oprot.WriteListEnd(ctx); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(ctx); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *Dummy1ServiceGetJobListResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
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

func (p *Dummy1ServiceGetJobListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Dummy1ServiceGetJobListResult(%+v)", *p)
}
