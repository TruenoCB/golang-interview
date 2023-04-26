package factory

import "fmt"

type absFactory interface {
	MakePhone() absProduct
}

func GetFactory(fname string) (absFactory, error) {
	if fname == "xiaomi" {
		return GetXiaomiFactory(), nil
	}
	if fname == "huawei" {
		return GetHuaweiFactory(), nil
	}
	return nil, fmt.Errorf("创建工厂失败")
}
