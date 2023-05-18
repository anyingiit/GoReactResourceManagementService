package models

import "gorm.io/gorm"

type WebService struct {
	gorm.Model
	Host             string          `gorm:"type:varchar(100);not null;comment:服务地址" json:"host"`
	Port             string          `gorm:"type:varchar(100);not null;comment:服务端口" json:"port"`
	WebServiceTypeID uint            `gorm:"type:int;not null;comment:服务类型ID"`
	WebServiceType   *WebServiceType `gorm:"foreignKey:WebServiceTypeID"`
	ServiceID        uint            `gorm:"not null"`
	Service          *Service        `gorm:"polymorphic:Owner;polymorphicValue:web_services"`
}

// 关于Service和WebServiceType为什么要用指针类型：
// 		当使用指针类型时，如果TaskParamType为nil，则表示该任务没有任务参数类型。
//		如果不使用指针类型，则无法区分任务是否具有任务参数类型，因为默认值将是该类型的零值。
//		此外，使用指针类型可以在查询时更容易地实现可选的任务参数类型。

// 关于gorm的关系：
// 1. 一对一关系，使用指针类型
// 2. 一对多关系，使用切片，切片中的元素指向的都是指针
// 3. 多对多关系，使用切片类型
// 4. 多态关系，使用指针类型
// 5. 多态多对多关系，使用切片类型

// table name
