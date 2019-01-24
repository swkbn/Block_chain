package main

import "os"

//判断是否由此文件
func isFileExist(fileName string)bool  {

	_,err:=os.Stat(fileName)
	//一定不要使用os.Isxist
	if os.IsNotExist(err) {

		return false
	}
	return true

}
