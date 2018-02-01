package main

import (
	"github.com/yanghai23/GoLib/atdb"
	"github.com/yanghai23/GoLib/atfile"
	"github.com/yanghai23/GoLib/aterr"
	"github.com/yanghai23/GoLib/atsafe"
	"database/sql"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"mime/multipart"
	"os"
)
var db *sql.DB
func main() {
	//读取配置文件
	config , err := atdb.ReadDbConfig(atfile.GetCurrentDirectory()+"/logDBConfig.json")
	aterr.CheckErr(err)
	//初始化数据库
	db,err = atdb.InitMysql(*config)
	aterr.CheckErr(err)

	http.HandleFunc("/uploadFile",atsafe.SafeHandle(uploadFile))
	http.HandleFunc("/download",atsafe.SafeHandle(download))
	err = http.ListenAndServe(":9999",nil)
	if err != nil{
		panic("监听端口失败")
	}
}


func uploadFile(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("TESTING .....")
	mf := r.MultipartForm
	var files []*multipart.FileHeader
	if mf != nil {
		fmt.Println("TESTING ...2..")
		files = mf.File["goBook"]
		return
	}

	_,file2 ,err := r.FormFile("goBook")
	aterr.CheckErr(err)
	Write2Local(file2)
	if files != nil {
		WriteFile2Local(files)
	}
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(200)

	w.Write([]byte(string("{\"code\":200,\"msg\":\"上传成功\"}")))
}

func WriteFile2Local(files []*multipart.FileHeader) error {
	for _, file := range files {
		Write2Local(file)
	}
	return nil
}
func Write2Local(fileSource *multipart.FileHeader) error {
		fmt.Println("file.Filename = ", fileSource.Filename)
		f, err := fileSource.Open()
		if err != nil {
			fmt.Println("err -- ", err)
			return err
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("err == ", err)
			return err
		}
		//路径是`yibaLog`+时间+uid组成
		path := "goBook/" + time.Now().Format("2006-01-02") + "/"
		fmt.Println("strTime", "strTime = "+path)
		os.MkdirAll(path, 0777) //创建文件夹
		fileNameStr := fileSource.Filename
		file, err := atfile.CreateFile(path, fileNameStr)
		aterr.CheckErr(err)
		file.Write([]byte(data))
		file.Close()
		//将日子文件对应关系，存入到数据库，方便查找日志和用户的关系
		insertLogData(time.Now().Format("2006-01-02"), path, fileNameStr)

	return nil
}

func insertLogData(time, path, fileName string) {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO FilePath(date,path,fileName) VALUES(?,?,?)")
	defer stmt.Close()
	aterr.CheckErr(err)
	res, err := stmt.Exec(time, path, fileName)
	aterr.CheckErr(err)
	id, err := res.LastInsertId()
	aterr.CheckErr(err)
	fmt.Println(id)
}

func download(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	if len(r.Form["path"]) > 0  && len(r.Form["fileName"]) > 0{
		localPath :=string( r.Form["path"][0])+string(r.Form["fileName"][0])
		buf,err := ioutil.ReadFile(localPath)
		aterr.CheckErr(err)
		w.Write(buf)

	}
}

