package gopay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func TestMd5(t *testing.T) {
	st := "appid=wxdaa2ab9ef87b5497&nonceStr=9k20rM66parD2U49&package=prepay_id=wx29164301554772fbc70d1d793335446010&signType=MD5&timeStamp=1548751382&key=GFDS8j98rewnmgl45wHTt980jg543wmg"
	hash := md5.New()
	hash.Write([]byte(st))
	sum := hash.Sum(nil)
	upper := strings.ToUpper(hex.EncodeToString(sum))
	fmt.Println(" ssad  ", upper)
}

func TestBodyMap_MarshalXML(t *testing.T) {

	maps := make(BodyMap)
	maps.Set("name", "jerry")
	maps.Set("age", 28)
	maps.Set("phone", "13212345678")

	bytes, err := xml.Marshal(&maps)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("ssss:", string(bytes))

}
