import request from '../utils/request';

export const traderListReq = (query, token) => {
    return request({
        url:"/trader/list",
        method:"get",
        params:query,
        headers: {"Authorization":"Bearer "+token},
    })
}
