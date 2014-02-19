cnamulator
==========

## Objective
```
Cnamulator is a utility written in the GO programming language. The
purpose of this tool is to perform a CNAM lookup on phone numbers
in order to derive the organization in which they belong to. This is
especially useful when performing non-disclosure (i.e., black box)
security assessments. For instance, scraping publicly available PDF's
for phone numbers and then using Cnamulator to verify the organization.
```


## Usage
```
$ cnamulator -h
Usage of cnamulator:
  -file="": a list of phone numbers
  -phone="": a single phone number
  -sid="": the opencnam api sid
  -token="": the opencnam api auth token
```

## Installation
```
# Installation
# -----------------------------------------------
# Cnamulator was tested on OSX ML
# ----------- OSX ---------------
# OSX Deps: Go Language in order to compile on system @ http://golang.org/
# OPENCNAM: Register @ https://www.opencnam.com/ for sid and token
```

## Sample Run
```
$ cnamulator -file=../numbers.txt -sid=opencnam-sid -token=opencnam-auth_token
	{
	    "results": [
	        {
	            "number": "16502530001",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530002",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530003",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530004",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530005",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530006",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530007",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530008",
	            "name": "GOOGLE INC"
	        },
	        {
	            "number": "16502530009",
	            "name": "GOOGLE INC"
	        }
	    ]
	}
```

## Developing
```
Alpha code under active development
```

## Thanks
```
# Contributor: Tom Steele
# Twitter: @_tomsteele
# Provided idiomatic GO syntax help
```

## Contact
```
# Author: Chris Patten
# Contact (Email): cpatten[t.a.]packetresearch[t.o.d]com
# Contact (Twitter): @packetassailant
```



