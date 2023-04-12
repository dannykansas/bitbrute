import random
import time

# Number of lines to write
num_lines = 1000000000

# Size of the buffer in bytes
buffer_size = 1024 * 1024 * 100  # 100 MB

# Open file in write mode with a buffer
with open('output_file.txt', 'w', buffering=buffer_size) as f:
    # Write lines in batches
    start_time = time.time()
    batch_size = 1000000
    for i in range(0, num_lines, batch_size):
        # Generate batch of random lines
        lines = [str(random.random()) + '\n' for _ in range(batch_size)]
        # Write batch of lines to file
        f.writelines(lines)
        # Flush buffer to disk after each batch
        f.flush()
    end_time = time.time()

# Print total time taken to write the file
print(f"Total time taken: {end_time - start_time:.2f} seconds")

