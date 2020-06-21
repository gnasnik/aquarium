import request from '../utils/request';

export const algorithmListReq = (query, token) => {
    return request({
        url:"/algorithm/list",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const addAlgorithmReq = (data,token) => {
    return request({
        url:"/algorithm/put",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const delAlgorithmReq = (data,token) => {
    return request({
        url:"/algorithm/del",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}