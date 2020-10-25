package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{Name: name}
}

type Player struct {
	Name string
	Sex  string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(player Player, monster Monster) Mission {
	return Mission{
		Player:  player,
		Monster: monster,
	}
}

func (m *Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

type EmployeeMessage struct {
	Body     []EmployeeMessageBody `json:"body"`
	Time     int64                 `json:"time"`
	Type     string                `json:"type"`
	TypeDesc string                `json:"typeDesc"`
}

type EmployeeMessageBody struct {
	NewValue EmployeeMessageBodyNewValue `json:"newValue"`
	OldValue struct {
		Leader string `json:"leader"`
	} `json:"oldValue"`
	ChangeType string `json:"changeType"`
}

type EmployeeMessageBodyNewValue struct {
	Position             string `json:"position"`
	Domain               string `json:"domain"`
	EntryDate            string `json:"entryDate"`
	City                 string `json:"city"`
	DisplayNumber        string `json:"displayNumber"`
	EmployTypeCode       string `json:"employTypeCode"`
	TalentTypeCode       string `json:"talentTypeCode"`
	DepartmentPathName   string `json:"departmentPathName"`
	Avatar               string `json:"avatar"`
	HrStatus             string `json:"hrStatus"`
	OfficeCode           string `json:"officeCode"`
	CommonName           string `json:"commonName"`
	Department           string `json:"department"`
	Sex                  string `json:"sex"`
	DepartmentPathNumber string `json:"departmentPathNumber"`
	LeaderPath           string `json:"leaderPath"`
	LeaderName           string `json:"leaderName"`
	LeaveDate            string `json:"leaveDate"`
	FormalDate           string `json:"formalDate"`
	PositionNumber       string `json:"positionNumber"`
	Office               string `json:"office"`
	Leader               string `json:"leader"`
	CompanyName          string `json:"companyName"`
	LastName             string `json:"lastName"`
	FirstName            string `json:"firstName"`
	EmployType           string `json:"employType"`
	TalentType           string `json:"talentType"`
	JobStatus            string `json:"jobStatus"`
	CityCode             string `json:"cityCode"`
	JobStatusCode        string `json:"jobStatusCode"`
	Company              string `json:"company"`
	Name                 string `json:"name"`
	SalaryCompanyNumber  string `json:"salaryCompanyNumber"`
	SalaryCompanyName    string `json:"salaryCompanyName"`
}

var str = "{\\\"body\\\":[{\\\"newValue\\\":{\\\"position\\\":\\\"课*问\\\",\\\"domain\\\":\\\"zhouyu01\\\",\\\"entryDate\\\":\\\"2019-01-07\\\",\\\"city\\\":\\\"北京市\\\",\\\"displayNumber\\\":\\\"10002646\\\",\\\"employTypeCode\\\":\\\"10\\\",\\\"talentTypeCode\\\":\\\"3\\\",\\\"departmentPathName\\\":\\\"百家互联-高*堂663-小*部683-销*售753\\\",\\\"avatar\\\":\\\"\\\",\\\"hrStatus\\\":\\\"A\\\",\\\"officeCode\\\":\\\"101\\\",\\\"commonName\\\":\\\"周*雨\\\",\\\"department\\\":\\\"10000753\\\",\\\"sex\\\":\\\"男\\\",\\\"departmentPathNumber\\\":\\\"10000001-10000663-10000683-10000753\\\",\\\"leaderPath\\\":\\\"\\\",\\\"leaderName\\\":\\\"贾*飞\\\",\\\"leaveDate\\\":\\\"\\\",\\\"formalDate\\\":\\\"2020-09-16\\\",\\\"positionNumber\\\":\\\"100350\\\",\\\"office\\\":\\\"北京博彦科技\\\",\\\"leader\\\":\\\"10006583\\\",\\\"companyName\\\":\\\"北京跟谁学科技有限公司\\\",\\\"lastName\\\":\\\"周\\\",\\\"firstName\\\":\\\"雨*雨\\\",\\\"employType\\\":\\\"正式员工\\\",\\\"talentType\\\":\\\"销售\\\",\\\"jobStatus\\\":\\\"正常在职\\\",\\\"cityCode\\\":\\\"101\\\",\\\"jobStatusCode\\\":\\\"10\\\",\\\"company\\\":\\\"103\\\",\\\"name\\\":\\\"周*雨\\\",\\\"salaryCompanyNumber\\\":\\\"505\\\",\\\"salaryCompanyName\\\":\\\"北京跟谁学科技有限公司\\\"},\\\"oldValue\\\":{\\\"leader\\\":\\\"\\\"},\\\"changeType\\\":\\\"update\\\"},{\\\"newValue\\\":{\\\"position\\\":\\\"行*管\\\",\\\"domain\\\":\\\"lihuijun\\\",\\\"entryDate\\\":\\\"2015-10-19\\\",\\\"city\\\":\\\"郑州市\\\",\\\"displayNumber\\\":\\\"10000379\\\",\\\"employTypeCode\\\":\\\"10\\\",\\\"talentTypeCode\\\":\\\"7\\\",\\\"departmentPathName\\\":\\\"百家互联-高*堂663-高*心713-综*部717\\\",\\\"avatar\\\":\\\"\\\",\\\"hrStatus\\\":\\\"A\\\",\\\"officeCode\\\":\\\"\\\",\\\"commonName\\\":\\\"李*君\\\",\\\"department\\\":\\\"10000717\\\",\\\"sex\\\":\\\"女\\\",\\\"departmentPathNumber\\\":\\\"10000001-10000663-10000713-10000717\\\",\\\"leaderPath\\\":\\\"\\\",\\\"leaderName\\\":\\\"周*鹏\\\",\\\"leaveDate\\\":\\\"\\\",\\\"formalDate\\\":\\\"\\\",\\\"positionNumber\\\":\\\"100068\\\",\\\"office\\\":\\\"\\\",\\\"leader\\\":\\\"10018239\\\",\\\"companyName\\\":\\\"北京百家互联科技有限公司郑州分公司\\\",\\\"lastName\\\":\\\"李\\\",\\\"firstName\\\":\\\"惠*君\\\",\\\"employType\\\":\\\"正式员工\\\",\\\"talentType\\\":\\\"通用管理及支持\\\",\\\"jobStatus\\\":\\\"试用期\\\",\\\"cityCode\\\":\\\"102\\\",\\\"jobStatusCode\\\":\\\"20\\\",\\\"company\\\":\\\"108\\\",\\\"name\\\":\\\"李*君\\\",\\\"salaryCompanyNumber\\\":\\\"\\\",\\\"salaryCompanyName\\\":\\\"\\\"},\\\"oldValue\\\":{\\\"leader\\\":\\\"10000001\\\"},\\\"changeType\\\":\\\"update\\\"}],\\\"time\\\":1600249196095,\\\"type\\\":\\\"employLeaderChange\\\",\\\"typeDesc\\\":\\\"员工直属leader变更\\\"}|medusa_ps_employ_job_leaderChange\""

func main() {
	//mission := InitMission("dj")
	//mission.Start()
	//fmt.Println(strings.Split(str, "|"))
	changes := strings.Split(str, "|")[0]
	fmt.Println(changes)
	newChanges := changes[1 : len(changes)-1]
	changesByte := []byte(newChanges)
	msg := &EmployeeMessage{}
	err := json.Unmarshal(changesByte, msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	//player := Player{
	//	Name: "闫旭",
	//}
	//jPlayer, err := json.Marshal(player)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(jPlayer))
	//}
	//
	//sPlayer := string(jPlayer) + "|你好"
	//sPlayer0 := strings.Split(sPlayer, "|")[0]
	//fmt.Println(sPlayer0)
	//json.RawMessage(player)
	//
	//player2 := &Player{}
	//err = json.Unmarshal([]byte(sPlayer), player2)
	//fmt.Println(err, player2)
}
