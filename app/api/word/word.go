package word

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"write_code_in_/app/ship/msg"

	"github.com/gogf/gf/os/gfile"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/unidoc/unioffice/document"
)

// Create 生成文档接口
func Create(r *ghttp.Request) {
	name := r.GetForm("name") //项目名称
	page := r.GetForm("page") //word页数
	path := r.GetForm("path") //项目路径

	files := []string{}
	files, _ = GetAll(path.(string), files)

	totalRow := 0
	pageRow := 46
	wordPath := "./resource/"

	doc := document.New()

	for _, item := range files {

		if totalRow < pageRow*gconv.Int(page) {
			fi, err := os.Open(item)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			defer fi.Close()

			br := bufio.NewReader(fi)
			for {
				a, _, c := br.ReadLine()
				if c == io.EOF {
					break
				}
				totalRow++
				// fmt.Println(string(a))
				createParaRun(doc, string(a))
			}
		} else {
			break
		}
	}
	doc.SaveToFile(wordPath + gconv.String(name) + "管理系统代码文档.docx")
	msg.Success(r, files, "成功")
}

// GetFile 获取文件列表
func GetFile(r *ghttp.Request) {
	files := []string{}
	name, _ := filepath.Abs(filepath.Dir("./resource"))
	fullPath := name + "\\resource"

	files, _ = GetAllWord(fullPath, files)

	var data []interface{} = make([]interface{}, len(files))
	for key, item := range files {
		fileInfo := gfile.MTime(fullPath + "\\" + item)
		var itemMap = make(map[string]string)
		itemMap["name"] = item
		itemMap["time"] = fileInfo.String()
		data[key] = itemMap
	}
	msg.Success(r, data, "成功")
}

// Download 下载文件列表
func Download(r *ghttp.Request) {
	name := r.Get("name") //项目名称
	r.Response.ServeFileDownload("./resource/" + name.(string))

}

// Delete 删除文件列表
func Delete(r *ghttp.Request) {
	name := r.Get("name") //项目名称

	err := os.Remove("./resource/" + name.(string))
	responseData := interface{}(nil)
	if err != nil {
		msg.Fail(r, responseData, "失败")

	} else {
		msg.Success(r, responseData, "成功")
	}

}

// GetAll 获取路径下的文件
func GetAll(filePath string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(filePath) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	array := garray.NewStrArrayFrom(g.SliceStr{"vendor", "config", "public", "cache", "log", "tmp"})
	// 如果是文件夹，进行递归处理
	for _, fi := range read {
		if fi.IsDir() { // 判断是否是文件夹
			fullDir := filePath + "\\" + fi.Name() //构造新的路径
			if !array.ContainsI(fi.Name()) {
				files, _ = GetAll(fullDir, files) //文件夹递归
			}
		} else {
			fullDir := filePath + "\\" + fi.Name() //构造新的路径
			fileExt := path.Ext(fi.Name())
			if fileExt == ".php" {
				files = append(files, fullDir) //追加路径
				// fmt.Println(filePath)
			}
		}
	}
	return files, nil
}

// GetAllWord 获取word
func GetAllWord(filePath string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(filePath) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	array := garray.NewStrArrayFrom(g.SliceStr{"vendor", "config", "public", "cache", "log", "tmp"})
	// 如果是文件夹，进行递归处理
	for _, fi := range read {
		if fi.IsDir() { // 判断是否是文件夹
			fullDir := filePath + "\\" + fi.Name() //构造新的路径
			if !array.ContainsI(fi.Name()) {
				files, _ = GetAll(fullDir, files) //文件夹递归
			}
		} else {
			fullDir := fi.Name() //构造新的路径
			fileExt := path.Ext(fi.Name())
			if fileExt == ".docx" {
				files = append(files, fullDir) //追加路径
				// fmt.Println(filePath)
			}
		}
	}
	return files, nil
}

//ReadFiles 读取文件内容
func ReadFiles(filePth string) (string, error) {
	// open readAll

	var str string

	fi, err := os.Open(filePth)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer fi.Close()
		strings, err := ioutil.ReadAll(fi)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			str = string(strings)
		}
	}
	return str, err
}

// ReF 读取文件内容
func ReF(filePth string) {
	fi, err := os.Open(filePth)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}

// Write0 aaa
func Write0(strTest string) {
	fileName := "1.txt"
	var d = []byte(strTest)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}

// createParaRun aaa
func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetFontFamily("Times New Roman")
	// run.Properties().SetSize(10)
	run.AddText(s)
	return run
}
