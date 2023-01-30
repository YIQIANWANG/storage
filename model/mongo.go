package model

// Storage 存储实例信息
type Storage struct {
	StorageAddress string `bson:"storageAddress" json:"storageAddress"`
	GroupID        string `bson:"groupID" json:"groupID"`
	UpdateTime     int64  `bson:"updateTime" json:"updateTime"`
}
