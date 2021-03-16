/*
   __author__ : stray_camel
  __description__ : 继承使用gorm构造orm对象
  __REFERENCES__:
  __date__: 2021-03-16
*/
package utils

import (
// "gorm.io/gorm"
)

func fetch_data_sql(sql string, params map[string]string) {

}

/*
————————————————
版权声明：本文为CSDN博主「王中阳」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/w425772719/article/details/113944875
*/
// BatchSave 批量插入数据
// func BatchSave(db *gorm.DB, emps []*CardInviteLogs) error {
// 	var buffer bytes.Buffer
// 	sql := "insert into `card_invite_logs` (`userid`,`invite_userid`,`has_chance`,`is_sign`) values"
// 	if _, err := buffer.WriteString(sql); err != nil {
// 		return err
// 	}
// 	for i, e := range emps {
// 		if i == len(emps)-1 {
// 			buffer.WriteString(fmt.Sprintf("('%d','%d',%d,%d);", e.Userid, e.InviteUserid, e.HasChance, e.IsSign))
// 		} else {
// 			buffer.WriteString(fmt.Sprintf("('%d','%d',%d,%d),", e.Userid, e.InviteUserid, e.HasChance, e.IsSign))
// 		}
// 	}
// 	return db.Exec(buffer.String()).Error
// }
