import request from '../utils/request';

export const fetchData = query => {
    return request({
        url: './table.json',
        method: 'get',
        params: query // form content-type
    });
};


export const loginReq = query => {
    return request({
        url:"/user/login",
        method:"post",
        data:query // json content-type
    })
}