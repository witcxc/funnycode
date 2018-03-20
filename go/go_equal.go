package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ss struct {
	ID int
}
type s struct {
	ID   ss
	name string
}

type DiskSpec struct {
	VolumeName           string            `json:",omitempty"`
	AllowUseHostBootDisk bool              `json:",omitempty"`
	Size                 int64             //:61440000000 //单位：字节。磁盘空间大小
	Type                 string            `json:",omitempty"` //:"SSD", "SATA"
	Iops                 int               `json:",omitempty"` //单位：次/秒。应用单实例的磁盘iops配额， NC的iops能力要大于这个值才能被选中；如果后面有iops隔离，那么这个值就是磁盘iops的quota值
	Iobps                int64             `json:",omitempty"` //单位：字节。应用单实例的磁盘带宽配额， NC的io吞吐能力要大于这个值才能被选中；如果后面有磁盘io带宽隔离，那么这个值就是磁盘带宽的quota值
	Rm                   string            `json:",omitempty`  //ro：readonly， rw：read write 含义等同于docker -v /a=/b:rw中冒号后面的部分
	Exclusive            string            `json:",omitempty"` //none:不独占，instance：实例独占；app：同一个appname可以共用，不和其他app共用
	Mandate              bool              `json:",omitempty"` //是否必需,默认为 true, 若非必需调度时可能不会给予分配
	Driver               string            `json:",omitempty"` //磁盘driver 类型, 默认为local,
	IncludeVolumeInParam bool              `json:",omitempty"` // 是否把docker create api 中的-v 选项指定的非bind mount volume 也使用这个磁盘
	IncludeVolumeInImage bool              `json:",omitempty"` // 是否把镜像中的 volume 也使用这个磁盘
	Opt                  map[string]string `json:",omitempty"` //driver 的选项
	Label                map[string]string `json:",omitempty"` //磁盘的 label
}

//{ true 1073741824  0 0   true  false false map[] map[]}
func main() {
	var p, q s
	p.ID.ID = 10
	p.name = "hello"
	q.ID.ID = 11
	q.name = "hello"
	fmt.Printf("hello!aaa")
	if !reflect.DeepEqual(p, q) {
		fmt.Printf("equal %v", p)
	}
	if !reflect.DeepEqual(nil, q) {
		fmt.Printf("---- %v, %v", p, nil)
	}
	var a, b DiskSpec
	a.VolumeName = ""
	a.AllowUseHostBootDisk = true
	a.Size = 1073741824
	a.Iops = 0
	a.Iobps = 0
	a.Mandate = true
	a.IncludeVolumeInParam = false
	a.IncludeVolumeInImage = false
	jsonStr, err := json.Marshal(a)
	if err != nil {

		fmt.Printf("\nmarshar error!")
	}
	if err := json.Unmarshal(jsonStr, &b); err != nil {
		fmt.Printf("\nun marshar error!")
	}

	//	b.VolumeName = ""
	//	b.AllowUseHostBootDisk = true
	//	b.Size = 1073741824
	//	b.Iops = 0
	//	b.Iobps = 0
	//	b.Mandate = true
	//	b.IncludeVolumeInParam = false
	//	b.IncludeVolumeInImage = false

	if !reflect.DeepEqual(a, b) {
		fmt.Printf("\ndisk not equal [%v] != [%v]", a, b)
	} else {
		fmt.Printf("\nequal [%v] == [%v]", a, b)
	}
	var c, d map[string]string
	if !reflect.DeepEqual(c, d) {
		fmt.Printf("\n-----disk not equal [%v] != [%v]", c, d)
	}

}
