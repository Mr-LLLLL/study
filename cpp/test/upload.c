#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include "deal_mysql.h"


int store_data()
{
    // 连接数据库
    MYSQL* conn = NULL;
    conn = msql_conn("root", "1234", "test");
    if(conn == NULL)
    {
        printf("upload_file mysql 数据库连接失败!");
        return -1;
    }
    // 设置数据库编码，否则，中文插入乱码
    mysql_query(conn, "set names utf8");
    // 插入数据
    //插入
    char buf[1024];
    sprintf(buf, "insert into file (name, field) values ('filename', 'field')");
    printf("<br> sql: %s<br>\n", buf);
    // 执行sql语句
    // 如果成功，它返回0。
    if ( mysql_query (conn, buf) != 0 ) {
        printf("upload_file mysql 数据插入失败");
        return -1;
    }
    // 关闭数据库连接
    mysql_close(conn);
    return 0;
}

int main ()
{
	store_data();

    return 0;
}
