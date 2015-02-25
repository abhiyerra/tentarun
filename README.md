# tentarun

Run commands on multiple machines.

## Examples

Run a command on multiple hosts.

With a username and password:

    tentarun -u username -p password -h host1,host2,host3 "tail /var/log/nginx/access.log"
    tentarun -p password -h host1,host2,host3 "tail /var/log/nginx/access.log"

With a username and keyfile

    tentarun -u username -k keyfile -h host1,host2,host3 "tail /var/log/nginx/access.log"
    tentarun -k keyfile -h host1,host2,host3 "tail /var/log/nginx/access.log"

Pass environment variables

    tentarun -e "FOO=bar BAR=foo" -h host1,host2,host3 "echo $FOO"

## Uses

 - Combine outputs of multiple machines and and parse it using
   local tools.
 - Get metrics on multiple machines.

## Credits

 - http://kukuruku.co/hub/golang/ssh-commands-execution-on-hundreds-of-servers-via-go
 - http://golang-basic.blogspot.com/2014/06/step-by-step-guide-to-ssh-using-go.html


## The MIT License (MIT)

Copyright (c) 2015 Abhi Yerra <abhi@berkeley.edu>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
