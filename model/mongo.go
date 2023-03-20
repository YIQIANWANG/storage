package model

// Storage 存储节点信息
type Storage struct {
	StorageAddress string `bson:"storageAddress" json:"storageAddress"`
	GroupID        string `bson:"groupID" json:"groupID"`
	Capacity       int    `bson:"capacity" json:"capacity"`
	AvailableCap   int    `bson:"availableCap" json:"availableCap"`
	UpdateTime     int64  `bson:"updateTime" json:"updateTime"`
}
