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
        try:
            # python dict() 转换为 JSON 字符串
            aweme_id = record['aweme_id']
            video_desc = record.get('desc', "")
            create_time = record['create_time']
            music_id = record['music']['id']
            source_id = aweme_id
            share_url = record.get('share_url', "")
            statistics_id = aweme_id
            status = json.dumps(record.get('status', []))
            text_extra = json.dumps(record.get('text_extra', []))
            is_top = record.get('is_top', False)
            share_info = json.dumps(record.get('share_info', []))
            duration = record.get('duration', 0)
            image_infos = record.get('image_infos', [])
            risk_infos = json.dumps(record.get('risk_infos', []))
            position = record.get('position', "")
            author_user_id = record['author_user_id']
            prevent_download = record.get('prevent_download', False)
            long_video = record.get('long_video', [])
            aweme_control = json.dumps(record.get('aweme_control', []))
            images = json.dumps(record.get('images', []))
            suggest_words = json.dumps(record.get('suggest_words', []))
            video_tag = json.dumps(record.get('video_tag', []))

            music_id = music_id
            title = record['music'].get('title', "")
            author = record['music'].get('author', "")
            cover_medium = json.dumps(record['music'].get('cover_medium', []))
            cover_thumb = json.dumps(record['music'].get('cover_thumb', []))
            play_url = json.dumps(record['music'].get('play_url', []))
            music_duration = record['music'].get('duration', 0)
            user_count = record['music'].get('user_count', 0)
            owner_nickname = record['music'].get('owner_nickname', "")
            is_original = record['music'].get('is_original', False)
            owner_id = record['music'].get('owner_id')

            # source_id = record['aweme_id']
            play_addr = json.dumps(record['video'].get('play_addr', []))
            cover = json.dumps(record['video'].get('cover', []))
            poster = record['video'].get('poster', "")
            video_height = record['video'].get('height', 0)
            video_width = record['video'].get('width', 0)
            ratio = record['video'].get('ratio', "default")
            use_static_cover = record['video'].get('use_static_cover', False)
            video_duration = record['video'].get('duration', 0)
            horizontal_type = record['video'].get('horizontal_type', 1)


            # statistics_id = record['aweme_id']
            admire_count = record['statistics'].get('admire_count', 0)
            comment_count = record['statistics'].get('comment_count', 0)
            digg_count = record['statistics'].get('digg_count', 0)
            collect_count = record['statistics'].get('collect_count', 0)
            play_count = record['statistics'].get('play_count', 0)
            share_count = record['statistics'].get('share_count', 0)


            # tb_videos insert
            sql = """INSERT INTO tb_videos (aweme_id, video_desc, create_time, music_id, source_id, share_url, statistics_id, status, text_extra, is_top, share_info, duration, image_infos, risk_infos, position, author_user_id, prevent_download, long_video, aweme_control, images, suggest_words, video_tag) 
                             VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""
            values = (aweme_id, video_desc, create_time, music_id, source_id,
                      share_url, statistics_id, status, text_extra, is_top,
                      share_info, duration, image_infos, risk_infos, position,
                      author_user_id, prevent_download, long_video, aweme_control,
                      images, suggest_words, video_tag)
            cursor.execute(sql, values)
            connection.commit()

            # tb_source insert
            sql_source = """INSERT INTO tb_source (id, play_addr, cover, poster, height, width, ratio, use_static_cover, duration, horizontal_type)
                                       VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""
            values_source = (source_id, play_addr, cover, poster, video_height, video_width, ratio, use_static_cover, video_duration, horizontal_type)
            cursor.execute(sql_source, values_source)
            connection.commit()

            # tb_statistics insert
            sql_statistics = """INSERT INTO tb_statistics (id, admire_count, comment_count, digg_count, collect_count, play_count, share_count)
                                           VALUES (%s, %s, %s, %s, %s, %s, %s)"""
            values_statistics = (statistics_id, admire_count, comment_count, digg_count, collect_count, play_count, share_count)
            cursor.execute(sql_statistics, values_statistics)
            connection.commit()

            # tb_music insert
            sql_music = """INSERT INTO tb_music (id, title, author, cover_medium, cover_thumb, play_url, duration, user_count, owner_nickname, is_original, owner_id)
                                               VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""
            values_music = (
            music_id, title, author, cover_medium, cover_thumb, play_url, music_duration, user_count, owner_nickname,
            is_original, owner_id)
            cursor.execute(sql_music, values_music)
            connection.commit()

        except mysql.connector.Error as err:
            print(f"Error: {err}, Record ID: {record['aweme_id']}")
            connection.rollback()

    cursor.close()
    connection.close()

if __name__ == "__main__":
    file_path = '/Users/sslee/Documents/RecSys/douyin/public/data/user_video_list/user-SUNMENG333.json'  # 指定你的 JSON 文件路径
    json_data = read_json_file(file_path)
    insert_into_database(json_data)
