namespace go hertzSvr.service

struct IDLInfo{
    1: string name (api.body="name")
    2: string idl (api.body="idl")
}

struct IDLMessage{
    1: string name (api.body="name")
}

struct IDLResponse{
    1: bool success (api.body="success")
    2: string message (api.body="message")
}

struct IDLQueryReq {
    1: string name (api.query="name")
}

service IDLService {
    IDLResponse AddIDL(1: IDLInfo idl)(api.post="/idl/add")
    IDLResponse DeleteIDL(1: IDLMessage idl)(api.post="/idl/delete")
    IDLResponse UpdateIDL(1: IDLInfo idl)(api.post="/idl/update")
    IDLInfo QueryIDL(1: IDLQueryReq idl)(api.get="/idl/query")
}