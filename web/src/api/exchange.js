import request from '../utils/request';

export const exchangeListReq = (query, token) => {
    return request({
        url:"/exchange/list",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const exchangeTypesReq = (query, token) => {
    return request({
        url:"/exchange/types",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const addExchangeReq = (data,token) => {
    return request({
        url:"/exchange/put",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const delExchangeReq = (data,token) => {
    return request({
        url:"/exchange/del",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}