from enum import Enum

class TextType(Enum):
  PLAIN_TEXT = "text"
  BOLD_TEXT = "bold"
  ITALIC_TEXT = "italic"
  CODE_TEXT = "code"
  LINK_TEXT = "link"
  IMAGES = "image"
  
class TextNode:
  def __init__(self, text, text_type, url=None):
    self.text = text
    self.text_type = TextType(text_type)
    self.url = url
    
  def __eq__(self, other) -> bool:
    return (
      self.text == other.text 
      and self.text_type == other.text_type
      and self.url == other.url
    )
  
  def __repr__(self) -> str:
    return f"TextNode({self.text}, {self.text_type.value}, {self.url})"