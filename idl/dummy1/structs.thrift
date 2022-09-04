namespace go facade.dummy1.structs
include "./enums.thrift"

struct Job {
    1: enums.Status status
    2: string name
    3: i64 createdAt
}