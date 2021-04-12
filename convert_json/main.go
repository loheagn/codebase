package main

import (
	"encoding/json"
	"fmt"
)

type CreateOption struct {
	Name       string `json:"projectName" binding:"Required"`
	OrgName    string `json:"organizationName"`
	ExpName    string `json:"expName" binding:"Required"`
	ExpID      int64  `json:"expId" binding:"Required"`
	CourseName string `json:"courseName" binding:"Required"`
	CourseID   int64  `json:"courseId" binding:"Required"`
	IsNewOrg   bool   `json:"isNewOrganization"`
}

func main() {
	c := CreateOption{}
	bs, _ := json.Marshal(c)
	fmt.Println(string(bs))
}
