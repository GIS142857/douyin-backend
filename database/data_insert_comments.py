import json
import mysql.connector

# 数据库连接配置
db_config = {
    'host': 'localhost',
    'user': 'root',
    'password': '04102410',
    'database': 'Mall_Db'
}

# 读取 JSON 文件
def read_json_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        data = json.load(file)
    return data

# 插入数据到数据库
def insert_into_database(data):
    connection = mysql.connector.connect(**db_config)
    cursor = connection.cursor()

    for record in data:
        print(record['comment_id'])
        try:
            # python dict() 转换为 JSON 字符串
            comment_id = record['comment_id']
            create_time = record.get('create_time')
            ip_location = record.get('ip_location')
            aweme_id = record.get('aweme_id')
            content = record.get('content')
            is_author_digged = record.get('is_author_digged')
            is_folded = record.get('is_folded')
            is_hot = record.get('is_hot')
            user_buried = record.get('user_buried')
            user_digged = record.get('user_digged')
            digg_count = record.get('digg_count')
            user_id = record.get('user_id')
            sec_uid = record.get('sec_uid')
            short_user_id = record.get('short_user_id')
            user_unique_id = record.get('user_unique_id')
            user_signature = record.get('user_signature')
            nickname = record.get('nickname')
            avatar = record.get('avatar')
            sub_comment_count = record.get('sub_comment_count')
            last_modify_ts = record.get('last_modify_ts')

            # tb_comments insert
            sql = """INSERT INTO tb_comments (comment_id,create_time,ip_location,aweme_id,content,is_author_digged,is_folded,is_hot,user_buried,user_digged,digg_count,user_id,sec_uid,short_user_id,user_unique_id,user_signature,nickname,avatar,sub_comment_count,last_modify_ts) 
                             VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""
            values = (comment_id,create_time,ip_location,aweme_id,content,is_author_digged,is_folded,is_hot,user_buried,user_digged,digg_count,user_id,sec_uid,short_user_id,user_unique_id,user_signature,nickname,avatar,sub_comment_count,last_modify_ts)
            cursor.execute(sql, values)
            connection.commit()

        except mysql.connector.Error as err:
            print(f"Error: {err}, Record ID: {comment_id}")
            connection.rollback()

    cursor.close()
    connection.close()

if __name__ == "__main__":
    file_path = '/Users/sslee/Documents/RecSys/douyin/public/data/comments/video_id_7321200290739326262.json'  # 指定你的 JSON 文件路径
    json_data = read_json_file(file_path)
    insert_into_database(json_data)
