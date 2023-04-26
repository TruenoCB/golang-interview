package factory

type xiaomi struct {
}

func GetXiaomiFactory() *xiaomi {
	return &xiaomi{}
}

func (f *xiaomi) MakePhone() absProduct {
	return GetXphone()
}

type huawei struct {
}

func GetHuaweiFactory() *huawei {
	return &huawei{}
}

func (f *huawei) MakePhone() absProduct {
	return GetHphone()
}
