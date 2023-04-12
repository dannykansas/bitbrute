import pycurl
import sys
from StringIO import StringIO

def apicheck(address):
	
	apiURL = 'http://blockchain.info/q/addressbalance/'
	
	store = StringIO()
	c = pycurl.Curl()
	c.setopt(pycurl.URL, apiURL + address)
	c.setopt(c.WRITEFUNCTION, store.write)
	c.perform()
	c.close()
	balance = float(store.getvalue())
	
	balance = balance / 100000000
	if balance == 0:
#		print "Zero balance, rejected."
		return
	else:
		print("Balance is: " + str(balance) + " for " + address)
		return address,balance

# Placeholder for local blockchain balance parse lookup
def localcheck(address):
	return

# Set manually, haven't added getopt yet
infile = 'test.txt'	
with open(infile) as f:
	for line in f:
		line = line.split(', ')
		apicheck(line[1].rstrip())
	print("Complete!")
