namespace go facade.dummy1

exception Dummy1Exception {}
service Dummy1Service {
    void ping() throws (1: Dummy1Exception e)
}