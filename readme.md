# Server Monitor
## NOTE: project currently unfinished in development.

This is a highly-custom server monitor, so I'm unsure if it will be useful to anyone else, however I figured I'd still upload it.

The purpose of the monitor is I have a raspberry pi running a number of social media bots; they log their posts to a log file and are running as seperate processes with seperate log files (just using tmux). The raspberry pi is behind a router I don't want to open up to the world or set up a static IP address to talk directly with it. My solution is this monitor service watches the log files, and if something is no longer getting logged (i.e. the process has crashed) it will post a notice on a designated URL on my site, which can be read from a sandboxed page on my phone.

On my site end, it's basically just a folder in the bucket I use to host my site. I've uploaded a simple html and css page, and "onClick" of a button it checks to see if there are pages uploaded. If so, it reads them and posts the content in javascript. Then when I get home, I can restart the server and clear the pages in the bucket.

Everything can be configured using the server.properties file in the server. The server is written in Go Lang, the site is obviously written with JS and is otherwise static.
