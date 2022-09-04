namespace go facade.dummy2

exception Dummy2Exception {}
service Dummy2Service {
    void ping() throws (1: Dummy2Exception e)
}