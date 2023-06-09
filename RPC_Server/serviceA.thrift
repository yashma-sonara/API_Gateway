namespace go api

struct Request {
    1: string userId
    2: string message
}

struct Response {
    1: string message
}

service serviceA {
    Response methodA(1: Request req),
    Response methodB(1: Request req),
    Response methodC(1: Request req)
}

service serviceB {
    Response methodA(1: Request req),
    Response methodB(1: Request req),
    Response methodC(1: Request req)
}