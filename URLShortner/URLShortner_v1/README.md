# Project UrlShortener version1

You certainly know that some addresses in the browser (called URLs) are (very) long and/or
complex, and that there are services on the web which turn these into a nice short url, to be used
instead. Our project is like that: it is a web service with 2 functionalities:

Add: given a long URL, it returns a short version, e.g.:
http://maps.google.com/maps?f=q&source=s_q&hl=en&geocode=&q=tokyo&sll=37.0625,-95.6
77068&sspn=68.684234,65.566406&ie=UTF8&hq=&hnear=Tokyo,+Japan&t=h&z=9

(A) becomes: http://goto/UrcGq

(B) and stores this pair of data

Redirect: whenever a shortened URL is requested, it redirects the user to the original, long URL:
so if you type (B) in a browser, it redirects you to the page of (A).
