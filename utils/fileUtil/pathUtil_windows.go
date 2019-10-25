// +build windows

package fileUtil

/*
获取程序运行路径，对windows的支持目前不完整
*/
func GetCurrentDirectory() string {
	return ""
}

func GetAbsUrl(relaPath string)string{
	//if strings.Contains(relaPath,":"){
	//	return relaPath
	//}else{
	//	return relaPath
	//}
	return relaPath
}