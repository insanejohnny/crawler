package model

type HouseProfile struct {
	Title          string //标题
	TotalPrice     int    //总价
	UnitPrice      int    //单价
	HouseType      string //房型
	Floor          string //楼层
	BuildingArea   string //建筑面积
	Direction      string //朝向
	HousesPerFloor string //梯户比例
	CommunityName  string //小区名称
}
