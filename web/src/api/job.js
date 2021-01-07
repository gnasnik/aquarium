import request from '../utils/request';

export const jobListReq = (query, token) => {
    return request({
        url:"/job/list",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const addJobReq = (data,token) => {
    return request({
        url:"/job/put",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}

export const delJobReq = (data,token) => {
    return request({
        url:"/job/del",
        method:"post",
        data:data,
        headers: {"Authorization":"Bearer "+token},
    })
}