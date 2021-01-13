import request from '../utils/request';

export const logListReq = (query, token) => {
    return request({
        url:"/log/list",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}

