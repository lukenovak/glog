# Goblog: a Golang blogging platform
I wanted a place to put some personal thoughts, and wasn't satisfied with
using Medium like a normal person would be, so I decided to write my own.
The ultimate end goal of this project is to have a portable, easy to run
blogging platform that can be launched from a machine with minimum overhead.
I'm currently a *long* way away from that, but ultimately the hope is that
the entire application can be downloaded and consequently run with a single
command.

## Running the server

If you want to run your own instance of goblog right now, there's some things
you'll have to do first:

### Depenencies
You will obviously need to install Go if you don't have it already. 

You will also need revel, a web application framework for golang. You can install it
using `go get github.com/revel/revel`. You will also need revel's command line
tool, which can be installed with `go get github.com/revel/cmd/revel`. For more 
information about Revel, you can [check it out here](https://revel.github.io/)!

Finally, you will need to install PostgreSQL, and run the schema in posts.sql.
I'm hoping to support other databases in the future.
