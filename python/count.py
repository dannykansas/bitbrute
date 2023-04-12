## Optimized file reader/counter
## Usage: python count.py <file>
import sys

def getcount(file_path):
    buffer_size = 1024 * 512  # 512KB buffer
    line_count = 0

    with open(file_path, 'rb') as f:
        buffer = f.read(buffer_size)
        while buffer:
            line_count += buffer.count(b'\n')
            buffer = f.read(buffer_size)

    print(f"{file_path} has {line_count} lines")

if __name__ == "__main__":
    getcount(sys.argv[1])
