package biz

import "com.lee/config-web/entity"

var RpcServerBiz = new(rpcServerBiz)

type rpcServerBiz struct {
}

func (biz *rpcServerBiz) GetRpcServices(name string, pageIndex int, pageSize int) (servers []entity.Service, totalCount int, err error) {
	//r, err = etcdx.GetClient().Get("service", true, true)
	return nil, 0, nil
}
