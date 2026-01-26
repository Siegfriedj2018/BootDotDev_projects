from stats import word_count, char_count, sort_contents
import sys


def get_book_text(filepath):
  with open(filepath) as f:
    file_contents = f.read()
    
  return file_contents

def format_print(file_path, book_data, words, chars, list_dict):
  print("============ BOOKBOT ============")
  print(f"Analyzing book found at {file_path}...")
  print(
    f'''----------- Word Count ----------
Found {words} total words
--------- Character Count -------'''
  )
  for val in list_dict:
    if not val["char"].isalpha():
      continue
    print(f"{val["char"]}: {val["num"]}")
    
  print("============= END ===============")

def main():
  if len(sys.argv) != 2:
    print("Usage: python3 main.py <path_to_book>")
    sys.exit(1)
  
  file_path = sys.argv[1]
  book_content = get_book_text(file_path)
  words = word_count(book_content)
  chars = char_count(book_content)
  list_dict = sort_contents(chars)
  
  format_print(file_path, book_content, words, chars, list_dict)
  
if __name__ == main():
  main()