package controllers

import (
	"fmt"
	"github.com/Comdex/imgo"
	"io/ioutil"
	"os"
	"strconv"
)

type ConvController struct {
	BaseController
}

// 判断尺寸是否允许
// 使用map储存，key对应宽，v对应高
func Isallowedsize(width int, height int) bool {
	stored_size := map[int]int{}
	stored_size[20] = 20
	stored_size[30] = 30
	stored_size[50] = 50
	stored_size[100] = 100
	stored_size[120] = 120
	stored_size[150] = 150
	stored_size[200] = 200
	stored_size[250] = 250
	stored_size[300] = 300

	if stored_size[width] == height {
		return true
	} else {
		return false
	}

}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func (c *ConvController) Get() {
	pa1 := c.Ctx.Input.Param(":pa1")
	pa2 := c.Ctx.Input.Param(":pa2")
	pa3 := c.Ctx.Input.Param(":pa3")
	pa4 := c.Ctx.Input.Param(":pa4")
	pa5 := c.Ctx.Input.Param(":pa5")
	pa_name := c.Ctx.Input.Param(":pa6")
	pa_width := c.Ctx.Input.Param(":pa7")
	pa_height := c.Ctx.Input.Param(":pa8")
	pa_suffix := c.Ctx.Input.Param(":pa9")

	pa_width_int, _ := strconv.Atoi(pa_width)
	pa_height_int, _ := strconv.Atoi(pa_height)

	file_save_path := pa1 + "/" + pa2 + "/" + pa3 + "/" + pa4 + "/" + pa5 + "/" + pa_name + "_" + pa_width + "_" + pa_height + "." + pa_suffix
	file_save_path = fmt.Sprintf("%s%s", c.Data["filepath"], file_save_path)
	file_real_path := pa1 + "/" + pa2 + "/" + pa3 + "/" + pa4 + "/" + pa5 + "/" + pa_name + "." + pa_suffix
	file_real_path = fmt.Sprintf("%s%s", c.Data["filepath"], file_real_path)

	if !Exist(file_real_path) {
		c.Data["error_msg"] = "file not found"
		c.Abort("404")
	}

	if Exist(file_save_path) {
		s, _ := ioutil.ReadFile(file_save_path)
		c.Ctx.Output.Header("Content-Type", "image/PNG; charset=utf-8")
		c.Ctx.Output.Body(s)
	} else {
	    
	    if !Isallowedsize(pa_width_int, pa_height_int){
	        c.Data["error_msg"] = "file size not allowed"
    		c.Abort("404")
	    }
	    
		img, _ := imgo.ResizeForMatrix(file_real_path, pa_width_int, pa_height_int)

		var err error = nil
		switch pa_suffix {
		case "PNG":
			fallthrough
		case "png":
			err = imgo.SaveAsPNG(file_save_path, img)
		default:
			//    //保存为jpeg,100为质量，1-100
			err = imgo.SaveAsJPEG(file_save_path, img, 30)
		}

		if err != nil {
			c.Data["error_msg"] = "file convert fail"
			c.Abort("404")
		}

		s, _ := ioutil.ReadFile(file_save_path)
		c.Ctx.Output.Header("Content-Type", "image/PNG; charset=utf-8")
		c.Ctx.Output.Body(s)
	}
}
