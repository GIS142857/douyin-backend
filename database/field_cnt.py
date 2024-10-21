import json


# 递归遍历 JSON 数据，收集所有字段
def collect_fields(data, field_set=None, parent_key=""):
    if field_set is None:
        field_set = set()

    if isinstance(data, dict):
        for key, value in data.items():
            full_key = f"{parent_key}.{key}" if parent_key else key
            field_set.add(full_key)
            collect_fields(value, field_set, full_key)
    elif isinstance(data, list):
        for item in data:
            collect_fields(item, field_set, parent_key)

    return field_set


# 读取 JSON 文件
def read_json_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        data = json.load(file)
    return data


if __name__ == "__main__":
    file_path = '/Users/sslee/Documents/RecSys/douyin/public/data/comments/video_id_6686589698707590411.json'  # 替换为你的 JSON 文件路径
    json_data = read_json_file(file_path)

    # 获取所有字段
    fields = collect_fields(json_data)

    # 打印字段
    print("JSON 中出现的所有字段:")
    i = 0
    for field in sorted(fields):
        print(i, field)
        i += 1
