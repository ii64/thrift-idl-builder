namespace go facade.dummy1

include "./structs.thrift"

exception Dummy1Exception {}
service Dummy1Service {
    void ping() throws (1: Dummy1Exception e)
    list<structs.Job> getJobList() throws (1: Dummy1Exception e)
}