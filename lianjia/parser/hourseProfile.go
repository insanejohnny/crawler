package parser

import (
	"regexp"
	"strconv"

	"github.com/crawler/engine"
	"github.com/crawler/model"
)

var totalPriceRe = regexp.MustCompile(`<div class="price "><span class="total">([0-9]+)</span><span class="unit"><span>万</span></span>`)
var unitPriceRe = regexp.MustCompile(`<div class="unitPrice"><span class="unitPriceValue">([0-9]+)<i>元/平米</i></span></div>`)
var houseTypeRe = regexp.MustCompile(`<li><span class="label">房屋户型</span>([^<]+)</li>`)
var floorRe = regexp.MustCompile(`<li><span class="label">所在楼层</span>([^<]+)</li>`)
var buildingAreaRe = regexp.MustCompile(`<li><span class="label">建筑面积</span>([^<]+)</li>`)
var directionRe = regexp.MustCompile(`<li><span class="label">房屋朝向</span>([^<]+)</li>`)
var housesPerFloorRe = regexp.MustCompile(`<li><span class="label">梯户比例</span>([^<]+)</li>`)
var communityNameRe = regexp.MustCompile(`<span class="label">小区名称</span><a href="/xiaoqu/[0-9]+/"[^>]*>([^<]+)</a>`)

func ParseHouseProfile(contents []byte, title string) engine.ParseResult {
	houseProfile := model.HouseProfile{}

	houseProfile.Title = title

	totalPrice, err := strconv.Atoi(extractString(contents, totalPriceRe))
	if err == nil {
		houseProfile.TotalPrice = totalPrice
	}
	// houseProfile.TotalPrice = extractString(contents, totalPriceRe)
	unitPrice, err := strconv.Atoi(extractString(contents, unitPriceRe))
	if err == nil {
		houseProfile.UnitPrice = unitPrice
	}

	houseProfile.HouseType = extractString(contents, houseTypeRe)
	houseProfile.Floor = extractString(contents, floorRe)
	houseProfile.BuildingArea = extractString(contents, buildingAreaRe)
	houseProfile.Direction = extractString(contents, directionRe)
	houseProfile.HousesPerFloor = extractString(contents, houseTypeRe)
	houseProfile.CommunityName = extractString(contents, communityNameRe)

	result := engine.ParseResult{
		Items: []interface{}{houseProfile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
