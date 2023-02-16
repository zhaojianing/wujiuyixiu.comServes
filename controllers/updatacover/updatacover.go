package updatacover

import (
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"os"
	"servers/utils"
	"strings"
)

type UpdataCover struct {
	beego.Controller
}

func (s *UpdataCover) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *UpdataCover) Post() {
	println("fill name")
	f, h, err := s.GetFile("file")
	println("fill name : ", h.Filename)

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	if err != nil {
		log.Fatal("err ", err)
		return
	}
	num, _ := utils.RadomTimeNum()
	// s.SaveToFile("file", "static/img/"+h.Filename)
	result := strings.LastIndex(h.Filename, ".")
	println("path", result)
	s.SaveToFile("file", "static/img/"+num+h.Filename[result:])
	reqs["data"] = "static/img/" + num + h.Filename[result:]
	defer f.Close()
}

func (s *UpdataCover) Delete() {
	imgName := s.GetString("img_name")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	result := strings.LastIndex(imgName, "/static/img/")
	imgUrl := "." + imgName[result:]
	println("imgUrl", imgUrl)
	err := os.Remove(imgUrl)
	if err != nil {
		reqs["code"] = utils.RECODE_REMOLEERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_REMOLEERR)
		reqs["data"] = err
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = "success"
}
