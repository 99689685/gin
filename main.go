package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

func main() {
	//// 引用数据库
	//model.InitDb()
	//// 路由地址
	//routers.InitRouter()
	// 加载身份证图像
	imagePath := "snipaste.png" // 身份证图像路径

	// 读取身份证图像
	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Failed to read image")
		return
	}
	defer img.Close()

	// 进行图像处理，如裁剪、灰度化、二值化等操作

	// 裁剪出姓名和身份证号码区域的图像
	nameImg := img.Region(nameRegion())         // 自定义函数，根据实际情况裁剪出姓名区域
	idNumberImg := img.Region(idNumberRegion()) // 自定义函数，根据实际情况裁剪出身份证号码区域
	defer nameImg.Close()
	defer idNumberImg.Close()

	// 转换为灰度图像
	grayName := gocv.NewMat()
	gocv.CvtColor(nameImg, &grayName, gocv.ColorBGRToGray)
	defer grayName.Close()

	grayIDNumber := gocv.NewMat()
	gocv.CvtColor(idNumberImg, &grayIDNumber, gocv.ColorBGRToGray)
	// 进行图像二值化处理
	binaryName := gocv.NewMat()
	gocv.Threshold(grayName, &binaryName, 0, 255, gocv.ThresholdBinaryInv|gocv.ThresholdOtsu)
	defer binaryName.Close()

	binaryIDNumber := gocv.NewMat()
	gocv.Threshold(grayIDNumber, &binaryIDNumber, 0, 255, gocv.ThresholdBinaryInv|gocv.ThresholdOtsu)
	defer binaryIDNumber.Close()

	// 保存二值化后的图像用于文字识别
	gocv.IMWrite("name_binary.jpg", binaryName)
	gocv.IMWrite("id_number_binary.jpg", binaryIDNumber)

	// 使用Tesseract OCR进行文字识别
	nameText := ocr("name_binary.jpg")
	idNumberText := ocr("id_number_binary.jpg")

	fmt.Println("姓名:", nameText)
	fmt.Println("身份证号码:", idNumberText)

}

// 自定义函数，根据实际情况裁剪出姓名区域
func nameRegion() image.Rectangle {
	// TODO: 根据实际情况修改裁剪区域的坐标和尺寸
	x := 100
	y := 200
	width := 300
	height := 50

	return image.Rect(x, y, x+width, y+height)
}

// 自定义函数，根据实际情况裁剪出身份证号码区域
func idNumberRegion() image.Rectangle {
	// TODO: 根据实际情况修改裁剪区域的坐标和尺寸
	x := 100
	y := 300
	width := 300
	height := 50

	return image.Rect(x, y, x+width, y+height)
}

// 使用Tesseract OCR进行文字识别
func ocr(imagePath string) string {
	// TODO: 调用Tesseract OCR库进行文字识别并返回结果
	return ""
}
