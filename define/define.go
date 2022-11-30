package define

var (
	DefaultPage = "1"
	DefaultSize = "20"
	StatusMsg   = map[int]string{
		-1: "等待评测",
		0:  "评测中",
		1:  "评测通过",
		2:  "评测失败",
		3:  "超时",
		4:  "内存超限",
		5:  "编译错误",
	}
)
