# thrift-idl-builder

Apache Thrift IDL builder.

## Example

Command-line example:

```bash
go run github.com/ii64/thrift-idl-builder \
    -errors \
    -wrk 10 \
    -source-dir ./idl \
    -o ./internal/test/gen \
    -bin thriftgo \
    -gen go
```

Or a minimal `Makefile` example.

```make
all: gen-stub-ext

PKG := leanon-tg/app/infra/service-discovery

THRIFT_LIB := github.com/apache/thrift/lib/go/thrift
THRIFT_GEN_PACKAGE_PREFIX := $(PKG)/pkg/gen/

THRIFT_IDL_REMOTE_PATTERN := *remote

THRIFTGO_GEN_FLAG := thrift_import_path=$(THRIFT_LIB),package_prefix=$(THRIFT_GEN_PACKAGE_PREFIX)
# thriftgo flag
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),reorder_fields=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),frugal_tag=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),keep_unknown_fields=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),reserve_comments=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),nil_safe=false
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),compatible_names=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),gen_type_meta=true

THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),value_type_in_container=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),validate_set=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),use_type_alias=true
THRIFTGO_GEN_FLAG := $(THRIFTGO_GEN_FLAG),gen_db_tag=true
#
THRIFTGO_GEN := go:"$(THRIFTGO_GEN_FLAG)"

# base
.PHONY: gen-stub
gen-stub:
   go run github.com/ii64/thrift-idl-builder \
      -errors \
      -wrk 10 \
      -tp httpclient \
      -source-dir $(THRIFT_DIR_SRC) \
      -o $(THRIFT_DIR_OUT) \
      -bin thriftgo \
      -gen $(THRIFTGO_GEN) && \
   bash -c 'find $(THRIFT_DIR_OUT) -name "$(THRIFT_IDL_REMOTE_PATTERN)" -prune -exec bash -c "echo {} && rm -r {}" \;' && \
   echo OK


# Extension
THRIFT_IDL_EXT_SRC := ./idl
THRIFT_IDL_EXT_OUT := ./pkg/gen
.PHONY: gen-stub-ext
gen-stub-ext:
   $(MAKE) gen-stub THRIFT_DIR_SRC=$(THRIFT_IDL_EXT_SRC) THRIFT_DIR_OUT=$(THRIFT_IDL_EXT_OUT)

```
