import sys

def getcount(argv):
	i = 0
	with open(argv) as f:
		for line in f:
			i = i+1
		print argv + " has " + str(i) + " lines"

if __name__ == "__main__":
	getcount(sys.argv[1])

