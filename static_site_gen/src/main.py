from textnode import TextNode

def main():
  test = TextNode(text="this is some anchor text", text_type="link", url="https://www.boot.dev")
  print(f"{test}")

if __name__ == "__main__":
  main()