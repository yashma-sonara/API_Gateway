namespace go api

struct Request {
    1: string userId
    2: string message
}

service serviceA {
    void methodA(1: Request req),
    void methodB(1: Request req),
    void methodC(1: Request req)
}