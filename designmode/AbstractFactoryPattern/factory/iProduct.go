package factory

type xiaomiPhone struct {
	Name  string
	price int
}

func GetXphone() *xiaomiPhone {
	return &xiaomiPhone{"xiaomi10s", 3000}
}

func (f *xiaomiPhone) GetName() string {
	return f.Name
}

type huaweiPhone struct {
	Name  string
	price int
}

func GetHphone() *huaweiPhone {
	return &huaweiPhone{"huaweip50", 3800}
}

func (f *huaweiPhone) GetName() string {
	return f.Name
}
