// Code generated by thriftgo (0.2.1). DO NOT EDIT.

package structs

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ii64/thrift-idl-builder/internal/test/gen/facade/dummy1/enums"
)

type Job struct {
	Status    enums.Status `thrift:"status,1" json:"status"`
	Name      string       `thrift:"name,2" json:"name"`
	CreatedAt int64        `thrift:"createdAt,3" json:"createdAt"`
}

func NewJob() *Job {
	return &Job{}
}

func (p *Job) GetStatus() (v enums.Status) {
	return p.Status
}

func (p *Job) GetName() (v string) {
	return p.Name
}

func (p *Job) GetCreatedAt() (v int64) {
	return p.CreatedAt
}

var fieldIDToName_Job = map[int16]string{
	1: "status",
	2: "name",
	3: "createdAt",
}

func (p *Job) Read(ctx context.Context, iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(ctx, iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(ctx, fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(ctx, iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(ctx, fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(ctx, iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Job[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Job) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return err
	} else {
		p.Status = enums.Status(v)
	}
	return nil
}

func (p *Job) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return err
	} else {
		p.Name = v
	}
	return nil
}

func (p *Job) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return err
	} else {
		p.CreatedAt = v
	}
	return nil
}

func (p *Job) Write(ctx context.Context, oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin(ctx, "Job"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(ctx, oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(ctx, oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(ctx, oprot); err != nil {
			fieldId = 3
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

func (p *Job) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin(ctx, "status", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(ctx, int32(p.Status)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(ctx); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Job) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin(ctx, "name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(ctx, p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(ctx); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Job) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin(ctx, "createdAt", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(ctx, p.CreatedAt); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(ctx); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Job) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Job(%+v)", *p)
}