namespace go Gateway

//--------------------request & response--------------
struct ServiceInfo {
    1: required string serviceName(api.body="name"),
    2: required string serviceIdlName(api.body="idlName"),
}

struct SuccessResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct ServiceReq {
    1: string serviceName(api.body='serviceName'),
}

service IdlService {
    SuccessResp AddService(1: ServiceInfo service)(api.post = '/add')
    SuccessResp DeleteService(1: ServiceReq serviceReq)(api.post = '/delete')
    SuccessResp UpdateService(1: ServiceInfo service)(api.post = '/update')
    ServiceInfo GetService(1: ServiceReq serviceReq)(api.post = '/get')
    list<ServiceInfo> ListService()(api.post = '/list')
}