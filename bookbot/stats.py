def word_count(book_contents):
  return len(book_contents.split())


def char_count(book_contents):
  contents = {}
  for char in book_contents:
    lowered = char.lower()
    if lowered not in contents:
      contents[lowered] = 0
    contents[lowered] += 1
    
  return contents

def sort_on(item):
  return item["num"]

def sort_contents(chars_dict):
  list_chars = []
  for d in chars_dict:
    list_chars.append(dict(char=d,num=chars_dict[d]))
  sorted = list_chars.sort(reverse=True, key=sort_on)
  return list_chars
  