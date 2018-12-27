package lib

import (
	"github.com/holdno/snowFlakeByGo"
)

//twitter snowFlake算法的golang实现
//github抄的,线程安全

var id_worket *snowFlakeByGo.Worker

func init() {
	registerSnowFlake()
}

func registerSnowFlake() {
	id_worket, _ = snowFlakeByGo.NewWorker(0) // 传入当前节点id 此id在机器集群中一定要唯一 且从0开始排最多1024个节点，可以根据节点的不同动态调整该算法每毫秒生成的id上限(如何调整会在后面讲到)
}

//得到一个uid,线程安全的，不用怕
func GetUid() (int64) {
	return id_worket.GetId()
}
