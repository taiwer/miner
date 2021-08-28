package dbPlotOpr

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)

func inserGlobal() {
	fmt.Println("insert Global ...")
	groups := [7]Global{
		{Name: "defaultStallItemNameList", Val: 0, Descs: "有哪些物品可以设置价格", Text: "灵魂晶石;魔魂晶石;幻魔晶石"},
		{Name: "defaultDropItemText", Val: 0, Descs: "默认丢弃的物品", Text: "幻兽锁;装备锁"},
		{Name: "defaultSellItemRuleText", Val: 0, Descs: "默认卖金币物品规则", Text: "*图鉴卡;:装备"},
		{Name: "defaultWareHouseRuleText", Val: 0, Descs: "默认存仓库物品规则列表", Text: "回城卷;*图鉴卡;*晶石;:装备;*%;伤害*;升级经验*;战斗力*;曙光战魂"},
		{Name: "defaultPickUpItemRuleText", Val: 0, Descs: "默认自动捡物品", Text: "战士攻防型;冥河戒灵·维列;*攻防型;异界精灵瑞拉;*奥伦;*晶石;金币;大堆金币;:1091000;:1091020;:装备;裁决骑士*;*图鉴卡;*凯恩;*凯特;*%;战斗力*;曙光战魂"},
		{Name: "defaultUseItemRuleText", Val: 0, Descs: "默认自动使用", Text: "*掠宝袋;1小时祝福礼包;神祝幻尘;*恒晶石礼盒"},
		{Name: "9*InstanceItemNameList", Val: 0, Descs: "(9星副本物品名称列表)", Text: strings.Join(InstanceItemNameList9, ";")},
	}
	for _, v := range groups {
		if _, err := AddGlobal(&v, nil); err != nil {
			zap.L().Error("insert Global", zap.String("err", err.Error()))
		}
	}
	fmt.Println("insert Global end", len(groups))
}

var InstanceItemNameList9 = []string{
	"魔魂晶石",
	"幻魔晶石",
	"灵魂晶石",
	"30星神兽碎片+1",
	"30星神兽碎片+2",
	"30星神兽碎片+3",
	"速效圣兽灵药",
	//"蓝色祝福碎片",
	//"紫色祝福碎片",
	//"卡利亚手记·十",
	"卡利亚手记·九",
	"卡利亚手记·八",
	"卡利亚手记·七", //灵魂
	"卡利亚手记·六",
	//"卡利亚手记·五", //魔魂
	//"卡利亚手记·四", //幻魔"
	"精英骑士纹章", //幻魔"s
	"12星XO礼包",
	"6星O礼包",
	"*星*礼包",
	"炎魔督军之杖", //星之祝福
}

//自动存仓库
var defaultCunChangKuItemNames = []string{
	"魔魂晶石",
	"幻魔晶石",
	"灵魂晶石",
	"超值转世水晶",
	"*星*礼包",
}

//自动使用
var defaultOpenItemNames = []string{
	"蓝色祝福碎片",
	"紫色祝福碎片",
	"卡利亚手记·十",
	"卡利亚手记·九",
	"卡利亚手记·八",
	"卡利亚手记·七", //灵魂
	"卡利亚手记·六",
	"卡利亚手记·五", //魔魂
	"卡利亚手记·四", //幻魔
	"精英骑士纹章",  //幻魔
}

//自动使用
var defaultSellItemNames = []string{
	"魔魂晶石",
	"幻魔晶石",
	"双子星*礼包",
	"白羊星*礼包",
	"金牛星*礼包",
}
