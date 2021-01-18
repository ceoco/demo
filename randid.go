package randid_test

import (
	//"database/sql"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
	//_ "github.com/go-sql-driver/mysql"
)

//RandomTestBase Generate several random numbers that are not repeated
func RandomTestBase() {
	//Test 5 time
	//for i := 0; i < 1; i++ {
	nums := generateRandomNumber(100000, 999999, 800000)
	fmt.Println(len(nums))
	//fmt.Println(nums)

	var wireteString = ""
	var filename = "./output1.txt"
	var f *os.File
	var err1 error
	//checkFileIsExist In the first way: use io.WriteString write to the file
	if checkFileIsExist(filename) { //if the file exists
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //open file
		fmt.Println("file does not exist")
	} else {
		f, err1 = os.Create(filename) //create file
		fmt.Println("file does not exist")
	}
	checkErr(err1)

	wireteString = "insert into king_identifier (recordid) values "
	n, err1 := io.WriteString(f, wireteString) //write to the file(string)
	checkErr(err1)
	fmt.Printf("write %d bytes n", n)

	for k, v := range nums {
		fmt.Println(k)
		fmt.Println(v)
		wireteString = "(" + strconv.Itoa(v) + "),"
		n, err1 := io.WriteString(f, wireteString) //write to file(string)
		checkErr(err1)
		fmt.Printf("write %d bytes n", n)
	}
	io.WriteString(f, ";")
	//}
}

//generate count non-repetitive random numbers from start to end.
func generateRandomNumber(start int, end int, count int) []int {
	//range check
	if end < start || (end-start) < count {
		return nil
	}

	//db, err := sql.Open("mysql", "test:test@/abc")
	//checkErr(err)

	//A slice in which the results are stored.
	nums := make([]int, 0)	
	//Random number generator,adding a timestamp to ensure that the random number generated each time is different
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//Generate random number
		num := r.Intn((end - start)) + start

		//Find duplicates
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
			//insert data
			//stmt, err := db.Prepare("insert into king_identifier (recordid) values (?)")
			//checkErr(err)
			//res, err := stmt.Exec(num)
			//checkErr(err)
			//id, err := res.LastInsertId()
			//checkErr(err)
			//
			//fmt.Println(id)
			//stmt.Close();
		}
	}

	return nums
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


//To determine whether the file exists,there is a return true,and there is no return false.
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func TestTransRate(t *testing.T) {

	RandomTestBase()

}
