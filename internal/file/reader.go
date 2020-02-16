package file

var filePath = "/Users/chenle/Workspaces/Code/company/pcps/internal/file/temp/camera.log"

//ReadFileLineByLine 一行行读取文件
// func ReadFileLineByLine(p string) (int, error) {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer f.Close()
// 	scanner := bufio.NewScanner(f)
// 	line := 1

// 	for scanner.Scan() {
// 		str := scanner.Text()

// 		pcpsd.ParseLogCameraInfo(str)

// 		time.Sleep(time.Second * 5)
// 		line++
// 	}
// 	return 0, nil
// 	if err := scanner.Err(); err != nil {
// 		// Handle the error
// 	}
// }

//parseCameraInfo 解析摄像头日志中的信息

// int1, _ := strconv.Atoi(nv)

// f,err := os.Open(path)
// if err != nil {
//     return 0,err
// }
// defer f.Close()

// // Splits on newlines by default.
// scanner := bufio.NewScanner(f)

// line := 1
// // https://golang.org/pkg/bufio/#Scanner.Scan
// for scanner.Scan() {
//     if strings.Contains(scanner.Text(),"yourstring") {
//         return line,nil
//     }

//     line++
// }

// if err := scanner.Err(); err != nil {
//     // Handle the error
// }
