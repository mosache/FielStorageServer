package router

import (
	"FileStorageServer/utils"
	"net/http"
	"strings"

	cache "FileStorageServer/cache/redis"
	"FileStorageServer/dao"
	"FileStorageServer/db"
	"FileStorageServer/model"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

//MutipartUploadInfo MutipartUploadInfo
type MutipartUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int64
	ChunkCount int
}

//MutipartUpLoadInitalization 初始化分块上传信息
func MutipartUpLoadInitalization(w http.ResponseWriter, r *http.Request) {
	//1.获取用户参数
	err := r.ParseForm()

	fileHash := r.FormValue("file_hash")
	fileSize, err := strconv.Atoi(r.FormValue("file_size"))
	username := r.FormValue("user_name")
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}
	//2.获取redis连接
	rConn := cache.RedisPool().Get()
	defer rConn.Close()
	//3.计算分块信息
	uploadInfo := MutipartUploadInfo{
		FileHash:   fileHash,
		FileSize:   fileSize,
		UploadID:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  5 * 1024 * 1024, //5MB
		ChunkCount: int(math.Ceil(float64(fileSize) / (5 * 1024 * 1024))),
	}
	//4.分块信息写入redisl
	rConn.Do("HSET", "MP"+uploadInfo.UploadID, "chunkCount", uploadInfo.ChunkCount)
	rConn.Do("HSET", "MP"+uploadInfo.UploadID, "filehash", uploadInfo.FileHash)
	rConn.Do("HSET", "MP"+uploadInfo.UploadID, "filesize", uploadInfo.FileSize)

	//5.响应客户端

	serveJSON(w, resultMap{
		"status": 1,
		"data":   uploadInfo,
	})
}

//UploadPart 上传文件分块
func UploadPart(w http.ResponseWriter, r *http.Request) {
	//1.获取请求参数
	err := r.ParseForm()
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	// username := r.Form.Get("username")
	uploadID := r.Form.Get("uploadID")
	chunkIndex := r.Form.Get("chunkIndex")
	//2.获取redis连接
	rConn := cache.RedisPool().Get()
	defer rConn.Close()
	//3.获得文件句柄，保存分块内容
	fd, err := os.OpenFile("/data/"+uploadID+"/"+chunkIndex, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}
	defer fd.Close()

	buf := make([]byte, 1024*1024) //1mb
	for {
		n, err := r.Body.Read(buf)
		fd.Write(buf[:n])
		if err != nil {
			break
		}
	}
	//4.更新redis状态
	rConn.Do("HSET", "MP"+uploadID, "chkidx_"+chunkIndex, 1)
	//5.响应客户端
	serveJSON(w, resultMap{
		"status": 1,
		"msg":    "OK",
	})

}

//CompleteUpload 通知上传合并
func CompleteUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uploadID := r.Form.Get("uploadID")
	fileHash := r.Form.Get("filehash")
	fileSize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
	}
	fileName := r.Form.Get("filename")

	rConn := cache.RedisPool().Get()
	defer rConn.Close()

	data, err := redis.Values(rConn.Do("HGETALL", "MP_"+uploadID))
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	totalCount := 0
	chunkCount := 0

	for i := 0; i < len(data); i += 2 {
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))

		if k == "chunkcount" {
			totalCount, _ = strconv.Atoi(v)
		} else if strings.HasPrefix(k, "chkidx_") && v == "1" {
			chunkCount++
		}
	}

	if totalCount != chunkCount {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    "合并文件出错",
		})
		return
	}

	tx, err := db.Db.Begin()
	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	err = dao.InsertFileMeta(tx, &model.FileMeta{FileSha1: fileHash, FileSize: fileSize, FileName: fileName})

	if err != nil {
		tx.Rollback()
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	tokenData := r.Context().Value(ctxKey).(*utils.TokenData)

	err = dao.InsertUserFile(tx, &model.UserFile{FileHash: fileHash, FileSize: fileSize, FileName: fileName, UserID: tokenData.UserID})

	if err != nil {
		tx.Rollback()
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	err = tx.Commit()

	if err != nil {
		serveJSON(w, resultMap{
			"status": 0,
			"msg":    err.Error(),
		})
		return
	}

	serveJSON(w, resultMap{
		"status": 1,
		"msg":   "ok",
	})
}
