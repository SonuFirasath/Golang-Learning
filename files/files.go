package main

import (
	"fmt"
	"os"
)

func main() {

	// Reading a file

	// f,err := os.Open("example.txt")

	// if err != nil {
	// 	// You can now log the error or panic
	// 	panic(err)
	// }

	// fileInfo,err := f.Stat()

	// if err!= nil {
	// 	panic(err)
	// }

	// fmt.Println("The name of the file is : ",fileInfo.Name())
	// fmt.Println("Is it is folder : ",fileInfo.IsDir())
	// fmt.Println("File size : ",fileInfo.Size())
	// fmt.Println("File Permission : ",fileInfo.Mode())
	// fmt.Println("File modified at :",fileInfo.ModTime())

	// Read file this is also one type of reading data and we will see the simpler version of this reading a file logic as well....

	// f,err := os.Open("example.txt")

	// if err != nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)

	// d,err := f.Read(buf)

	// if err!= nil{
	// 	panic(err)
	// }

	// for i := range(buf){

	// 	fmt.Println("data : ",d,string(buf[i]))
	// }

	// Simpler version of reading a file in golang...

	// This is a simpler form of reading a file but we should not use it often because the ReadFile will load all the data of file only once....if the file is small that's not a problem if the file is big that is not a viable solution because if bigger file which contains some GB's of data is loaded at a time the recourses might not be enough for that work....

	// f,err := os.ReadFile("example.txt")

	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	// Read Folders

	// dir , err :=os.Open("..")

	// if err != nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo , err := dir.ReadDir(-1)

	// for _,fi := range(fileInfo){
	// 	fmt.Println(fi.Name())
	// }

	// creating a file

	// f, err := os.Create("example2.txt")

	// if err!= nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hello Firasath")

	// To over write the existing data we need to truncate the data and using seek we can go to the beginning of the file and then write a content over there...

	// f.Truncate(0)
	// f.Seek(0,0)
	// f.WriteString("Hi Readers How are you ?")

	// Easier method of writing to a file if we need to override the existing data in a file.
	
	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	// // Read and write to another file (streaming fashion of the code .....)

	// sourceFile , err := os.Open("example.txt")

	// if err!=nil{
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")

	// if err!=nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)

	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b,err := reader.ReadByte()

	// 	if err != nil{
	// 		if err.Error() != "EOF"{

	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)

	// 	if e != nil{
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()

	// fmt.Println("Written to a new file successfully.....")

	// Deleting a file

	err := os.Remove("example2.txt")

	if err!=nil{
		panic(err)
	}

	fmt.Println("The file is deleted successfully .....")

}