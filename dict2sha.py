#!/usr/bin/python
# 
# Convert dictionary input file to outfile containing 
# bitcoin private key and matching public address.
#
# Don't be rude.
#
# - dannykansas
#

import binascii
import hashlib
import pubgen
import sys, getopt
import count

# Gets SHA256 hash of password and generates matching
# bitcoin public address, then stores to <outfile>
# for later use by checkbal.py
def gethash(line):
	m = hashlib.sha256()
	m.update(line)
	byte = m.digest()
	linehash = binascii.hexlify(byte)
	linehash = '0x'+linehash
	
	# Special thanks to luper~
	byte = int(linehash, 16)
	p = pubgen.addy(byte)
	return linehash, p

def main(argv):
	infile = ''
	outfile = ''
	try:
		opts, args = getopt.getopt(argv, "hi:o:",["ifile=","ofile="])
	except getopt.GetopError:
		print 'dict2sha.py -i <inputfile> -o <outputfile>'
		sys.exit(2)
	for opt, arg in opts:
		if opt == '-h':
			print 'text.py -i <inputfile> -o <outputfile>'
			sys.exit()
		elif opt in ("-i", "--infile"):
			infile = arg
		elif opt in ("-o", "--outfile"):
			outfile = arg
	count.getcount(infile)
	i = 0
	fout = open(outfile,'w')

	with open(infile) as f:
		for line in f:
			print "Hashing: " + line.rstrip()
			linehash, p = gethash(line.rstrip())
#			print "linehash is %s", linehash
			fout.write(linehash + ', ' + p + '\n')
	return		

if __name__ == "__main__":
   main(sys.argv[1:])


