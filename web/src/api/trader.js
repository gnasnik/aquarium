import request from '../utils/request';

export const traderListReq = (query, token) => {
    return request({
        url:"/trader/list",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const addTraderReq = (data,token) => {
    return request({
        url:"/trader/put",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const delTraderReq = (data,token) => {
    return request({
        url:"/trader/del",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const swithTrader = (data, token) =>  {
    return request({
        url:"/trader/switch",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}