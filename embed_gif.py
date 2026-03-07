import sys
import os

readme_path = 'README.md'
base64_path = 'media/base64_final.txt'

if not os.path.exists(base64_path):
    print(f"Error: {base64_path} not found.")
    sys.exit(1)

with open(base64_path, 'rb') as f:
    data = f.read()
    if data.startswith(b'\xff\xfe'): # UTF-16 LE BOM
        base64_content = data[2:].decode('utf-16').strip()
    elif data.startswith(b'\xef\xbb\xbf'): # UTF-8 BOM
        base64_content = data[3:].decode('utf-8').strip()
    else:
        try:
            base64_content = data.decode('utf-8').strip()
        except UnicodeDecodeError:
            base64_content = data.decode('utf-16').strip()

if not os.path.exists(readme_path):
    print(f"Error: {readme_path} not found.")
    sys.exit(1)

# Re-read original README (I'll try to find a clean version or just fix the corrupted one)
# Since I only added content, I can try to find the target and everything before/after it.
with open(readme_path, 'r', encoding='utf-8', errors='ignore') as f:
    content = f.read()

# Target text
target = '**Instant, frictionless file sharing and receiving across your local network via CLI.**'

if target in content:
    parts = content.split(target)
    # The first part is everything before the target
    # The second part contains the corrupted insertion and everything after it
    
    # We want to find the </div> that follows the target in the second part
    second_part = parts[1]
    closing_div_index = second_part.find('</div>')
    
    if closing_div_index != -1:
        after_div = second_part[closing_div_index:]
        
        insertion = f'\n\n<br>\n\n<img src="{base64_content}" alt="magshare demo" width="800">\n\n'
        new_readme = parts[0] + target + insertion + after_div
        
        with open(readme_path, 'w', encoding='utf-8') as f:
            f.write(new_readme)
        print("Success: README.md updated and fixed encoding.")
    else:
        print("Error: Could not find closing </div> after target.")
        sys.exit(1)
else:
    print("Error: Could not find target text in README.md.")
    sys.exit(1)
