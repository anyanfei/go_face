package main

import (
	"fmt"
	"github.com/Kagami/go-face"
	"log"
)

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	modelDir  = dataDir + "/models"
)

// 图片中的人名
var labels = []string{
	"萧敬腾",
	"周杰伦",
	"苏永康",
	"王力宏",
	"陶喆",
	"林俊杰",
}

func main() {
	fmt.Println("Face Recognition...")
	// 初始化识别器，训练模型后组成神经
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()
	fmt.Println("Recognizer Initialized")
	// 先来一张合影，让其从左到右识别后并给这些人唯一id
	faces, err := rec.RecognizeFile("heying.jpg")
	if err !=nil{
		log.Println(err)
		return
	}
	// 将合影中人物映射到唯一 ID, 然后将唯一 ID 和对应人物相关联
	var samples []face.Descriptor
	var peoples []int32
	for i,f := range faces{
		samples = append(samples,f.Descriptor)
		//给每张脸唯一id
		peoples = append(peoples,int32(i))
	}
	rec.SetSamples(samples,peoples)
	RecognizePeople(rec,"jay.jpg")
	RecognizePeople(rec,"linjunjie.jpg")
	RecognizePeople(rec,"taozhe.jpg")
}

// RecognizePeople 人脸识别方法，传入识别器和照片路径，打印对应人物 ID，人物名字
func RecognizePeople(rec *face.Recognizer,file string){
	oneFace , err := rec.RecognizeSingleFile(file)
	if err !=nil{
		log.Fatalf("无法识别：%v",err)
	}
	if oneFace == nil{
		log.Fatalf("图片上不止一张人脸")
	}
	// 根据特征得到唯一id
	faceId := rec.Classify(oneFace.Descriptor)
	if faceId <0 {
		log.Fatalf("无法区分")
	}
	fmt.Println(faceId)
	fmt.Println(labels[faceId])
}
