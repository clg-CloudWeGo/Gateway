namespace go hertzSvr.idlManager

struct IDLMessage {
    1: string svcName(api.query="svcName")
}

struct IDLResponse {
    1: bool success(api.body="success")
    2: string message(api.body="message")
}

service IDLService {
    IDLResponse updateIDL(1:IDLMessage req)(api.post="/idl/update")
}
