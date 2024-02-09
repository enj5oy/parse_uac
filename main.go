package main
import uac "github.com/audibleblink/msldapuac"
import "fmt"
import "flag"
import "log"
func main() {
var uacnumber int64
flag.Int64Var(&uacnumber, "uacnumber", 0 ,"Specify UAC value from ldap")
flag.Parse()
if uacnumber==0{
 log.Fatal("Please provide -uacnumber argument")
}
uacProp := int64(uacnumber)
fmt.Println(uac.ParseUAC(uacProp))
}
