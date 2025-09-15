# coding=utf-8
import copy
import json
import re
from collections import defaultdict


class ImageInfo:
    line_index: int
    name: str
    ver: str

def main():
    with open('rewrite_setting.json') as f:
        rewrite_setting = json.load(f)

    target_files: list[str] = rewrite_setting["target_files"]
    for target_file in target_files:
        with open(target_file, encoding='utf-8') as f:
            lines = f.readlines()

        imageinfo_map = get_imageinfo_map(lines, rewrite_setting)
        if len(imageinfo_map) > 0:
            print(f"rewrite {target_file}\n")
            rewrite_lines = rewrite(lines, imageinfo_map, rewrite_setting)
            with open(target_file, 'w') as f:
              f.writelines(rewrite_lines)


def get_imageinfo_map(lines: list[str], replace_ver_map: dict[str, str]) -> dict[str, list[ImageInfo]]:
    imageinfo_map: dict[str, list[ImageInfo]] = defaultdict(lambda: [])
    pat = r'(\s*#?\s*)image\s*:\s*([^\s]+)\s*:\s*([^\s,]+)'
    for i, line in enumerate(lines):
        ret = re.match(pat, line)
        if ret is None or len(ret.groups()) < 3:
            continue
        name = ret.group(2)
        ver = ret.group(3)
        img = ImageInfo()
        img.line_index = i
        img.name = name
        img.ver = ver
        if img.name in replace_ver_map:
            imageinfo_map[name].append(img)
    return imageinfo_map

def rewrite(lines: list[str], imageinfo_map: dict[str, list[ImageInfo]], rewrite_setting: dict[str, str]) -> list[str]:
    def get_text(name: str, ver: str, is_comment: bool) -> str:
        if is_comment:
            item_name = "    #image"
        else:
            item_name = "    image"
        return f"{item_name}: {name}:{ver}\n"
    rewrite_lines = copy.deepcopy(lines)
    keys = list(reversed(list(imageinfo_map.keys())))
    for k in keys:
        imageinfo_list = imageinfo_map[k]
        found_rewrite = False
        for image in imageinfo_list:
            if image.name in rewrite_setting and image.ver == rewrite_setting[image.name]:
                found_rewrite = True
                break
        for image in imageinfo_list:
            rewrite_lines[image.line_index] = get_text(image.name, image.ver, True)

        if found_rewrite:
            for image in imageinfo_list:
                if image.name in rewrite_setting and image.ver == rewrite_setting[image.name]:
                    rewrite_lines[image.line_index] = get_text(image.name, image.ver, False)
        else:
            slist = sorted(imageinfo_list, key=lambda x: x.line_index, reverse=True)
            new_imageinfo = slist[0]
            new_name = new_imageinfo.name
            new_ver = rewrite_setting[new_name]
            image_text = get_text(new_name, new_ver, False)
            rewrite_lines.insert(new_imageinfo.line_index + 1, image_text)
    return rewrite_lines

if __name__ == "__main__":
    # /usr/bin/python3 -version
    main()