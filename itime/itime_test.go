package itime

import (
	"fmt"
	"testing"
	"time"
)

func TestNowMinuteStart(t *testing.T) {
	fmt.Println(time.Now().Unix())
	
	tt, _ := time.ParseInLocation("2006-01", time.Now().Format("2006-01"), time.Local)
fmt.Println(	tt.AddDate(11, 48, 0))

	//
	//fmt.Println(last)
	//
	//ttt, _ := time.ParseInLocation("2006-01", time.Unix(last, 0).Format("2006-01"), time.Local)
	//
	//fmt.Println(ttt.Unix(), ttt.Month())

	//x, y := imath.Divmod(5, 12)
	//fmt.Println(x, y)
	//
	//year := time.Now().Year()
	//year -= int(x)
	//month := time.Now().Month()
	//div := int64(month) - y
	//if  div>0{
	//	month = time.Month(div)
	//}else if div == 0{
	//	month = time.December
	//	year -= 1
	//
	//}else {
	//	month = time.Month(div+div)
	//	year-=1
	//}
	//
	//fmt.Println(	time.Date(year, month,1,0,0,0,0,time.Local))

}
