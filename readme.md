Elevate is a simple and efficient website status checker built using Go. It checks whether a list of websites is reachable by sending HTTP requests and reporting their status.


Release link : (You can download the .exe file and run on your local computer) : 
https://github.com/devWisz/Elevate/releases/tag/1.0


Features
1.Concurrent website checking using goroutines
2.Fast and efficient execution
3.Timeout handling to prevent hanging requests
4.Clear success and failure responses for each URL
5.Minimal and easy-to-understand codebase


Example Output : 

http://sarjakkhanal.com is up and running!
http://google.com is up and running!
http://facebook.com is up and running!
http://this-site-should-fail.com might be down! Error: Get "...": dial tcp: lookup failed

License

This project is open-source and free to use for learning and experimentation.