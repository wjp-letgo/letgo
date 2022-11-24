package mediaspace

import (
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	mediaspaceEntity "github.com/wjpxxx/letgo/x/api/shopee/mediaspace/entity"
	"github.com/wjpxxx/letgo/lib"
)

//MediaSpace
type MediaSpace struct{
	Config *shopeeConfig.Config
}

//InitVideoUpload
//@Title Initiate video upload session. Video duration should be between 10s and 60s (inclusive).
//@Description https://open.shopee.com/documents?module=91&type=1&id=531&version=2
func (m *MediaSpace)InitVideoUpload(fileMd5 string,fileSize int)mediaspaceEntity.InitVideoUploadResult{
	method:="media_space/init_video_upload"
	result:=mediaspaceEntity.InitVideoUploadResult{}
	params:=lib.InRow{
		"file_md5":fileMd5,
		"file_size":fileSize,
	}
	err:=m.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UploadVideoPart
//@Title Upload video file by part using the upload_id in initiate_video_upload. The request Content-Type of this API should be of multipart/form-data
//@Description https://open.shopee.com/documents?module=91&type=1&id=532&version=2
func (m *MediaSpace)UploadVideoPart(videoUploadID string,partSeq int,contentMd5 string,partContentPath string)mediaspaceEntity.UploadVideoPartResult{
	method:="media_space/upload_video_part"
	result:=mediaspaceEntity.UploadVideoPartResult{}
	params:=lib.InRow{
		"video_upload_id":videoUploadID,
		"part_seq":partSeq,
		"content_md5":contentMd5,
		"@part_content":partContentPath,
	}
	err:=m.Config.HttpPostFile(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//CompleteVideoUpload
//@Title Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
//@Description https://open.shopee.com/documents?module=91&type=1&id=533&version=2
func (m *MediaSpace)CompleteVideoUpload(videoUploadID string,partSeqList []int,reportData mediaspaceEntity.ReportDataEntity)mediaspaceEntity.CompleteVideoUploadResult{
	method:="media_space/complete_video_upload"
	result:=mediaspaceEntity.CompleteVideoUploadResult{}
	params:=lib.InRow{
		"video_upload_id":videoUploadID,
		"part_seq_list":partSeqList,
		"report_data":reportData,
	}
	err:=m.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetVideoUploadResult
//@Title Query the upload status and result of video upload.
//@Description https://open.shopee.com/documents?module=91&type=1&id=534&version=2
func (m *MediaSpace)GetVideoUploadResult(videoUploadID string)mediaspaceEntity.GetVideoUploadResult{
	method:="media_space/get_video_upload_result"
	result:=mediaspaceEntity.GetVideoUploadResult{}
	params:=lib.InRow{
		"video_upload_id":videoUploadID,
	}
	err:=m.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//CancelVideoUpload
//@Title Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
//@Description https://open.shopee.com/documents?module=91&type=1&id=533&version=2
func (m *MediaSpace)CancelVideoUpload(videoUploadID string)mediaspaceEntity.CancelVideoUploadResult{
	method:="media_space/cancel_video_upload"
	result:=mediaspaceEntity.CancelVideoUploadResult{}
	params:=lib.InRow{
		"video_upload_id":videoUploadID,
	}
	err:=m.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UploadImage
//@Title Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
//@Description https://open.shopee.com/documents?module=91&type=1&id=533&version=2
func (m *MediaSpace)UploadImage(image string)mediaspaceEntity.UploadImageResult{
	method:="media_space/upload_image"
	result:=mediaspaceEntity.UploadImageResult{}
	params:=lib.InRow{
		"@image":image,
	}
	err:=m.Config.HttpPostFile(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}