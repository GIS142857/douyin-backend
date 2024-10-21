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
        print(record['uid'])
        try:
            # python dict() 转换为 JSON 字符串
            uid = record['uid']
            short_id = record.get('short_id', 0)
            unique_id = record.get('unique_id', "")
            gender = record.get('gender')
            user_age = record.get('user_age', -1)
            nickname = record.get('nickname', "")
            country = record.get('country', "")
            province = record.get('province', "")
            district = record.get('district', "")
            city = record.get('city', "")
            signature = record.get('signature', "")
            ip_location = record.get('ip_location', "")
            birthday_hide_level = record.get('birthday_hide_level')
            can_show_group_card = record.get('can_show_group_card')
            aweme_count = record.get('aweme_count', 0)
            total_favorited = record.get('total_favorited', 0)
            favoriting_count = record.get('favoriting_count', 0)
            follower_count = record.get('follower_count', 0)
            following_count = record.get('following_count', 0)
            forward_count = record.get('forward_count', 0)
            public_collects_count = record.get('public_collects_count', 0)
            mplatform_followers_count = record.get('mplatform_followers_count', 0)
            max_follower_count = record.get('max_follower_count', 0)
            follow_status = record.get('follow_status')
            follower_status = record.get('follower_status')
            follower_request_status = record.get('follower_request_status')
            cover_colour = record.get('cover_colour', "#03997706")
            cover_url = json.dumps(record.get('cover_url', []))
            white_cover_url = json.dumps(record.get('white_cover_url', []))
            share_info = json.dumps(record.get('share_info', []))
            commerce_info = json.dumps(record.get('commerce_info', []))
            commerce_user_info = json.dumps(record.get('commerce_user_info', []))
            commerce_user_level = record.get('commerce_user_level', 0)
            card_entries = json.dumps(record.get('card_entries', []))
            avatar_168x168 = json.dumps(record.get('avatar_168x168', []))
            avatar_300x300 = json.dumps(record.get('avatar_300x300', []))


            # tb_users insert
            sql = """INSERT INTO tb_users (uid,short_id,unique_id,gender,user_age,nickname,country,province,district,city,signature,ip_location,birthday_hide_level,can_show_group_card,aweme_count,total_favorited,favoriting_count,follower_count,following_count,forward_count,public_collects_count,mplatform_followers_count,max_follower_count,follow_status,follower_status,follower_request_status,cover_colour,cover_url,white_cover_url,share_info,commerce_info,commerce_user_info,commerce_user_level,card_entries,avatar_168x168,avatar_300x300) 
                             VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""
            values = (uid,short_id,unique_id,gender,user_age,nickname,country,province,district,city,signature,ip_location,birthday_hide_level,can_show_group_card,aweme_count,total_favorited,favoriting_count,follower_count,following_count,forward_count,public_collects_count,mplatform_followers_count,max_follower_count,follow_status,follower_status,follower_request_status,cover_colour,cover_url,white_cover_url,share_info,commerce_info,commerce_user_info,commerce_user_level,card_entries,avatar_168x168,avatar_300x300)
            cursor.execute(sql, values)
            connection.commit()

        except mysql.connector.Error as err:
            print(f"Error: {err}, Record ID: {uid}")
            connection.rollback()

    cursor.close()
    connection.close()

if __name__ == "__main__":
    file_path = '/Users/sslee/Documents/RecSys/douyin/public/data/users.json'  # 指定你的 JSON 文件路径
    json_data = read_json_file(file_path)
    insert_into_database(json_data)
